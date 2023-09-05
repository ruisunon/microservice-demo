// Copyright 2022 Linkall Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package trigger

import (
	"context"
	stdErr "errors"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"

	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"

	eb "github.com/vanus-labs/vanus/client"
	"github.com/vanus-labs/vanus/internal/controller/member"
	"github.com/vanus-labs/vanus/internal/controller/trigger/metadata"
	"github.com/vanus-labs/vanus/internal/controller/trigger/secret"
	"github.com/vanus-labs/vanus/internal/controller/trigger/storage"
	"github.com/vanus-labs/vanus/internal/controller/trigger/subscription"
	"github.com/vanus-labs/vanus/internal/controller/trigger/validation"
	"github.com/vanus-labs/vanus/internal/controller/trigger/worker"
	"github.com/vanus-labs/vanus/internal/convert"
	"github.com/vanus-labs/vanus/internal/primitive"
	"github.com/vanus-labs/vanus/internal/primitive/vanus"
	"github.com/vanus-labs/vanus/observability/log"
	"github.com/vanus-labs/vanus/observability/metrics"
	"github.com/vanus-labs/vanus/pkg/cluster"
	"github.com/vanus-labs/vanus/pkg/errors"
	"github.com/vanus-labs/vanus/pkg/util"
	ctrlpb "github.com/vanus-labs/vanus/proto/pkg/controller"
	metapb "github.com/vanus-labs/vanus/proto/pkg/meta"
)

var _ ctrlpb.TriggerControllerServer = &controller{}

const (
	defaultGcSubscriptionInterval = time.Second * 10
)

func NewController(config Config, mem member.Member) *controller {
	ctrl := &controller{
		config:                config,
		member:                mem,
		needCleanSubscription: map[vanus.ID]string{},
		state:                 primitive.ServerStateCreated,
		cl:                    cluster.NewClusterController(config.ControllerAddr, insecure.NewCredentials()),
		ebClient:              eb.Connect(config.ControllerAddr),
	}
	ctrl.ctx, ctrl.stopFunc = context.WithCancel(context.Background())
	return ctrl
}

type controller struct {
	config                Config
	member                member.Member
	storage               storage.Storage
	secretStorage         secret.Storage
	subscriptionManager   subscription.Manager
	workerManager         worker.Manager
	scheduler             *worker.SubscriptionScheduler
	needCleanSubscription map[vanus.ID]string
	lock                  sync.Mutex
	membershipMutex       sync.Mutex
	isLeader              bool
	ctx                   context.Context
	stopFunc              context.CancelFunc
	state                 primitive.ServerState
	cl                    cluster.Cluster
	ebClient              eb.Client
}

