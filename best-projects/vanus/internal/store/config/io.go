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

package config

import (
	// this project.
	"github.com/vanus-labs/vanus/internal/store/io/engine"
	"github.com/vanus-labs/vanus/internal/store/io/engine/psync"
)

type IOEngineType string

const (
	Psync IOEngineType = "psync"
	Uring IOEngineType = "io_uring"
)

type IO struct {
	Engine   IOEngineType `yaml:"engine"`
	Parallel int          `yaml:"parallel"`
}

func buildIOEngine(cfg IO) engine.Interface {
	switch cfg.Engine {
	case Psync:
		return buildPsyncEngine(cfg)
	default:
		return buildIOEngineEx(cfg)
	}
}

func buildPsyncEngine(cfg IO) engine.Interface {
	var opts []psync.Option
	if cfg.Parallel > 0 {
		opts = append(opts, psync.WithParallel(cfg.Parallel))
	}
	return psync.New(opts...)
}
