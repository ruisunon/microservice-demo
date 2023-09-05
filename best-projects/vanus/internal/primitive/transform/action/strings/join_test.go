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

package strings_test

import (
	"testing"

	cetest "github.com/cloudevents/sdk-go/v2/test"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/vanus-labs/vanus/internal/primitive/transform/action/strings"
	"github.com/vanus-labs/vanus/internal/primitive/transform/context"
	"github.com/vanus-labs/vanus/internal/primitive/transform/runtime"
)

func TestJoinAction(t *testing.T) {
	funcName := strings.NewJoinAction().Name()
	Convey("test join action", t, func() {
		e := cetest.MinEvent()
		eventCtx := &context.EventContext{
			Event: &e,
			Data: map[string]interface{}{
				"array": []map[string]interface{}{
					{"key1": "value1"},
					{"key1": "value11"},
					{"key1": "value111"},
				},
			},
		}
		Convey("test string", func() {
			Convey("test one param", func() {
				a, err := runtime.NewAction([]interface{}{funcName, "$.test1", ",", "abc"})
				So(err, ShouldBeNil)
				err = a.Execute(eventCtx)
				So(err, ShouldBeNil)
				So(eventCtx.Event.Extensions()["test1"], ShouldEqual, "abc")
			})
			Convey("test many param", func() {
				a, err := runtime.NewAction([]interface{}{funcName, "$.test2", ",", "abc", "123"})
				So(err, ShouldBeNil)
				err = a.Execute(eventCtx)
				So(err, ShouldBeNil)
				So(eventCtx.Event.Extensions()["test2"], ShouldEqual, "abc,123")
			})
		})
		Convey("test string array", func() {
			Convey("test one param", func() {
				a, err := runtime.NewAction([]interface{}{funcName, "$.array1", ",", "$.data.array[:].key1"})
				So(err, ShouldBeNil)
				err = a.Execute(eventCtx)
				So(err, ShouldBeNil)
				So(eventCtx.Event.Extensions()["array1"], ShouldEqual, "value1,value11,value111")
			})
			Convey("test many mixture param", func() {
				a, err := runtime.NewAction([]interface{}{funcName, "$.array2", ",", "$.data.array[:].key1", "abc"})
				So(err, ShouldBeNil)
				err = a.Execute(eventCtx)
				So(err, ShouldBeNil)
				So(eventCtx.Event.Extensions()["array2"], ShouldEqual, "value1,value11,value111,abc")
			})
		})
	})
}
