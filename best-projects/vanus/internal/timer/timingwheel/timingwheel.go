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

package timingwheel

import (
	"container/list"
	"context"
	"encoding/json"
	stderr "errors"
	"fmt"
	"io"
	"sync"
	"time"

	ce "github.com/cloudevents/sdk-go/v2"
	"github.com/vanus-labs/vanus/client"
	"github.com/vanus-labs/vanus/client/pkg/api"
	"github.com/vanus-labs/vanus/internal/kv"
	"github.com/vanus-labs/vanus/internal/kv/etcd"
	"github.com/vanus-labs/vanus/internal/primitive/vanus"
	"github.com/vanus-labs/vanus/internal/timer/metadata"
	"github.com/vanus-labs/vanus/observability/log"
	"github.com/vanus-labs/vanus/observability/metrics"
	"github.com/vanus-labs/vanus/pkg/cluster"
	"github.com/vanus-labs/vanus/pkg/errors"
	ctrlpb "github.com/vanus-labs/vanus/proto/pkg/controller"
	"google.golang.org/grpc/credentials/insecure"
	"k8s.io/apimachinery/pkg/util/wait"
)

const (
	// check waiting period every 1/defaultCheckWaitingPeriodRatio tick time by default.
	defaultCheckWaitingPeriodRatio = 10

	// frequent check waiting period every 1/defaultFrequentCheckWaitingPeriodRatio tick time by default.
	defaultFrequentCheckWaitingPeriodRatio = 100

	// number of tick flow in advance by default.
	defaultNumberOfTickFlowInAdvance = 1

	// number of events read each time by default.
	defaultNumberOfEventsRead = 10

	// the max number of workers by default.
	defaultMaxNumberOfWorkers = 1000

	recycleInterval = 60 * time.Second
)

var newEtcdClientV3 = etcd.NewEtcdClientV3

type Manager interface {
	Init(ctx context.Context) error
	Start(ctx context.Context) error
	Push(ctx context.Context, e *ce.Event) bool
	SetLeader(isleader bool)
	IsLeader() bool
	IsDeployed(ctx context.Context) bool
	Recover(ctx context.Context) error
	StopNotify() <-chan struct{}
	Stop(ctx context.Context)
}

// timingWheel timewheel contains multiple layers.
type timingWheel struct {
	config  *Config
	kvStore kv.Client
	ctrlCli ctrlpb.EventbusControllerClient
	client  client.Client
	cache   sync.Map
	twList  *list.List // element: *timingWheelElement

	ctrl cluster.Cluster

	receivingStation    *bucket
	distributionStation *bucket

	leader bool
	exitC  chan struct{}
	wg     sync.WaitGroup
}

func NewTimingWheel(c *Config) Manager {
	log.Info().
		Dur("tick", c.Tick).
		Int64("layers", c.Layers).
		Int64("wheel_size", c.WheelSize).
		Str("key_prefix", c.KeyPrefix).
		Strs("etcd_endpoints", c.EtcdEndpoints).
		Strs("ctrl_endpoints", c.CtrlEndpoints).
		Msg("new timingwheel manager")

	metrics.TimingWheelTickGauge.Set(float64(c.Tick))
	metrics.TimingWheelSizeGauge.Set(float64(c.WheelSize))
	metrics.TimingWheelLayersGauge.Set(float64(c.Layers))
	return &timingWheel{
		config: c,
		client: client.Connect(c.CtrlEndpoints),
		twList: list.New(),
		leader: false,
		exitC:  make(chan struct{}),
	}
}

