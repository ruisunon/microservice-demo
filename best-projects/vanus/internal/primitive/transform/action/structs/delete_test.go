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

package structs_test

import (
	"testing"

	cetest "github.com/cloudevents/sdk-go/v2/test"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/vanus-labs/vanus/internal/primitive/transform/action/structs"
	"github.com/vanus-labs/vanus/internal/primitive/transform/context"
	"github.com/vanus-labs/vanus/internal/primitive/transform/runtime"
)

func TestDeleteAction(t *testing.T) {
	funcName := structs.NewDeleteAction().Name()
	Convey("test delete", t, func() {
		Convey("delete cant not delete key", func() {
			a, err := runtime.NewAction([]interface{}{funcName, "$.id"})
			So(err, ShouldBeNil)
			e := cetest.MinEvent()
			err = a.Execute(&context.EventContext{
				Event: &e,
			})
			So(err, ShouldNotBeNil)
		})
		Convey("delete", func() {
			a, err := runtime.NewAction([]interface{}{funcName, "$.test"})
			So(err, ShouldBeNil)
			e := cetest.MinEvent()
			e.SetExtension("test", "abc")
			err = a.Execute(&context.EventContext{
				Event: &e,
			})
			So(err, ShouldBeNil)
			So(len(e.Extensions()), ShouldEqual, 0)
		})
	})
}
