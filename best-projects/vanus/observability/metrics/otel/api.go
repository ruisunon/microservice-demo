// Copyright 2023 Linkall Inc.
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

package otel

import (
	"context"
	"sync"

	"go.opentelemetry.io/otel/attribute"
)

var (
	countMap          = make(map[string]ICounter, 0)
	gaugeMap          = make(map[string]IGauge, 0)
	histogramMap      = make(map[string]IHistogram, 0)
	metricCreateMutex = sync.Mutex{}
	emptyCount        = &promCounter{}
	emptyGauge        = &promGauge{}
	emptyHistogram    = &promHistogram{}
)

type ICounter interface {
	IncrInt(int64, ...attribute.KeyValue)
	IncrFloat(float64, ...attribute.KeyValue)
	Async(func(context.Context, ICounter))
}

type unit string

const (
	UnitMillisecond   = unit("ms")
	UnitByte          = unit("byte")
	UnitDimensionless = unit("1")
)

type metricKey struct {
	name        string
	description string
	unit        unit
}

func NewMetricKey(name string, u unit, desc string) *metricKey {
	return &metricKey{
		name:        name,
		description: desc,
		unit:        u,
	}
}

func newCounter(k *metricKey) ICounter {
	return emptyCount
}

type IGauge interface {
	IncrInt(int64, ...attribute.KeyValue)
	IncrFloat(float64, ...attribute.KeyValue)
	Async(func(context.Context, IGauge))
}

func newGauge(k *metricKey) IGauge {
	return emptyGauge
}

type IHistogram interface {
	RecordInt(int64, ...attribute.KeyValue)
	RecordFloat(float64, ...attribute.KeyValue)
	Async(func(context.Context, IHistogram))
}

func newHistogram(k *metricKey) IHistogram {
	return emptyHistogram
}

func GetCounter(key *metricKey) ICounter {
	if !isValidKey(key) {
		return emptyCount
	}
	v, exist := countMap[key.name]
	if !exist {
		metricCreateMutex.Lock()
		v, exist = countMap[key.name]
		if !exist {
			v = newCounter(key)
			countMap[key.name] = v
		}
		metricCreateMutex.Unlock()
	}
	return v
}

func GetGauge(key *metricKey) IGauge {
	if !isValidKey(key) {
		return emptyGauge
	}
	v, exist := gaugeMap[key.name]
	if !exist {
		metricCreateMutex.Lock()
		v, exist = gaugeMap[key.name]
		if !exist {
			v = newGauge(key)
			gaugeMap[key.name] = v
		}
		metricCreateMutex.Unlock()
	}
	return v
}

func GetHistogram(key *metricKey) IHistogram {
	if !isValidKey(key) {
		return emptyHistogram
	}
	v, exist := histogramMap[key.name]
	if !exist {
		metricCreateMutex.Lock()
		v, exist = histogramMap[key.name]
		if !exist {
			v = newHistogram(key)
			histogramMap[key.name] = v
		}
		metricCreateMutex.Unlock()
	}
	return v
}

func isValidKey(k *metricKey) bool {
	if k == nil {
		return false
	}
	if k.name == "" {
		return false
	}
	if k.unit == "" {
		return false
	}
	return true
}
