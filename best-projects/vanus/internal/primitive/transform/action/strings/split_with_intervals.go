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

package strings

import (
	"fmt"

	"github.com/vanus-labs/vanus/internal/primitive/transform/action"
	"github.com/vanus-labs/vanus/internal/primitive/transform/arg"
	"github.com/vanus-labs/vanus/internal/primitive/transform/common"
	"github.com/vanus-labs/vanus/internal/primitive/transform/context"
)

type splitWithIntervalsAction struct {
	action.CommonAction
}

// NewSplitWithIntervalsAction["sourceJSONPath", "startPosition", "splitInterval", "targetJsonPath"].
func NewSplitWithIntervalsAction() action.Action {
	return &splitWithIntervalsAction{
		CommonAction: action.CommonAction{
			ActionName: "SPLIT_WITH_INTERVALS",
			FixedArgs:  []arg.TypeList{arg.EventList, arg.All, arg.All, []arg.Type{arg.EventData}},
		},
	}
}

func (a *splitWithIntervalsAction) Init(args []arg.Arg) error {
	a.TargetArg = args[3]
	a.Args = args[:3]
	a.ArgTypes = []common.Type{common.String, common.Int, common.Int}
	return nil
}

func (a *splitWithIntervalsAction) Execute(ceCtx *context.EventContext) error {
	args, err := a.RunArgs(ceCtx)
	if err != nil {
		return err
	}

	v, _ := a.TargetArg.Evaluate(ceCtx)
	if v != nil {
		return fmt.Errorf("key %s exists", a.TargetArg.Original())
	}

	sourceJSONPath, _ := args[0].(string)
	startPosition, _ := args[1].(int)
	splitInterval, _ := args[2].(int)

	// split string
	var substrings []string
	if startPosition > len(sourceJSONPath) {
		// if startPosition is beyond the end of the string, return an error
		return a.TargetArg.SetValue(ceCtx, []string{sourceJSONPath})
	}

	// split the string according to the specified interval
	substrings = []string{sourceJSONPath[:startPosition]}
	for i := startPosition; i < len(sourceJSONPath); i += splitInterval {
		end := i + splitInterval
		if end > len(sourceJSONPath) {
			end = len(sourceJSONPath)
		}
		substrings = append(substrings, sourceJSONPath[i:end])
	}

	return a.TargetArg.SetValue(ceCtx, substrings)
}