// Init the current timing wheel.
func (tw *timingWheel) Init(ctx context.Context) error {
	log.Info(ctx).Msg("init timingwheel")
	// Init Hierarchical Timing Wheels.
	ctrl := cluster.NewClusterController(tw.config.CtrlEndpoints, insecure.NewCredentials())
	if err := ctrl.WaitForControllerReady(true); err != nil {
		panic("wait for controller ready timeout")
	}
	tw.ctrl = ctrl
	tw.ctrlCli = ctrl.EventbusService().RawClient()

	store, err := newEtcdClientV3(tw.config.EtcdEndpoints, tw.config.KeyPrefix)
	if err != nil {
		log.Error().Err(err).
			Str("key_prefix", tw.config.KeyPrefix).
			Strs("endpoints", tw.config.EtcdEndpoints).
			Msg("new etcd client v3 failed")
		panic("new etcd client failed")
	}
	tw.kvStore = store

	for layer := int64(1); layer <= tw.config.Layers+1; layer++ {
		tick := exponent(tw.config.Tick, tw.config.WheelSize, layer-1)
		twe := newTimingWheelElement(tw, tick, layer)
		twe.setElement(tw.twList.PushBack(twe))
		if layer <= tw.config.Layers {
			buckets := make(map[int64]*bucket, tw.config.WheelSize+defaultNumberOfTickFlowInAdvance)
			for i := int64(0); i < tw.config.WheelSize+defaultNumberOfTickFlowInAdvance; i++ {
				ebName := fmt.Sprintf(timerBuiltInEventbus, layer, i)
				buckets[i] = newBucket(tw, twe.element, tick, ebName, layer, i)
			}
			twe.buckets = buckets
		} else {
			twe.buckets = make(map[int64]*bucket)
		}
	}
	tw.receivingStation = newBucket(tw, nil, 0, timerBuiltInEventbusReceivingStation, 0, 0)
	tw.distributionStation = newBucket(tw, nil, 0, timerBuiltInEventbusDistributionStation, 0, 0)

	return nil
}

// Start starts the current timing wheel.
func (tw *timingWheel) Start(ctx context.Context) error {
	var err error
	log.Info(ctx).
		Bool("leader", tw.leader).
		Msg("start timingwheel")

	// here is to wait for the leader to complete the creation of all eventbus
	waitCtx, cancel := context.WithCancel(ctx)
	wait.Until(func() {
		if tw.IsLeader() || tw.IsDeployed(ctx) {
			cancel()
		} else {
			log.Info(ctx).Msg("wait for the leader to be ready")
		}
	}, time.Second, waitCtx.Done())

	// start distribution station for scheduled events distributing
	if err = tw.startDistributionStation(ctx); err != nil {
		return err
	}

	// start all bucket of each layer
	for e := tw.twList.Front(); e != nil; e = e.Next() {
		for _, bucket := range e.Value.(*timingWheelElement).getBuckets() {
			if err = bucket.start(ctx); err != nil {
				log.Error(ctx).Err(err).
					Str("eventbus", bucket.getEventbus()).
					Msg("start bucket failed")
				return err
			}
		}
	}

	// start receiving station for scheduled events receiving
	if err = tw.startReceivingStation(ctx); err != nil {
		return err
	}

	// start bucket recycling
	tw.startRecycling(ctx)

	return nil
}

func (tw *timingWheel) StopNotify() <-chan struct{} {
	return tw.exitC
}

// Stop stops the current timing wheel.
func (tw *timingWheel) Stop(ctx context.Context) {
	log.Info(ctx).Msg("stop timingwheel")
	// wait for all goroutine to end
	for e := tw.twList.Front(); e != nil; e = e.Next() {
		for _, bucket := range e.Value.(*timingWheelElement).getBuckets() {
			bucket.stop(ctx)
		}
		e.Value.(*timingWheelElement).wait(ctx)
	}
	tw.receivingStation.wait(ctx)
	tw.distributionStation.wait(ctx)
	close(tw.exitC)
	tw.wg.Wait()
	if closer, ok := tw.ctrlCli.(io.Closer); ok {
		_ = closer.Close()
	}
}

func (tw *timingWheel) SetLeader(isLeader bool) {
	tw.leader = isLeader
}

func (tw *timingWheel) IsLeader() bool {
	return tw.leader
}

