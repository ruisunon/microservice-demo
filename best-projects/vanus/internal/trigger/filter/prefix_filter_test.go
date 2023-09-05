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

package filter_test

import (
	"testing"

	ce "github.com/cloudevents/sdk-go/v2"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/vanus-labs/vanus/internal/trigger/filter"
)

func TestPrefixFilter(t *testing.T) {
	event := ce.NewEvent()
	event.SetID("testID")
	event.SetSource("testSource")
	Convey("prefix filter nil", t, func() {
		f := filter.NewPrefixFilter(map[string]string{
			"": "testID",
		})
		So(f, ShouldBeNil)
		f = filter.NewPrefixFilter(map[string]string{
			"k": "",
		})
		So(f, ShouldBeNil)
	})
	Convey("prefix filter pass", t, func() {
		f := filter.NewPrefixFilter(map[string]string{
			"id":     "test",
			"source": "test",
		})
		result := f.Filter(event)
		So(result, ShouldEqual, filter.PassFilter)
	})
	Convey("prefix filter fail", t, func() {
		f := filter.NewPrefixFilter(map[string]string{
			"id":     "un",
			"source": "test",
		})
		result := f.Filter(event)
		So(result, ShouldEqual, filter.FailFilter)
	})
}
