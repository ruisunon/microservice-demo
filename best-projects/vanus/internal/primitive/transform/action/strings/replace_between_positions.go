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
	"github.com/vanus-labs/vanus/internal/primitive/transform/action"
	"github.com/vanus-labs/vanus/internal/primitive/transform/arg"
	"github.com/vanus-labs/vanus/internal/primitive/transform/function"
)

// NewReplaceBetweenPositionsAction ["path","startPosition","endPosition","targetValue"].
func NewReplaceBetweenPositionsAction() action.Action {
	a := &action.SourceTargetSameAction{}
	a.CommonAction = action.CommonAction{
		ActionName: "REPLACE_BETWEEN_POSITIONS",
		FixedArgs:  []arg.TypeList{arg.EventList, arg.All, arg.All, arg.All},
		Fn:         function.ReplaceBetweenPositionsFunction,
	}
	return a
}