func (tw *timingWheel) IsDeployed(ctx context.Context) bool {
	var err error
	if _, err = tw.ctrl.EventbusService().GetSystemEventbusByName(ctx, tw.receivingStation.eventbus); err != nil {
		return false
	}
	if _, err = tw.ctrl.EventbusService().GetSystemEventbusByName(ctx, tw.distributionStation.eventbus); err != nil {
		return false
	}
	return true
}

func (tw *timingWheel) Recover(ctx context.Context) error {
	offsetPairs, err := tw.kvStore.List(ctx, metadata.OffsetKeyPrefixInKVStore)
	if err != nil {
		return err
	}
	// no offset metadata, no recovery required
	if len(offsetPairs) == 0 {
		return nil
	}
	offsetMetaMap := make(map[string]*metadata.OffsetMeta, tw.config.Layers+1)
	for _, v := range offsetPairs {
		md := &metadata.OffsetMeta{}
		_ = json.Unmarshal(v.Value, md)
		if md.Layer > tw.config.Layers &&
			tw.twList.Back().Value.(*timingWheelElement).makeSureBucketExist(ctx, md.Slot) != nil {
			return err
		}
		offsetMetaMap[md.Eventbus] = md
	}

	for e := tw.twList.Front(); e != nil; e = e.Next() {
		for _, bucket := range e.Value.(*timingWheelElement).getBuckets() {
			if v, ok := offsetMetaMap[bucket.getEventbus()]; ok {
				log.Info(ctx).
					Int64("layer", v.Layer).
					Int64("slot", v.Slot).
					Int64("offset", v.Offset).
					Str("eventbus", v.Eventbus).
					Msg("recover offset metadata")
				bucket.offset = v.Offset
			}
		}
	}

	if _, ok := offsetMetaMap[timerBuiltInEventbusReceivingStation]; ok {
		log.Info(ctx).
			Int64("offset", offsetMetaMap[timerBuiltInEventbusReceivingStation].Offset).
			Str("eventbus", tw.receivingStation.getEventbus()).
			Msg("recover receiving station metadata")
		tw.receivingStation.offset = offsetMetaMap[timerBuiltInEventbusReceivingStation].Offset
	}
	if _, ok := offsetMetaMap[timerBuiltInEventbusDistributionStation]; ok {
		log.Info(ctx).
			Int64("offset", offsetMetaMap[timerBuiltInEventbusReceivingStation].Offset).
			Str("eventbus", tw.receivingStation.getEventbus()).
			Msg("recover distribution station metadata")
		tw.distributionStation.offset = offsetMetaMap[timerBuiltInEventbusDistributionStation].Offset
	}

	return nil
}

// Push the scheduled event to the timingwheel.
func (tw *timingWheel) Push(ctx context.Context, e *ce.Event) bool {
	tm := newTimingMsg(ctx, e)
	log.Info(ctx).
		Str("event_id", e.ID()).
		Str("expiration", tm.getExpiration().Format(time.RFC3339Nano)).
		Msg("push event to timingwheel")

	metrics.TimerScheduledEventDelayTime.WithLabelValues(metrics.LabelScheduledEventDelayTime).
		Observe(time.Until(tm.getExpiration()).Seconds())

	if tm.hasExpired() {
		// Already expired
		return tw.getDistributionStation().push(ctx, tm)
	}

	return tw.twList.Front().Value.(*timingWheelElement).pushHandler(ctx, tm)
}

func (tw *timingWheel) getReceivingStation() *bucket {
	return tw.receivingStation
}

func (tw *timingWheel) getDistributionStation() *bucket {
	return tw.distributionStation
}

func (tw *timingWheel) startRecycling(ctx context.Context) {
	tw.wg.Add(1)
	go func() {
		defer tw.wg.Done()
		ticker := time.NewTicker(recycleInterval)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				log.Debug(ctx).Msg("context canceled at timingwheel recycling")
				return
			case <-ticker.C:
				if !tw.IsLeader() {
					break
				}
				tw.twList.Back().Value.(*timingWheelElement).recycling(ctx)
			}
		}
	}()
}

