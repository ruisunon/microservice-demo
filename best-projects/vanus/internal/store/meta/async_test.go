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

package meta

import (
	// standard libraries.
	"context"
	"testing"

	// third-party libraries.
	. "github.com/smartystreets/goconvey/convey"
)

func TestAsyncStore(t *testing.T) {
	ctx := context.Background()

	Convey("AsyncStore", t, func() {
		dir := t.TempDir()

		Convey("new empty AsyncStore by recovery", func() {
			ss, err := RecoverAsyncStore(ctx, dir)

			So(err, ShouldBeNil)
			So(ss, ShouldNotBeNil)

			ss.Close()
		})

		Convey("setup AsyncStore", func() {
			ss, err := RecoverAsyncStore(ctx, dir)
			So(err, ShouldBeNil)
			ss.Store(ctx, key0, "value0")
			ss.Store(ctx, key1, "value1")
			ss.Close()

			Convey("recover AsyncStore", func() {
				ss, err = RecoverAsyncStore(ctx, dir)
				So(err, ShouldBeNil)

				value0, ok0 := ss.Load(key0)
				So(ok0, ShouldBeTrue)
				So(value0, ShouldResemble, "value0")

				value1, ok1 := ss.Load(key1)
				So(ok1, ShouldBeTrue)
				So(value1, ShouldResemble, "value1")

				_, ok2 := ss.Load(key2)
				So(ok2, ShouldBeFalse)

				Convey("modify AsyncStore", func() {
					ss.Delete(key1)
					_, ok1 = ss.Load(key1)
					So(ok1, ShouldBeFalse)

					ss.Store(ctx, key2, "value2")
					value2, ok2 := ss.Load(key2)
					So(ok2, ShouldBeTrue)
					So(value2, ShouldResemble, "value2")

					ss.Close()

					Convey("recover AsyncStore again", func() {
						ss, err = RecoverAsyncStore(ctx, dir)
						So(err, ShouldBeNil)

						value0, ok0 := ss.Load(key0)
						So(ok0, ShouldBeTrue)
						So(value0, ShouldResemble, "value0")

						_, ok1 := ss.Load(key1)
						So(ok1, ShouldBeFalse)

						value2, ok2 := ss.Load(key2)
						So(ok2, ShouldBeTrue)
						So(value2, ShouldResemble, "value2")

						ss.Close()
					})
				})
			})
		})
	})
}