func (ctrl *controller) SetDeadLetterEventOffset(
	ctx context.Context, request *ctrlpb.SetDeadLetterEventOffsetRequest,
) (*emptypb.Empty, error) {
	if ctrl.state != primitive.ServerStateRunning {
		return nil, errors.ErrServerNotStart
	}
	subID := vanus.ID(request.SubscriptionId)
	err := ctrl.subscriptionManager.SaveDeadLetterOffset(ctx, subID, request.GetOffset())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (ctrl *controller) GetDeadLetterEventOffset(
	ctx context.Context, request *ctrlpb.GetDeadLetterEventOffsetRequest,
) (*ctrlpb.GetDeadLetterEventOffsetResponse, error) {
	if ctrl.state != primitive.ServerStateRunning {
		return nil, errors.ErrServerNotStart
	}
	subID := vanus.ID(request.SubscriptionId)
	offset, err := ctrl.subscriptionManager.GetDeadLetterOffset(ctx, subID)
	if err != nil {
		return nil, err
	}
	return &ctrlpb.GetDeadLetterEventOffsetResponse{Offset: offset}, err
}

func (ctrl *controller) CommitOffset(
	ctx context.Context, request *ctrlpb.CommitOffsetRequest,
) (*ctrlpb.CommitOffsetResponse, error) {
	if ctrl.state != primitive.ServerStateRunning {
		return nil, errors.ErrServerNotStart
	}
	resp := new(ctrlpb.CommitOffsetResponse)
	for _, subInfo := range request.SubscriptionInfo {
		if len(subInfo.Offsets) == 0 {
			continue
		}
		id := vanus.ID(subInfo.SubscriptionId)
		offsets := convert.FromPbOffsetInfos(subInfo.Offsets)
		err := ctrl.subscriptionManager.SaveOffset(ctx, id, offsets, request.ForceCommit)
		if err != nil {
			resp.FailSubscriptionId = append(resp.FailSubscriptionId, subInfo.SubscriptionId)
			log.Warn(ctx).Err(err).
				Stringer(log.KeySubscriptionID, id).
				Msg("commit offset error")
		}
	}
	return resp, nil
}

func (ctrl *controller) ResetOffsetToTimestamp(
	ctx context.Context, request *ctrlpb.ResetOffsetToTimestampRequest,
) (*ctrlpb.ResetOffsetToTimestampResponse, error) {
	if ctrl.state != primitive.ServerStateRunning {
		return nil, errors.ErrServerNotStart
	}
	if request.Timestamp == 0 {
		return nil, errors.ErrInvalidRequest.WithMessage("timestamp is invalid")
	}
	subID := vanus.ID(request.SubscriptionId)
	sub := ctrl.subscriptionManager.GetSubscription(ctx, subID)
	if sub == nil {
		return nil, errors.ErrResourceNotFound.WithMessage("subscription not exist")
	}
	if sub.Phase != metadata.SubscriptionPhaseStopped {
		return nil, errors.ErrResourceCanNotOp.WithMessage(
			"subscription must be disable can reset offset")
	}
	offsets, err := ctrl.subscriptionManager.ResetOffsetByTimestamp(ctx, subID, request.Timestamp)
	if err != nil {
		return nil, errors.ErrInternal.WithMessage("reset offset by timestamp error").Wrap(err)
	}
	return &ctrlpb.ResetOffsetToTimestampResponse{
		Offsets: convert.ToPbOffsetInfos(offsets),
	}, nil
}

func (ctrl *controller) CreateSubscription(
	ctx context.Context, request *ctrlpb.CreateSubscriptionRequest,
) (*metapb.Subscription, error) {
	if ctrl.state != primitive.ServerStateRunning {
		return nil, errors.ErrServerNotStart
	}
	err := validation.ValidateSubscriptionRequest(ctx, request.Subscription)
	if err != nil {
		log.Info(ctx).Err(err).Msg("create subscription validate fail")
		return nil, err
	}
	sub := convert.FromPbSubscriptionRequest(request.Subscription)
	sub.ID, err = vanus.NewID()
	sub.CreatedAt = time.Now()
	sub.UpdatedAt = time.Now()
	if err != nil {
		return nil, err
	}
	if request.Subscription.Disable {
		sub.Phase = metadata.SubscriptionPhaseStopped
	} else {
		sub.Phase = metadata.SubscriptionPhaseCreated
	}
	err = ctrl.subscriptionManager.AddSubscription(ctx, sub)
	if err != nil {
		return nil, err
	}
	if !request.Subscription.Disable {
		ctrl.scheduler.EnqueueNormalSubscription(sub.ID)
	}
	resp := convert.ToPbSubscription(sub, nil)
	return resp, nil
}

func (ctrl *controller) UpdateSubscription(
	ctx context.Context, request *ctrlpb.UpdateSubscriptionRequest,
) (*metapb.Subscription, error) {
	if ctrl.state != primitive.ServerStateRunning {
		return nil, errors.ErrServerNotStart
	}
	subID := vanus.ID(request.Id)
	sub := ctrl.subscriptionManager.GetSubscription(ctx, subID)
	if sub == nil {
		return nil, errors.ErrResourceNotFound.WithMessage("subscription not exist")
	}
	if sub.Phase != metadata.SubscriptionPhaseStopped {
		return nil, errors.ErrResourceCanNotOp.WithMessage(
			"subscription must be disabled can update")
	}
	if err := validation.ValidateSubscriptionRequest(ctx, request.Subscription); err != nil {
		return nil, err
	}
	if request.Subscription.EventbusId != uint64(sub.EventbusID) {
		return nil, errors.ErrInvalidRequest.WithMessage("can not change eventbus")
	}
	update := convert.FromPbSubscriptionRequest(request.Subscription)
	transChange := 0
	if !sub.Transformer.Exist() && update.Transformer.Exist() {
		transChange = 1
	} else if sub.Transformer.Exist() && !update.Transformer.Exist() {
		transChange = -1
	}
	change := sub.Update(update)
	if !change {
		return nil, errors.ErrInvalidRequest.WithMessage("no change")
	}
	sub.UpdatedAt = time.Now()
	if err := ctrl.subscriptionManager.UpdateSubscription(ctx, sub); err != nil {
		return nil, err
	}
	if transChange != 0 {
		metrics.SubscriptionTransformerGauge.WithLabelValues(sub.EventbusID.Key()).Add(float64(transChange))
	}
	resp := convert.ToPbSubscription(sub, nil)
	return resp, nil
}

func (ctrl *controller) DeleteSubscription(
	ctx context.Context, request *ctrlpb.DeleteSubscriptionRequest,
) (*emptypb.Empty, error) {
	if ctrl.state != primitive.ServerStateRunning {
		return nil, errors.ErrServerNotStart
	}
	subID := vanus.ID(request.Id)
	sub := ctrl.subscriptionManager.GetSubscription(ctx, subID)
	if sub != nil {
		sub.Phase = metadata.SubscriptionPhaseToDelete
		err := ctrl.subscriptionManager.UpdateSubscription(ctx, sub)
		if err != nil {
			return nil, err
		}
		go func(subID vanus.ID, addr string) {
			err := ctrl.gcSubscription(ctrl.ctx, subID, addr)
			if err != nil {
				ctrl.lock.Lock()
				defer ctrl.lock.Unlock()
				ctrl.needCleanSubscription[subID] = addr
			}
		}(subID, sub.TriggerWorker)
	}
	return &emptypb.Empty{}, nil
}

func (ctrl *controller) DisableSubscription(
	ctx context.Context, request *ctrlpb.DisableSubscriptionRequest,
) (*emptypb.Empty, error) {
	if ctrl.state != primitive.ServerStateRunning {
		return nil, errors.ErrServerNotStart
	}
	subID := vanus.ID(request.Id)
	sub := ctrl.subscriptionManager.GetSubscription(ctx, subID)
	if sub == nil {
		return nil, errors.ErrResourceNotFound.WithMessage(
			fmt.Sprintf("subscription %d not exist", subID))
	}
	switch sub.Phase {
	case metadata.SubscriptionPhaseStopped:
		if request.Declaratively {
			return &emptypb.Empty{}, nil
		}
		return nil, errors.ErrResourceCanNotOp.WithMessage("subscription is disabled")
	case metadata.SubscriptionPhaseStopping:
		if request.Declaratively {
			return &emptypb.Empty{}, nil
		}
		return nil, errors.ErrResourceCanNotOp.WithMessage("subscription is disabling")
	}
	sub.Phase = metadata.SubscriptionPhaseStopping
	err := ctrl.subscriptionManager.UpdateSubscription(ctx, sub)
	if err != nil {
		return nil, err
	}
	ctrl.scheduler.EnqueueSubscription(sub.ID)
	return &emptypb.Empty{}, nil
}

func (ctrl *controller) ResumeSubscription(
	ctx context.Context, request *ctrlpb.ResumeSubscriptionRequest,
) (*emptypb.Empty, error) {
	if ctrl.state != primitive.ServerStateRunning {
		return nil, errors.ErrServerNotStart
	}
	subID := vanus.ID(request.Id)
	sub := ctrl.subscriptionManager.GetSubscription(ctx, subID)
	if sub == nil {
		return nil, errors.ErrResourceNotFound.WithMessage(
			fmt.Sprintf("subscrption %d not exist", subID))
	}
	if sub.Phase != metadata.SubscriptionPhaseStopped {
		return nil, errors.ErrResourceCanNotOp.WithMessage("subscription is not disable")
	}
	sub.Phase = metadata.SubscriptionPhasePending
	err := ctrl.subscriptionManager.UpdateSubscription(ctx, sub)
	if err != nil {
		return nil, err
	}
	ctrl.scheduler.EnqueueSubscription(sub.ID)
	return &emptypb.Empty{}, nil
}

func (ctrl *controller) GetSubscription(
	ctx context.Context, request *ctrlpb.GetSubscriptionRequest,
) (*metapb.Subscription, error) {
	if ctrl.state != primitive.ServerStateRunning {
		return nil, errors.ErrServerNotStart
	}
	sub := ctrl.subscriptionManager.GetSubscription(ctx, vanus.ID(request.Id))
	if sub == nil {
		return nil, errors.ErrResourceNotFound.WithMessage("subscription not exist")
	}
	offsets, _ := ctrl.subscriptionManager.GetOffset(ctx, sub.ID)
	resp := convert.ToPbSubscription(sub, offsets)
	return resp, nil
}

func (ctrl *controller) TriggerWorkerHeartbeat(
	heartbeat ctrlpb.TriggerController_TriggerWorkerHeartbeatServer,
) error {
	ctx := ctrl.ctx
	for {
		select {
		case <-ctx.Done():
			_ = heartbeat.SendAndClose(&ctrlpb.TriggerWorkerHeartbeatResponse{})
			return nil
		default:
		}
		if !ctrl.member.IsLeader() {
			_ = heartbeat.SendAndClose(&ctrlpb.TriggerWorkerHeartbeatResponse{})
			return nil
		}
		req, err := heartbeat.Recv()
		if err != nil {
			if !stdErr.Is(err, io.EOF) {
				log.Warn(ctx).Err(err).Msg("heartbeat recv error")
			}
			log.Info(ctx).Msg("heartbeat close")
			return nil
		}
		log.Debug(ctx).
			Str(log.KeyTriggerWorkerAddr, req.Address).
			Interface("subscriptionInfo", req.SubscriptionInfo).
			Msg("heartbeat")
		err = ctrl.triggerWorkerHeartbeatRequest(ctx, req)
		if err != nil {
			return err
		}
	}
}

func (ctrl *controller) triggerWorkerHeartbeatRequest(
	ctx context.Context, req *ctrlpb.TriggerWorkerHeartbeatRequest,
) error {
	now := time.Now()
	for _, subInfo := range req.SubscriptionInfo {
		subscriptionID := vanus.ID(subInfo.SubscriptionId)
		err := ctrl.subscriptionManager.Heartbeat(ctx, subscriptionID, req.Address, now)
		if err != nil {
			log.Warn(ctx).Err(err).
				Str(log.KeyTriggerWorkerAddr, req.Address).
				Stringer(log.KeySubscriptionID, subscriptionID).
				Msg("heartbeat subscription heartbeat error")
		}
	}
	err := ctrl.workerManager.UpdateTriggerWorkerInfo(ctx, req.Address)
	if err != nil {
		log.Info().Str(log.KeyTriggerWorkerAddr, req.Address).
			Msg("unknown trigger worker")
		return errors.ErrResourceNotFound.WithMessage("unknown trigger worker")
	}
	for _, subInfo := range req.SubscriptionInfo {
		if len(subInfo.Offsets) == 0 {
			continue
		}
		offsets := convert.FromPbOffsetInfos(subInfo.Offsets)
		err = ctrl.subscriptionManager.SaveOffset(ctx, vanus.ID(subInfo.SubscriptionId), offsets, false)
		if err != nil {
			log.Warn(ctx).Err(err).
				Uint64(log.KeySubscriptionID, subInfo.SubscriptionId).
				Msg("heartbeat commit offset error")
		}
	}
	return nil
}

func (ctrl *controller) RegisterTriggerWorker(
	ctx context.Context, request *ctrlpb.RegisterTriggerWorkerRequest,
) (*ctrlpb.RegisterTriggerWorkerResponse, error) {
	log.Info(ctx).
		Str(log.KeyTriggerWorkerAddr, request.Address).
		Msg("register trigger worker")
	err := ctrl.workerManager.AddTriggerWorker(ctx, request.Address)
	if err != nil {
		log.Warn(ctx).Err(err).
			Str(log.KeyTriggerWorkerAddr, request.Address).
			Msg("register trigger worker error")
		return nil, err
	}
	return &ctrlpb.RegisterTriggerWorkerResponse{}, nil
}

func (ctrl *controller) UnregisterTriggerWorker(
	ctx context.Context, request *ctrlpb.UnregisterTriggerWorkerRequest,
) (*ctrlpb.UnregisterTriggerWorkerResponse, error) {
	log.Info(ctx).
		Str(log.KeyTriggerWorkerAddr, request.Address).
		Msg("unregister trigger worker")

	ctrl.workerManager.RemoveTriggerWorker(context.TODO(), request.Address)
	return &ctrlpb.UnregisterTriggerWorkerResponse{}, nil
}

func (ctrl *controller) ListSubscription(
	ctx context.Context, request *ctrlpb.ListSubscriptionRequest,
) (*ctrlpb.ListSubscriptionResponse, error) {
	subscriptions := ctrl.subscriptionManager.ListSubscription(ctx)
	list := make([]*metapb.Subscription, 0, len(subscriptions))
	for _, sub := range subscriptions {
		if request.NamespaceId != 0 && request.NamespaceId != sub.NamespaceID.Uint64() {
			continue
		}
		if request.EventbusId != 0 && request.EventbusId != sub.EventbusID.Uint64() {
			continue
		}
		if request.Name != "" && !strings.Contains(sub.Name, request.Name) {
			continue
		}
		offsets, _ := ctrl.subscriptionManager.GetOffset(ctx, sub.ID)
		list = append(list, convert.ToPbSubscription(sub, offsets))
	}
	return &ctrlpb.ListSubscriptionResponse{Subscription: list}, nil
}

// gcSubscription before delete subscription,need
//
// 1.trigger worker remove subscription
// 2.delete offset
// 3.delete subscription .
func (ctrl *controller) gcSubscription(ctx context.Context, id vanus.ID, addr string) error {
	tWorker := ctrl.workerManager.GetTriggerWorker(addr)
	if tWorker != nil {
		err := tWorker.UnAssignSubscription(id)
		if err != nil {
			return err
		}
	}
	err := ctrl.subscriptionManager.DeleteSubscription(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (ctrl *controller) gcSubscriptions(ctx context.Context) {
	util.UntilWithContext(ctx, func(ctx context.Context) {
		ctrl.lock.Lock()
		defer ctrl.lock.Unlock()
		for ID, addr := range ctrl.needCleanSubscription {
			err := ctrl.gcSubscription(ctx, ID, addr)
			if err == nil {
				delete(ctrl.needCleanSubscription, ID)
			}
		}
	}, defaultGcSubscriptionInterval)
}

func (ctrl *controller) requeueSubscription(ctx context.Context, id vanus.ID, addr string) error {
	sub := ctrl.subscriptionManager.GetSubscription(ctx, id)
	if sub == nil {
		return nil
	}
	if sub.TriggerWorker != addr {
		// data is not consistent, record
		log.Error(ctx).
			Str(log.KeyTriggerWorkerAddr, sub.TriggerWorker).
			Str("runningAddr", addr).
			Msg("requeue subscription invalid")
	}
	switch sub.Phase {
	case metadata.SubscriptionPhaseCreated, metadata.SubscriptionPhaseRunning, metadata.SubscriptionPhasePending:
		sub.TriggerWorker = ""
		sub.Phase = metadata.SubscriptionPhasePending
		err := ctrl.subscriptionManager.UpdateSubscription(ctx, sub)
		if err != nil {
			return err
		}
	}
	metrics.CtrlTriggerGauge.WithLabelValues(sub.TriggerWorker).Dec()
	ctrl.scheduler.EnqueueSubscription(id)
	return nil
}

func (ctrl *controller) init(ctx context.Context) error {
	ctrl.initTriggerSystemEventbus()
	err := ctrl.subscriptionManager.Init(ctx)
	if err != nil {
		return err
	}
	err = ctrl.workerManager.Init(ctx)
	if err != nil {
		return err
	}
	// restart,need reschedule
	for _, sub := range ctrl.subscriptionManager.ListSubscription(ctx) {
		switch sub.Phase {
		case metadata.SubscriptionPhaseCreated:
			ctrl.scheduler.EnqueueNormalSubscription(sub.ID)
		case metadata.SubscriptionPhasePending, metadata.SubscriptionPhaseStopping:
			ctrl.scheduler.EnqueueSubscription(sub.ID)
		case metadata.SubscriptionPhaseToDelete:
			ctrl.needCleanSubscription[sub.ID] = sub.TriggerWorker
		}
	}
	return nil
}

func (ctrl *controller) membershipChangedProcessor(
	ctx context.Context, event member.MembershipChangedEvent,
) error {
	log.Info(ctx).
		Interface("event", event).
		Msg("start to process membership change event")
	ctrl.membershipMutex.Lock()
	defer ctrl.membershipMutex.Unlock()
	switch event.Type {
	case member.EventBecomeLeader:
		if ctrl.isLeader {
			return nil
		}
		log.Info().Msg("trigger become leader")
		err := ctrl.init(ctx)
		if err != nil {
			_err := ctrl.stop(ctx)
			if _err != nil {
				log.Error(ctx).Err(err).Msg("controller stop has error")
			}
			log.Error(ctx).Err(err).Msg("controller init has error")
			return err
		}
		ctrl.workerManager.Start()
		ctrl.subscriptionManager.Start()
		ctrl.scheduler.Run()
		go ctrl.gcSubscriptions(ctx)
		ctrl.state = primitive.ServerStateRunning
		ctrl.isLeader = true
	case member.EventBecomeFollower:
		if !ctrl.isLeader {
			return nil
		}
		log.Info().Msg("become flower")
		_err := ctrl.stop(ctx)
		if _err != nil {
			log.Error(ctx).Err(_err).Msg("controller stop has error")
		}
	}
	return nil
}

func (ctrl *controller) stop(_ context.Context) error {
	ctrl.member.ResignIfLeader()
	ctrl.state = primitive.ServerStateStopping
	ctrl.stopFunc()
	ctrl.scheduler.Stop()
	ctrl.workerManager.Stop()
	ctrl.subscriptionManager.Stop()
	ctrl.storage.Close()
	ctrl.state = primitive.ServerStateStopped
	return nil
}

func (ctrl *controller) Start() error {
	s, err := storage.NewStorage(ctrl.config.Storage)
	if err != nil {
		return err
	}
	ctrl.storage = s
	secretStorage, err := storage.NewSecretStorage(ctrl.config.Storage, ctrl.config.SecretEncryptionSalt)
	if err != nil {
		return err
	}
	ctrl.secretStorage = secretStorage
	ctrl.subscriptionManager = subscription.NewSubscriptionManager(ctrl.storage, ctrl.secretStorage,
		ctrl.ebClient, ctrl.cl)
	ctrl.workerManager = worker.NewTriggerWorkerManager(worker.Config{}, ctrl.storage,
		ctrl.subscriptionManager, ctrl.requeueSubscription)
	ctrl.scheduler = worker.NewSubscriptionScheduler(ctrl.workerManager, ctrl.subscriptionManager)

	ctrl.member.RegisterMembershipChangedProcessor(ctrl.membershipChangedProcessor)
	return nil
}

func (ctrl *controller) Stop(ctx context.Context) {
	if err := ctrl.stop(ctx); err != nil {
		log.Warn(ctx).Err(err).Msg("stop trigger controller error")
	}
}

func (ctrl *controller) initTriggerSystemEventbus() {
	// avoid blocking starting
	go func() {
		ctx := context.Background()
		log.Info(ctx).Msg("trigger controller is ready to check system eventbus")
		if err := ctrl.cl.WaitForControllerReady(true); err != nil {
			log.Error(ctx).Err(err).
				Msg("trigger controller try to create system eventbus, but Vanus cluster hasn't ready, exit")
			os.Exit(-1)
		}

		if _, err := ctrl.cl.EventbusService().CreateSystemEventbusIfNotExist(ctx, primitive.RetryEventbusName,
			"System Eventbus For Trigger Service"); err != nil {
			log.Error(ctx).Err(err).Msg("failed to create RetryEventbus, exit")
			os.Exit(-1)
		}
		log.Info(ctx).Msg("trigger controller has finished for checking system eventbus")
	}()
}
