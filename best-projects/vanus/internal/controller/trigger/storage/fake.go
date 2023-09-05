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

package storage

import (
	"context"

	"github.com/vanus-labs/vanus/internal/controller/trigger/metadata"
	"github.com/vanus-labs/vanus/internal/kv"
	pInfo "github.com/vanus-labs/vanus/internal/primitive/info"
	"github.com/vanus-labs/vanus/internal/primitive/vanus"
)

type fake struct {
	subs     map[vanus.ID]*metadata.Subscription
	offset   map[vanus.ID]map[vanus.ID]pInfo.OffsetInfo
	tWorkers map[string]*metadata.TriggerWorkerInfo
}

func NewFakeStorage() Storage {
	s := &fake{
		subs:     map[vanus.ID]*metadata.Subscription{},
		offset:   map[vanus.ID]map[vanus.ID]pInfo.OffsetInfo{},
		tWorkers: map[string]*metadata.TriggerWorkerInfo{},
	}
	return s
}

func (f *fake) Close() {

}

func (f *fake) CreateSubscription(_ context.Context, sub *metadata.Subscription) error {
	f.subs[sub.ID] = sub
	return nil
}

func (f *fake) UpdateSubscription(_ context.Context, sub *metadata.Subscription) error {
	f.subs[sub.ID] = sub
	return nil
}

func (f *fake) DeleteSubscription(_ context.Context, id vanus.ID) error {
	delete(f.subs, id)
	return nil
}

func (f *fake) GetSubscription(_ context.Context, id vanus.ID) (*metadata.Subscription, error) {
	return f.subs[id], nil
}

func (f *fake) ListSubscription(_ context.Context) ([]*metadata.Subscription, error) {
	list := make([]*metadata.Subscription, 0)
	for _, sub := range f.subs {
		list = append(list, sub)
	}
	return list, nil
}

func (f *fake) CreateOffset(_ context.Context, subscriptionID vanus.ID, info pInfo.OffsetInfo) error {
	sub, exist := f.offset[subscriptionID]
	if !exist {
		sub = map[vanus.ID]pInfo.OffsetInfo{}
		f.offset[subscriptionID] = sub
	}
	sub[info.EventlogID] = info
	return nil
}
func (f *fake) UpdateOffset(_ context.Context, subscriptionID vanus.ID, info pInfo.OffsetInfo) error {
	sub, exist := f.offset[subscriptionID]
	if !exist {
		return kv.ErrKeyNotFound
	}
	sub[info.EventlogID] = info
	return nil
}
func (f *fake) GetOffsets(_ context.Context, subscriptionID vanus.ID) (pInfo.ListOffsetInfo, error) {
	sub, exist := f.offset[subscriptionID]
	if !exist {
		return nil, nil
	}
	var infos pInfo.ListOffsetInfo
	for _, v := range sub {
		infos = append(infos, v)
	}
	return infos, nil
}

func (f *fake) DeleteOffset(_ context.Context, subscriptionID vanus.ID) error {
	delete(f.offset, subscriptionID)
	return nil
}

func (f *fake) SaveTriggerWorker(_ context.Context, info metadata.TriggerWorkerInfo) error {
	f.tWorkers[info.ID] = &info
	return nil
}
func (f *fake) GetTriggerWorker(_ context.Context, id string) (*metadata.TriggerWorkerInfo, error) {
	return f.tWorkers[id], nil
}
func (f *fake) DeleteTriggerWorker(_ context.Context, id string) error {
	delete(f.tWorkers, id)
	return nil
}
func (f *fake) ListTriggerWorker(_ context.Context) ([]*metadata.TriggerWorkerInfo, error) {
	list := make([]*metadata.TriggerWorkerInfo, 0)
	for _, data := range f.tWorkers {
		list = append(list, data)
	}
	return list, nil
}
