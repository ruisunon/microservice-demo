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

package metrics

import (
	"context"
	"fmt"
	"net/http"
	"regexp"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/vanus-labs/vanus/observability/log"
)

const (
	namespace = "vanus"
)

func Init(ctx context.Context, cfg Config, getCollectors func() []prometheus.Collector) {
	if !cfg.Enable {
		log.Info(ctx).Msg("metrics module has been disabled")
		return
	}
	if getCollectors == nil {
		log.Info(ctx).Msg("metrics module has been disabled due to empty collectors")
		return
	}
	colls := getCollectors()
	if len(colls) == 0 {
		log.Info(ctx).Msg("metrics module has been disabled due to empty collectors")
		return
	}

	reg := prometheus.NewRegistry()
	reg.MustRegister(colls...)
	http.Handle("/metrics", promhttp.HandlerFor(
		reg,
		promhttp.HandlerOpts{
			EnableOpenMetrics: false,
		},
	))
	go func() {
		if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.GetPort()), nil); err != nil {
			log.Error().Err(err).Msg("Metrics listen and serve failed.")
		}
	}()
	log.Info().Int("port", cfg.GetPort()).Msg("metrics module started")
}

func GetControllerMetrics() []prometheus.Collector {
	coll := []prometheus.Collector{
		ControllerLeaderGaugeVec,
		EventbusGauge,
		EventbusUpdatedGauge,
		EventbusDeletedGauge,
		EventlogGaugeVec,
		SegmentGaugeVec,
		SegmentSizeGaugeVec,
		SegmentCapacityGaugeVec,
		SegmentEventNumberGaugeVec,
		SegmentCreatedByCacheMissing,
		SegmentCreatedByScaleTask,
		SegmentDeletedCounterVec,
		SubscriptionGauge,
		SubscriptionTransformerGauge,
		CtrlTriggerGauge,
	}
	return append(coll, getGoRuntimeMetrics()...)
}

func GetGatewayMetrics() []prometheus.Collector {
	coll := []prometheus.Collector{
		GatewayEventReceivedCountVec,
		// GatewayEventWriteLatencyHistogramVec,
		GatewayEventWriteLatencySummaryVec,
	}
	return append(coll, getGoRuntimeMetrics()...)
}

func GetTriggerMetrics() []prometheus.Collector {
	coll := []prometheus.Collector{
		TriggerGauge,
		TriggerPullEventCounter,
		TriggerFilterCostSecond,
		TriggerTransformCostSecond,
		TriggerFilterMatchEventCounter,
		TriggerFilterMatchRetryEventCounter,
		TriggerRetryEventCounter,
		TriggerRetryEventAppendSecond,
		TriggerDeadLetterEventCounter,
		TriggerDeadLetterEventAppendSecond,
		TriggerPushEventCounter,
		TriggerPushEventTime,
	}
	return append(coll, getGoRuntimeMetrics()...)
}

func GetTimerMetrics() []prometheus.Collector {
	coll := []prometheus.Collector{
		TimingWheelTickGauge,
		TimingWheelSizeGauge,
		TimingWheelLayersGauge,
		TimerPushEventTPSCounterVec,
		TimerDeliverEventTPSCounterVec,
		TimerScheduledEventDelayTime,
		TimerPushEventTime,
		TimerDeliverEventTime,
	}
	return append(coll, getGoRuntimeMetrics()...)
}

func GetSegmentServerMetrics() []prometheus.Collector {
	coll := []prometheus.Collector{
		WriteThroughputCounterVec,
		WriteTPSCounterVec,
		ReadTPSCounterVec,
		ReadThroughputCounterVec,
	}
	return append(coll, getGoRuntimeMetrics()...)
}

func getGoRuntimeMetrics() []prometheus.Collector {
	return []prometheus.Collector{
		collectors.NewBuildInfoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
		collectors.NewGoCollector(collectors.WithGoCollectorRuntimeMetrics(
			collectors.GoRuntimeMetricsRule{Matcher: regexp.MustCompile("/.*")},
		)),
	}
}
