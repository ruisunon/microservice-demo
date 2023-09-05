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

package vsb

import (
	// standard libraries.
	"encoding/binary"
	"testing"
	"time"

	// third-party libraries.
	. "github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"

	// first-party libraries.
	cepb "github.com/vanus-labs/vanus/proto/pkg/cloudevents"

	// this project.
	"github.com/vanus-labs/vanus/internal/store/block"
	ceschema "github.com/vanus-labs/vanus/internal/store/schema/ce"
	"github.com/vanus-labs/vanus/internal/store/schema/ce/convert"
	cetest "github.com/vanus-labs/vanus/internal/store/schema/ce/testing"
	"github.com/vanus-labs/vanus/internal/store/vsb/codec"
	"github.com/vanus-labs/vanus/internal/store/vsb/entry"
	vsbtest "github.com/vanus-labs/vanus/internal/store/vsb/testing"
)

func TestFragment(t *testing.T) {
	ctrl := NewController(t)
	defer ctrl.Finish()

	ent0 := cetest.MakeStoredEntry0(ctrl)
	ent1 := cetest.MakeStoredEntry1(ctrl, true)

	enc := codec.NewEncoder()

	var offBuf [8]byte

	Convey("fragment", t, func() {
		frag0 := newFragment(vsbtest.EntryOffset0, []block.Entry{ent0}, enc)
		So(frag0.Size(), ShouldEqual, vsbtest.EntrySize0)
		So(frag0.Payload(), ShouldResemble, vsbtest.EntryData0)
		So(frag0.StartOffset(), ShouldEqual, vsbtest.EntryOffset0)
		So(frag0.EndOffset(), ShouldEqual, vsbtest.EntryOffset0+vsbtest.EntrySize0)
		mshl0, ok0 := frag0.(block.FragmentMarshaler)
		So(ok0, ShouldBeTrue)
		data0, err0 := mshl0.MarshalFragment()
		So(err0, ShouldBeNil)
		binary.LittleEndian.PutUint64(offBuf[:], uint64(vsbtest.EntryOffset0))
		So(data0, ShouldResemble, append(offBuf[:], vsbtest.EntryData0...))

		frag1 := newFragment(vsbtest.EntryOffset1, []block.Entry{ent1}, enc)
		So(frag1.Size(), ShouldEqual, vsbtest.EntrySize1)
		So(frag1.Payload(), ShouldResemble, vsbtest.EntryData1)
		So(frag1.StartOffset(), ShouldEqual, vsbtest.EntryOffset1)
		So(frag1.EndOffset(), ShouldEqual, vsbtest.EntryOffset1+vsbtest.EntrySize1)
		mshl1, ok1 := frag1.(block.FragmentMarshaler)
		So(ok1, ShouldBeTrue)
		data1, err1 := mshl1.MarshalFragment()
		So(err1, ShouldBeNil)
		binary.LittleEndian.PutUint64(offBuf[:], uint64(vsbtest.EntryOffset1))
		So(data1, ShouldResemble, append(offBuf[:], vsbtest.EntryData1...))

		frag2 := newFragment(vsbtest.EntryOffset0, []block.Entry{ent0, ent1}, enc)
		So(frag2.Size(), ShouldEqual, vsbtest.EntrySize0+vsbtest.EntrySize1)
		So(frag2.Payload(), ShouldResemble, append(vsbtest.EntryData0, vsbtest.EntryData1...))
		So(frag2.StartOffset(), ShouldEqual, vsbtest.EntryOffset0)
		So(frag2.EndOffset(), ShouldEqual, vsbtest.EntryOffset0+vsbtest.EntrySize0+vsbtest.EntrySize1)
		mshl2, ok2 := frag2.(block.FragmentMarshaler)
		So(ok2, ShouldBeTrue)
		data2, err2 := mshl2.MarshalFragment()
		So(err2, ShouldBeNil)
		binary.LittleEndian.PutUint64(offBuf[:], uint64(vsbtest.EntryOffset0))
		So(data2, ShouldResemble, append(append(offBuf[:], vsbtest.EntryData0...), vsbtest.EntryData1...))
	})
}

func BenchmarkSize(b *testing.B) {
	c := codec.NewEncoder()
	e := getEntry()
	b.Run("Size", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c.Size(e)
		}
	})
}

func BenchmarkFragment_MarshalFragment(b *testing.B) {
	c := codec.NewEncoder()
	e := getEntry()
	buf := make([]byte, c.Size(e))
	b.Run("MarshalFragment", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = c.MarshalTo(e, buf)
		}
	})
}

func getEntry() block.Entry {
	ce := &cepb.CloudEvent{
		Id:          "benchmark1",
		Source:      "benchmark2",
		SpecVersion: "1.0",
		Type:        "benchmark3",
		Data: &cepb.CloudEvent_TextData{
			TextData: "benchmark4",
		},
	}
	e := convert.ToEntry(ce)

	return entry.Wrap(e, ceschema.CloudEvent, 111, time.Now().UnixMilli())
}
