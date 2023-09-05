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

package pipeline

import (
	"github.com/vanus-labs/vanus/internal/primitive"
	"github.com/vanus-labs/vanus/internal/primitive/transform/action"
	"github.com/vanus-labs/vanus/internal/primitive/transform/context"
	"github.com/vanus-labs/vanus/internal/primitive/transform/runtime"
	"github.com/vanus-labs/vanus/observability/log"
)

type Pipeline struct {
	actions []action.Action
}

func NewPipeline() *Pipeline {
	return &Pipeline{
		actions: make([]action.Action, 0),
	}
}

func (p *Pipeline) Parse(actions []*primitive.Action) {
	p.actions = make([]action.Action, 0, len(actions))
	for i := range actions {
		_action, err := runtime.NewAction(actions[i].Command)
		if err != nil {
			// it has check in controller so err must be nil otherwise controller check has a bug
			log.Warn().Err(err).
				Interface("command", actions[i].Command).
				Msg("new action error")
			continue
		}
		p.actions = append(p.actions, _action)
	}
}

func (p *Pipeline) Run(ceCtx *context.EventContext) error {
	for _, a := range p.actions {
		err := a.Execute(ceCtx)
		if err != nil {
			log.Warn().Err(err).
				Str("command", a.Name()).
				Msg("action execute error")
		}
	}
	return nil
}