func (tw *timingWheel) startReceivingStation(ctx context.Context) error {
	var err error
	if err = tw.getReceivingStation().createEventbus(ctx); err != nil {
		return err
	}

	tw.getReceivingStation().connectEventbus(ctx)
	tw.runReceivingStation(ctx)
	return nil
}

const receiveGoroutineNum = 2

// runReceivingStation as the unified entrance of scheduled events and pushed to the timingwheel.
func (tw *timingWheel) runReceivingStation(ctx context.Context) {
	offsetC := make(chan waitGroup, defaultMaxNumberOfWorkers)
	tw.wg.Add(receiveGoroutineNum)
	// update offset asynchronously
	go func() {
		defer tw.wg.Done()
		for {
			select {
			case <-ctx.Done():
				log.Debug(ctx).Msg("context canceled at receiving station update offset metadata")
				return
			case offset := <-offsetC:
				// wait for all goroutines to finish before updating offset metadata
				offset.wg.Wait()
				log.Debug(ctx).
					Str("eventbus", tw.receivingStation.getEventbus()).
					Int64("update_to", offset.data).
					Msg("update offset metadata")
				tw.receivingStation.updateOffsetMeta(ctx, offset.data)
			}
		}
	}()
	go func() {
		defer tw.wg.Done()
		// limit the number of goroutines to no more than defaultMaxNumberOfWorkers
		glimitC := make(chan struct{}, defaultMaxNumberOfWorkers)
		for {
			select {
			case <-ctx.Done():
				log.Debug(ctx).Msg("context canceled at receiving station running")
				return
			default:
				// batch read
				events, err := tw.receivingStation.getEvent(ctx, defaultNumberOfEventsRead)
				if err != nil {
					if !errors.Is(err, errors.ErrOffsetOnEnd) {
						log.Error(ctx).Err(err).
							Str("eventbus", tw.receivingStation.getEventbus()).
							Msg("get event failed when receiving station running")
					}
					time.Sleep(sleepDuration)
					break
				}
				if len(events) == 0 {
					time.Sleep(sleepDuration)
					log.Info(ctx).
						Str("function", "runReceivingStation").
						Msg("no more message")
					break
				}
				// concurrent write
				numberOfEvents := int64(len(events))
				log.Debug(ctx).
					Str("eventbus", tw.receivingStation.getEventbus()).
					Int64("offset", tw.receivingStation.getOffset()).
					Int64("number_of_events", numberOfEvents).
					Msg("got events when receiving station running")

				wg := sync.WaitGroup{}
				for _, event := range events {
					// TODO(jiangkai): check event dst eventbus is vaild.
					wg.Add(1)
					glimitC <- struct{}{}
					go func(ctx context.Context, e *ce.Event) {
						defer wg.Done()
						waitCtx, cancel := context.WithCancel(ctx)
						wait.Until(func() {
							startTime := time.Now()
							if tw.Push(ctx, e) {
								metrics.TimerPushEventTime.WithLabelValues(
									metrics.LabelTimerPushScheduledEventTime).
									Observe(time.Since(startTime).Seconds())
								metrics.TimerPushEventTPSCounterVec.WithLabelValues(metrics.LabelTimer).Inc()
								cancel()
							} else {
								log.Warn(ctx).
									Str("event_id", e.ID()).
									Interface("eventbus", e.Extensions()[xVanusEventbus]).
									Str("delivery_time", newTimingMsg(ctx, e).getExpiration().Format(time.RFC3339Nano)).
									Msg("push event to timingwheel failed, retry until it succeed")
							}
						}, tw.config.Tick/defaultCheckWaitingPeriodRatio, waitCtx.Done())
						<-glimitC
					}(ctx, event)
				}
				// asynchronously update offset after the same batch of events are successfully written
				offsetC <- waitGroup{
					wg:   &wg,
					data: tw.receivingStation.getOffset() + numberOfEvents,
				}
				tw.receivingStation.incOffset(numberOfEvents)
			}
		}
	}()
}

