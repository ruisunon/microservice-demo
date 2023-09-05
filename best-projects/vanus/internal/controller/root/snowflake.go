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

package root

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"path"
	"sync"
	"time"

	"github.com/vanus-labs/vanus/internal/controller/member"
	"github.com/vanus-labs/vanus/internal/kv"
	"github.com/vanus-labs/vanus/internal/kv/etcd"
	"github.com/vanus-labs/vanus/observability/log"
	ctrlpb "github.com/vanus-labs/vanus/proto/pkg/controller"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	spinInterval = 100 * time.Millisecond
)

type Config struct {
	KVEndpoints []string
	KVPrefix    string
}

func NewSnowflakeController(cfg Config, mem member.Member) *snowflake { //nolint:revive // it's ok
	sf := &snowflake{
		cfg:    cfg,
		member: mem,
		r:      rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	mem.RegisterMembershipChangedProcessor(sf.membershipChangedProcessor)
	return sf
}

var (
	_ ctrlpb.PingServerServer = &snowflake{}
)

type snowflake struct {
	startAt  time.Time
	cfg      Config
	kvStore  kv.Client
	isLeader bool
	member   member.Member
	nodes    map[uint16]*node
	mutex    sync.RWMutex
	r        *rand.Rand
}

func (sf *snowflake) Ping(_ context.Context, _ *emptypb.Empty) (*ctrlpb.PingResponse, error) {
	return &ctrlpb.PingResponse{
		LeaderAddr: sf.member.GetLeaderAddr(),
	}, nil
}

type node struct {
	StartAt time.Time
	ID      uint16
}

func (sf *snowflake) Start(_ context.Context) error {
	store, err := etcd.NewEtcdClientV3(sf.cfg.KVEndpoints, sf.cfg.KVPrefix)
	if err != nil {
		return err
	}

	sf.kvStore = store
	return nil
}

func (sf *snowflake) GetClusterStartTime(_ context.Context, _ *emptypb.Empty) (*timestamppb.Timestamp, error) {
	sf.mutex.Lock()
	defer sf.mutex.Unlock()

	for sf.member.GetLeaderID() == "" {
		time.Sleep(spinInterval)
	}

	if !sf.isLeader {
		return nil, errors.New("i'm not leader")
	}
	return timestamppb.New(sf.startAt), nil
}

func (sf *snowflake) RegisterNode(ctx context.Context, in *wrapperspb.UInt32Value) (*emptypb.Empty, error) {
	sf.mutex.Lock()
	defer sf.mutex.Unlock()

	for sf.member.GetLeaderID() == "" {
		time.Sleep(spinInterval)
	}

	if !sf.isLeader {
		return nil, errors.New("i'm not leader")
	}

	id := uint16(in.Value)
	// TODO(wenfeng) find a good solution in future
	// _, exist := sf.nodes[id]
	//
	// if exist {
	//	return nil, errors.New("node has been register")
	// }

	n := &node{
		ID:      id,
		StartAt: time.Now(),
	}
	sf.nodes[id] = n

	data, _ := json.Marshal(n)

	if err := sf.kvStore.Set(ctx, GetNodeIDKey(n.ID), data); err != nil {
		return nil, errors.New("save node to kv failed")
	}
	log.Info(ctx).
		Uint16("node_id", id).
		Msg("a new node registered")
	return &emptypb.Empty{}, nil
}

func (sf *snowflake) UnregisterNode(ctx context.Context, in *wrapperspb.UInt32Value) (*emptypb.Empty, error) {
	sf.mutex.Lock()
	defer sf.mutex.Unlock()

	if !sf.isLeader {
		return nil, errors.New("i'm not leader")
	}

	node, exist := sf.nodes[uint16(in.Value)]
	if !exist {
		return &emptypb.Empty{}, nil
	}

	delete(sf.nodes, uint16(in.Value))

	if err := sf.kvStore.Delete(ctx, GetNodeIDKey(node.ID)); err != nil {
		return nil, errors.New("delete node from kv failed")
	}

	log.Info(ctx).
		Uint16("node_id", node.ID).
		Msg("a node unregistered")
	return &emptypb.Empty{}, nil
}

func (sf *snowflake) Stop() {
	_ = sf.kvStore.Close()
}

func (sf *snowflake) membershipChangedProcessor(ctx context.Context, event member.MembershipChangedEvent) error {
	log.Info(ctx).
		Interface("event", event).
		Msg("start to process membership change event")
	start := time.Now()
	defer func() {
		log.Info(ctx).
			Dur("duration", time.Since(start)).
			Msg("processing membership change event is finished")
	}()

	sf.mutex.Lock()
	defer sf.mutex.Unlock()

	switch event.Type {
	case member.EventBecomeLeader:
		if sf.isLeader {
			return nil
		}

		exist, err := sf.kvStore.Exists(ctx, kv.ClusterStart)
		if err != nil {
			return err
		}

		if !exist {
			now := time.Now()
			data, _ := now.MarshalJSON()
			if err = sf.kvStore.Set(ctx, kv.ClusterStart, data); err != nil {
				return err
			}
		}

		val, err := sf.kvStore.Get(ctx, kv.ClusterStart)
		if err != nil {
			return err
		}

		startAt := time.Time{}
		if err = startAt.UnmarshalJSON(val); err != nil {
			return err
		}
		sf.startAt = startAt

		pairs, err := sf.kvStore.List(ctx, kv.ClusterNode)
		if err != nil {
			return err
		}

		sf.nodes = map[uint16]*node{}
		for _, v := range pairs {
			n := &node{}
			if err = json.Unmarshal(v.Value, n); err != nil {
				return err
			}
			sf.nodes[n.ID] = n
		}
		sf.isLeader = true
	case member.EventBecomeFollower:
		if !sf.isLeader {
			return nil
		}
		sf.isLeader = false
		sf.nodes = nil
	}
	return nil
}

func GetNodeIDKey(nodeID uint16) string {
	return path.Join(kv.ClusterNode, fmt.Sprintf("%d", nodeID))
}
