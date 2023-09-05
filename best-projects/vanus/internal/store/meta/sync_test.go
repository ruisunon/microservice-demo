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

var (
	key0 = []byte("key0")
	key1 = []byte("key1")
	key2 = []byte("key2")
)

func TestSyncStore(t *testing.T) {
	Convey("SyncStore", t, func() {
		dir := t.TempDir()

		Convey("new empty SyncStore by recovery", func() {
			ss, err := RecoverSyncStore(context.Background(), dir)

			So(err, ShouldBeNil)
			So(ss, ShouldNotBeNil)

			ss.Close(context.Background())
		})

		Convey("setup SyncStore", func() {
			ss, err := RecoverSyncStore(context.Background(), dir)
			So(err, ShouldBeNil)
			Store(context.Background(), ss, key0, "value0")
			Store(context.Background(), ss, key1, "value1")
			ss.Close(context.Background())

			Convey("recover SyncStore", func() {
				ss, err = RecoverSyncStore(context.Background(), dir)
				So(err, ShouldBeNil)

				value0, ok0 := ss.Load(key0)
				So(ok0, ShouldBeTrue)
				So(value0, ShouldResemble, "value0")

				value1, ok1 := ss.Load(key1)
				So(ok1, ShouldBeTrue)
				So(value1, ShouldResemble, "value1")

				_, ok2 := ss.Load(key2)
				So(ok2, ShouldBeFalse)

				Convey("modify SyncStore", func() {
					Delete(context.Background(), ss, key1)
					_, ok1 = ss.Load(key1)
					So(ok1, ShouldBeFalse)

					Store(context.Background(), ss, key2, "value2")
					value2, ok2 := ss.Load(key2)
					So(ok2, ShouldBeTrue)
					So(value2, ShouldResemble, "value2")

					ss.Close(context.Background())

					Convey("recover SyncStore again", func() {
						ss, err = RecoverSyncStore(context.Background(), dir)
						So(err, ShouldBeNil)

						value0, ok0 := ss.Load(key0)
						So(ok0, ShouldBeTrue)
						So(value0, ShouldResemble, "value0")

						_, ok1 := ss.Load(key1)
						So(ok1, ShouldBeFalse)

						value2, ok2 := ss.Load(key2)
						So(ok2, ShouldBeTrue)
						So(value2, ShouldResemble, "value2")

						ss.Close(context.Background())
					})
				})
			})
		})
	})
}