func (tw *timingWheel) startDistributionStation(ctx context.Context) error {
	var err error
	if err = tw.getDistributionStation().createEventbus(ctx); err != nil {
		return err
	}

	tw.getDistributionStation().connectEventbus(ctx)
	tw.runDistributionStation(ctx)
	return nil
}

// runDistributionStation as the unified exit of scheduled events and popped to the timingwheel.
func (tw *timingWheel) runDistributionStation(ctx context.Context) {
	offsetC := make(chan waitGroup, defaultMaxNumberOfWorkers)
	tw.wg.Add(1)
	// update offset asynchronously
	go func() {
		defer tw.wg.Done()
		for {
			select {
			case <-ctx.Done():
				log.Debug(ctx).Msg("context canceled at distribution station update offset metadata")
				return
			case offset := <-offsetC:
				// wait for all goroutines to finish before updating offset metadata
				offset.wg.Wait()
				log.Debug(ctx).
					Str("eventbus", tw.distributionStation.getEventbus()).
					Int64("update_to", offset.data).
					Msg("update offset metadata")
				tw.distributionStation.updateOffsetMeta(ctx, offset.data)
			}
		}
	}()

	tw.wg.Add(1)
	go func() {
		defer tw.wg.Done()
		// limit the number of goroutines to no more than defaultMaxNumberOfWorkers
		glimitC := make(chan struct{}, defaultMaxNumberOfWorkers)
		for {
			select {
			case <-ctx.Done():
				log.Debug(ctx).Msg("context canceled at distribution station running")
				return
			default:
				// batch read
				events, err := tw.distributionStation.getEvent(ctx, defaultNumberOfEventsRead)
				if err != nil {
					if !errors.Is(err, errors.ErrOffsetOnEnd) {
						log.Error(ctx).Err(err).
							Str("eventbus", tw.distributionStation.getEventbus()).
							Msg("get event failed when distribution station running")
					}
					time.Sleep(sleepDuration)
					break
				}
				if len(events) == 0 {
					time.Sleep(sleepDuration)
					log.Debug(ctx).Str("function", "runDistributionStation").Msg("no more message")
					break
				}
				// concurrent write
				numberOfEvents := int64(len(events))
				log.Debug(ctx).
					Str("eventbus", tw.distributionStation.getEventbus()).
					Int64("offset", tw.distributionStation.getOffset()).
					Int64("number_of_events", numberOfEvents).
					Msg("got events when distribution station running")

				wg := sync.WaitGroup{}
				for _, event := range events {
					wg.Add(1)
					glimitC <- struct{}{}
					go func(ctx context.Context, e *ce.Event) {
						defer wg.Done()
						waitCtx, cancel := context.WithCancel(ctx)
						wait.Until(func() {
							startTime := time.Now()
							if err = tw.deliver(ctx, e); err == nil {
								metrics.TimerDeliverEventTime.WithLabelValues(
									metrics.LabelTimerDeliverScheduledEventTime).
									Observe(time.Since(startTime).Seconds())
								metrics.TimerDeliverEventTPSCounterVec.WithLabelValues(metrics.LabelTimer).Inc()
								cancel()
							} else {
								log.Warn(ctx).
									Str("event_id", e.ID()).
									Interface("eventbus", e.Extensions()[xVanusEventbus]).
									Time("delivery_time", newTimingMsg(ctx, e).getExpiration()).
									Msg("deliver event failed, retry until it succeed")
							}
						}, tw.config.Tick/defaultCheckWaitingPeriodRatio, waitCtx.Done())
						<-glimitC
					}(ctx, event)
				}
				// asynchronously update offset after the same batch of events are successfully written
				offsetC <- waitGroup{
					wg:   &wg,
					data: tw.distributionStation.getOffset() + numberOfEvents,
				}
				tw.distributionStation.incOffset(numberOfEvents)
			}
		}
	}()
}

