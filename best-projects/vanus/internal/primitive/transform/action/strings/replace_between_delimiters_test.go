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

package strings_test

import (
	"testing"

	cetest "github.com/cloudevents/sdk-go/v2/test"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/vanus-labs/vanus/internal/primitive/transform/action/strings"
	"github.com/vanus-labs/vanus/internal/primitive/transform/context"
	"github.com/vanus-labs/vanus/internal/primitive/transform/runtime"
)

func TestReplaceBetweenDelimitersAction(t *testing.T) {
	funcName := strings.NewReplaceBetweenDelimitersAction().Name()

	Convey("test startDelimiter and endDelimiter present in string testcase", t, func() {
		a, err := runtime.NewAction([]interface{}{funcName, "$.test", "&&", "&&", "Vanus"})
		So(err, ShouldBeNil)
		e := cetest.MinEvent()
		e.SetExtension("test", "Hello, &&World&&!")
		ceCtx := &context.EventContext{
			Event: &e,
		}
		err = a.Execute(ceCtx)
		So(err, ShouldBeNil)
		So(e.Extensions()["test"], ShouldEqual, "Hello, Vanus!")
	})

	Convey("test startDelimiter and endDelimiter present in string testcase", t, func() {
		a, err := runtime.NewAction([]interface{}{funcName, "$.test", "^^", "^^", "lots of"})
		So(err, ShouldBeNil)
		e := cetest.MinEvent()
		e.SetExtension("test", "Vanus has ^^many^^ beginner friendly open issues!")
		ceCtx := &context.EventContext{
			Event: &e,
		}
		err = a.Execute(ceCtx)
		So(err, ShouldBeNil)
		So(e.Extensions()["test"], ShouldEqual, "Vanus has lots of beginner friendly open issues!")
	})

	Convey("test startDelimiter and endDelimiter not present in string testcase", t, func() {
		a, err := runtime.NewAction([]interface{}{funcName, "$.test", "**", "**", "fun"})
		So(err, ShouldBeNil)
		e := cetest.MinEvent()
		e.SetExtension("test", "Contributing to Vanus Opensource project is %%an eye opener%%!")
		ceCtx := &context.EventContext{
			Event: &e,
		}
		err = a.Execute(ceCtx)
		So(err, ShouldNotBeNil)
		So(e.Extensions()["test"], ShouldEqual, "Contributing to Vanus Opensource project is %%an eye opener%%!")
	})

	Convey("test endDelimiter before startDelimiter in string testcase", t, func() {
		a, err := runtime.NewAction([]interface{}{funcName, "$.test", "&&", "!!", "love"})
		So(err, ShouldBeNil)
		e := cetest.MinEvent()
		e.SetExtension("test", "I !!like&& opensource contributions")
		ceCtx := &context.EventContext{
			Event: &e,
		}
		err = a.Execute(ceCtx)
		So(err, ShouldNotBeNil)
		So(e.Extensions()["test"], ShouldEqual, "I !!like&& opensource contributions")
	})

	Convey("test Only endDelimiter present in string testcase", t, func() {
		a, err := runtime.NewAction([]interface{}{funcName, "$.test", "&&", "**", "supported"})
		So(err, ShouldBeNil)
		e := cetest.MinEvent()
		e.SetExtension("test", "FOSS is !!powered** by open communities")
		ceCtx := &context.EventContext{
			Event: &e,
		}
		err = a.Execute(ceCtx)
		So(err, ShouldNotBeNil)
		So(e.Extensions()["test"], ShouldEqual, "FOSS is !!powered** by open communities")
	})
}
