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

package main

import (
	"context"
	"flag"
	"os"

	"github.com/vanus-labs/vanus/observability"
	"github.com/vanus-labs/vanus/observability/log"
	"github.com/vanus-labs/vanus/observability/metrics"
	"github.com/vanus-labs/vanus/pkg/util/signal"

	"github.com/vanus-labs/vanus/internal/timer"
	"github.com/vanus-labs/vanus/internal/timer/leaderelection"
	"github.com/vanus-labs/vanus/internal/timer/timingwheel"
)

var configPath = flag.String("config", "./config/timer.yaml", "the configuration file of timer")

func main() {
	var (
		err error
		ctx context.Context
	)

	flag.Parse()
	ctx = signal.SetupSignalContext()
	cfg, err := timer.InitConfig(*configPath)
	if err != nil {
		log.Error(ctx).Err(err).Msg("init config error")
		os.Exit(-1)
	}

	_ = observability.Initialize(ctx, cfg.Observability, metrics.GetTimerMetrics)

	// new leaderelection manager
	leaderelectionMgr := leaderelection.NewLeaderElection(cfg.GetLeaderElectionConfig())
	// new timingwheel manager
	timingwheelMgr := timingwheel.NewTimingWheel(cfg.GetTimingWheelConfig())

	// init timingwheel
	if err = timingwheelMgr.Init(ctx); err != nil {
		log.Error(ctx).Err(err).Msg("init timer wheel failed")
		os.Exit(-1)
	}

	// define leaderelection callback
	callbacks := leaderelection.LeaderCallbacks{
		OnStartedLeading: func(ctx context.Context) {
			log.Info(ctx).Msg("leaderelection finish, become leader")
			if timingwheelMgr.IsDeployed(ctx) {
				err := timingwheelMgr.Recover(ctx)
				if err != nil {
					log.Error(ctx).Err(err).Msg("recover for fail-over failed, keeping follower")
					return
				}
			}
			timingwheelMgr.SetLeader(true)
		},
		OnStoppedLeading: func(ctx context.Context) {
			log.Info(ctx).Msg("leaderelection lost, become follower")
			timingwheelMgr.SetLeader(false)
		},
	}

	// start leaderelection
	if err = leaderelectionMgr.Start(ctx, callbacks); err != nil {
		log.Error(ctx).Err(err).Msg("start leader election failed")
		os.Exit(-1)
	}

	// start timingwheel
	if err = timingwheelMgr.Start(ctx); err != nil {
		log.Error(ctx).Err(err).Msg("start timer wheel failed")
		os.Exit(-1)
	}

	select {
	case <-ctx.Done():
		log.Info(ctx).Msg("received system signal, preparing exit")
	case <-timingwheelMgr.StopNotify():
		log.Info(ctx).Msg("received timingwheel manager ready to stop, preparing exit")
		signal.RequestShutdown()
	}

	_ = leaderelectionMgr.Stop(context.Background())
	timingwheelMgr.Stop(context.Background())

	log.Info(ctx).Msg("the tiemr has been shutdown gracefully")
}