func (tw *timingWheel) deliver(ctx context.Context, e *ce.Event) error {
	var (
		err  error
		ebID string
	)

	err = e.ExtensionAs(xVanusEventbus, &ebID)
	if err != nil {
		log.Error(ctx).Err(err).Msg("get eventbus failed when delivering")
		return err
	}
	eventbusID, err := vanus.NewIDFromString(ebID)
	if err != nil {
		log.Error(ctx).Err(err).
			Str("eventbus_id", ebID).
			Msg("eventbus id string to uint64 failed when delivering")
		return err
	}
	v, exist := tw.cache.Load(eventbusID)
	if !exist {
		v, _ = tw.cache.LoadOrStore(eventbusID, tw.client.Eventbus(ctx, api.WithID(eventbusID.Uint64())).Writer())
	}
	writer, _ := v.(api.BusWriter)
	_, err = api.AppendOne(ctx, writer, e)
	if err != nil {
		if errors.Is(err, errors.ErrOffsetOnEnd) {
			log.Warn(ctx).Err(err).
				Str("eventbus_id", ebID).
				Str("event_id", e.ID()).
				Interface("eventbus", e.Extensions()[xVanusEventbus]).
				Time("delivery_time", newTimingMsg(ctx, e).getExpiration()).
				Msg("eventbus not found, discard this event")
			return nil
		}
		log.Error(ctx).Err(err).Str("eventbus_id", ebID).Msg("append failed")
		return err
	}
	log.Debug(ctx).
		Str("event_id", e.ID()).
		Str("eventbus_id", ebID).
		Time("delivery_time", newTimingMsg(ctx, e).getExpiration()).
		Msg("event delivered")
	return nil
}

// timingWheelElement timingwheelelement has N number of buckets, every bucket is an eventbus.
type timingWheelElement struct {
	config   *Config
	kvStore  kv.Client
	ctrlCli  ctrlpb.EventbusControllerClient
	tick     time.Duration
	layer    int64
	interval time.Duration
	buckets  map[int64]*bucket

	exitC chan struct{}
	mu    sync.RWMutex
	wg    sync.WaitGroup

	timingwheel *timingWheel
	element     *list.Element

	pushHandler func(ctx context.Context, tm *timingMsg) bool
}

// newTimingWheel is an internal helper function that really creates an instance of TimingWheel.
func newTimingWheelElement(tw *timingWheel, tick time.Duration, layer int64) *timingWheelElement {
	if tick <= 0 {
		panic(stderr.New("tick must be greater than or equal to 1s"))
	}

	twe := &timingWheelElement{
		config:      tw.config,
		kvStore:     tw.kvStore,
		ctrlCli:     tw.ctrlCli,
		tick:        tick,
		layer:       layer,
		interval:    tick * time.Duration(tw.config.WheelSize),
		exitC:       make(chan struct{}),
		timingwheel: tw,
	}

	if layer > tw.config.Layers {
		twe.pushHandler = twe.pushBack
	} else {
		twe.pushHandler = twe.push
	}
	return twe
}

func (twe *timingWheelElement) push(ctx context.Context, tm *timingMsg) bool {
	if twe.allowPush(tm) {
		index := tm.getExpiration().UnixNano() % twe.interval.Nanoseconds() / twe.tick.Nanoseconds()
		// Put it into its own bucket
		return twe.buckets[index].push(ctx, tm)
	}
	// Out of the interval. Put it into the overflow wheel
	return twe.next().pushHandler(ctx, tm)
}

func (twe *timingWheelElement) pushBack(ctx context.Context, tm *timingMsg) bool {
	index := tm.getExpiration().UnixNano() / twe.tick.Nanoseconds()
	// Put it into its own bucket
	if twe.makeSureBucketExist(ctx, index) != nil {
		log.Error(ctx).
			Str("eventbus", twe.buckets[index].getEventbus()).
			Time("expiration", tm.getExpiration()).
			Msg("push timing message failed because bucket not exist")
		return false
	}
	return twe.buckets[index].push(ctx, tm)
}

