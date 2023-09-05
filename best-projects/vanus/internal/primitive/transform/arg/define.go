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

package arg

import (
	"github.com/vanus-labs/vanus/internal/primitive/transform/context"
)

type define struct {
	name     string
	original string
}

// newDefine name format is <var> .
func newDefine(name string) Arg {
	return define{
		name:     name[1 : len(name)-1],
		original: name,
	}
}

func (arg define) Type() Type {
	return Define
}

func (arg define) Name() string {
	return arg.name
}

func (arg define) Original() string {
	return arg.original
}

func (arg define) Evaluate(ceCtx *context.EventContext) (interface{}, error) {
	if len(ceCtx.Define) == 0 {
		return nil, ErrArgValueNil
	}
	v, exist := ceCtx.Define[arg.name]
	if !exist {
		return nil, ErrArgValueNil
	}
	return v, nil
}

func (arg define) SetValue(*context.EventContext, interface{}) error {
	return ErrOperationNotSupport
}

func (arg define) DeleteValue(*context.EventContext) error {
	return ErrOperationNotSupport
}