func (twe *timingWheelElement) allowPush(tm *timingMsg) bool {
	now := time.Now()
	timeOfBufferBoundaryLine := now.UnixNano() - (now.UnixNano() % twe.tick.Nanoseconds()) + twe.interval.Nanoseconds()
	return tm.getExpiration().UnixNano() < timeOfBufferBoundaryLine
}

func (twe *timingWheelElement) flow(ctx context.Context, tm *timingMsg) bool {
	index := twe.calculateIndex(tm)
	// Put it into its own bucket
	return twe.buckets[index].push(ctx, tm)
}

func (twe *timingWheelElement) calculateIndex(tm *timingMsg) int64 {
	// the timing message comes from the timingwheel of the upper layer
	startTimeOfBucket := tm.getExpiration().UnixNano() -
		(tm.getExpiration().UnixNano() % twe.interval.Nanoseconds())
	timeOfEarlyFlow := defaultNumberOfTickFlowInAdvance * twe.tick.Nanoseconds()
	timeOfBufferBoundaryLine := startTimeOfBucket - timeOfEarlyFlow + twe.interval.Nanoseconds()
	if tm.getExpiration().UnixNano() >= timeOfBufferBoundaryLine {
		// Put it into its buffer bucket
		return (tm.getExpiration().UnixNano()-timeOfBufferBoundaryLine)/
			twe.tick.Nanoseconds() + twe.config.WheelSize
	}
	// Put it into its own bucket
	return tm.getExpiration().UnixNano() % twe.interval.Nanoseconds() / twe.tick.Nanoseconds()
}

func (twe *timingWheelElement) makeSureBucketExist(ctx context.Context, index int64) error {
	// TODO(jiangkai): redesign locks if here is a performance bottleneck in the future, by jiangkai, 2022.09.16
	// the segmented lock may solve the problem.
	twe.mu.RLock()
	if _, ok := twe.buckets[index]; ok {
		twe.mu.RUnlock()
		return nil
	}
	twe.mu.RUnlock()
	twe.mu.Lock()
	defer twe.mu.Unlock()
	if _, ok := twe.buckets[index]; ok {
		return nil
	}
	ebName := fmt.Sprintf(timerBuiltInEventbus, twe.layer, index)
	twe.buckets[index] = newBucket(twe.timingwheel, twe.element, twe.tick, ebName, twe.layer, index)
	if err := twe.buckets[index].start(ctx); err != nil {
		log.Error(ctx).Err(err).
			Str("eventbus", twe.buckets[index].getEventbus()).
			Msg("start bucket failed when makesure bucket exist")
		return err
	}
	exist, err := twe.buckets[index].existsOffsetMeta(ctx)
	if !exist && err == nil {
		twe.buckets[index].updateOffsetMeta(ctx, twe.buckets[index].offset)
	}
	return nil
}

func (twe *timingWheelElement) recycling(ctx context.Context) {
	twe.mu.Lock()
	defer twe.mu.Unlock()
	for idx, bucket := range twe.buckets {
		if time.Now().UnixNano()/bucket.tick.Nanoseconds() > idx && bucket.hasOnEnd(ctx) {
			log.Info(ctx).
				Str("bucket", bucket.eventbus).
				Msg("recycle expired bucket")
			bucket.stop(ctx)
			bucket.recycle(ctx)
			delete(twe.buckets, idx)
		}
	}
}

func (twe *timingWheelElement) wait(_ context.Context) {
	twe.wg.Wait()
}

func (twe *timingWheelElement) getBuckets() map[int64]*bucket {
	return twe.buckets
}

func (twe *timingWheelElement) setElement(element *list.Element) {
	twe.element = element
}

func (twe *timingWheelElement) prev() *timingWheelElement {
	return twe.element.Prev().Value.(*timingWheelElement)
}

func (twe *timingWheelElement) next() *timingWheelElement {
	return twe.element.Next().Value.(*timingWheelElement)
}
