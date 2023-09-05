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

package wal

import (
	// standard libraries.
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync/atomic"
	"testing"
	"time"

	// third-party libraries.
	. "github.com/smartystreets/goconvey/convey"

	// this project.
	"github.com/vanus-labs/vanus/internal/store/wal/record"
)

var (
	fileSize int64 = 8 * defaultBlockSize
	data0          = []byte{0x41, 0x42, 0x43}
	data1          = []byte{0x44, 0x45, 0x46, 0x47}
)

func TestWAL_AppendOne(t *testing.T) {
	ctx := context.Background()

	Convey("wal append one", t, func() {
		walDir := t.TempDir()

		wal, err := Open(ctx, walDir, WithFileSize(fileSize))
		So(err, ShouldBeNil)

		Convey("append one with callback", func() {
			var done bool

			wal.AppendOne(ctx, data0, func(Range, error) {
				done = true
			})
			n, _ := AppendOne(ctx, wal, data1)

			// Invoke callback of append data0, before append data1 return.
			So(done, ShouldBeTrue)
			So(n.EO, ShouldEqual, 21)
			So(wal.s.WriteOffset(), ShouldEqual, 21)
			// So(wal.wb.Committed(), ShouldEqual, 21)

			filePath := filepath.Join(walDir, fmt.Sprintf("%020d.log", 0))
			data, err2 := os.ReadFile(filePath)
			So(err2, ShouldBeNil)

			So(data[:21+record.HeaderSize], ShouldResemble,
				[]byte{
					0x7D, 0x7F, 0xEB, 0x7A, 0x00, 0x03, 0x01, 0x41, 0x42, 0x43,
					0x52, 0x74, 0x2F, 0x51, 0x00, 0x04, 0x01, 0x44, 0x45, 0x46, 0x47,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				})
		})

		Convey("append one with large data", func() {
			data := make([]byte, fileSize)

			n, err := DirectAppendOne(ctx, wal, data)

			So(err, ShouldBeNil)
			So(n.EO, ShouldEqual, fileSize+9*record.HeaderSize)
			So(wal.sf.Len(), ShouldEqual, 2)
		})

		Reset(func() {
			wal.Close()
			wal.Wait()
		})
	})

	Convey("flush wal when timeout", t, func() {
		walDir := t.TempDir()

		flushTimeout := time.Second
		wal, err := Open(ctx, walDir, WithFileSize(fileSize), WithFlushDelayTime(flushTimeout))
		So(err, ShouldBeNil)

		data := make([]byte, defaultBlockSize)

		startTime := time.Now()
		var t0, t1 time.Time
		wal.AppendOne(ctx, data0, func(Range, error) {
			t0 = time.Now()
		})
		wal.AppendOne(ctx, data, func(Range, error) {
			t1 = time.Now()
		})
		AppendOne(ctx, wal, data1)
		t2 := time.Now()

		So(t0, ShouldHappenBefore, startTime.Add(flushTimeout))
		So(t1, ShouldHappenAfter, startTime.Add(flushTimeout))
		So(t1, ShouldHappenAfter, t0)
		So(t2, ShouldHappenAfter, t1)

		wal.Close()
		wal.Wait()
	})

	Convey("wal append one after close", t, func() {
		walDir := t.TempDir()

		flushTimeout := 100 * time.Millisecond
		wal, err := Open(ctx, walDir, WithFileSize(fileSize), WithFlushDelayTime(flushTimeout))
		So(err, ShouldBeNil)

		var inflight int32 = 100
		for i := inflight; i > 0; i-- {
			wal.AppendOne(ctx, data0, func(Range, error) {
				atomic.AddInt32(&inflight, -1)
			})
		}

		wal.Close()

		_, err = AppendOne(ctx, wal, data1)
		So(err, ShouldNotBeNil)

		wal.Wait()

		// NOTE: All appends are guaranteed to return before wal is closed.
		So(atomic.LoadInt32(&inflight), ShouldBeZeroValue)

		// NOTE: There is no guarantee that data0 will be successfully written.
		// So(wal.wb.Size(), ShouldEqual, 10)
		// So(wal.wb.Committed(), ShouldEqual, 10)
	})
}

func TestWAL_Append(t *testing.T) {
	ctx := context.Background()

	Convey("wal append", t, func() {
		walDir := t.TempDir()

		wal, err := Open(ctx, walDir, WithFileSize(fileSize))
		So(err, ShouldBeNil)

		Convey("direct append", func() {
			ranges, err := DirectAppend(ctx, wal, [][]byte{data0, data1})

			So(err, ShouldBeNil)
			So(len(ranges), ShouldEqual, 2)
			So(ranges[0].EO, ShouldEqual, 10)
			So(ranges[1].EO, ShouldEqual, 21)
			So(wal.s.WriteOffset(), ShouldEqual, 21)
			// So(wal.wb.Committed(), ShouldEqual, 21)

			filePath := filepath.Join(walDir, fmt.Sprintf("%020d.log", 0))
			data, err2 := os.ReadFile(filePath)
			So(err2, ShouldBeNil)

			So(data[:21+record.HeaderSize], ShouldResemble,
				[]byte{
					0x7D, 0x7F, 0xEB, 0x7A, 0x00, 0x03, 0x01, 0x41, 0x42, 0x43,
					0x52, 0x74, 0x2F, 0x51, 0x00, 0x04, 0x01, 0x44, 0x45, 0x46, 0x47,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				})
		})

		Reset(func() {
			wal.Close()
			wal.Wait()
		})
	})
}

func TestWAL_Compact(t *testing.T) {
	ctx := context.Background()
	data := make([]byte, defaultBlockSize-record.HeaderSize)
	copy(data, []byte("hello world!"))

	Convey("wal compaction", t, func() {
		walDir := t.TempDir()

		wal, err := Open(ctx, walDir, WithFileSize(fileSize))
		So(err, ShouldBeNil)

		_, err = Append(ctx, wal, [][]byte{data, data, data, data, data, data, data, data})
		So(err, ShouldBeNil)
		So(wal.sf.Len(), ShouldEqual, 1)

		r, err := AppendOne(ctx, wal, data)
		So(err, ShouldBeNil)
		So(r, ShouldResemble, Range{SO: fileSize, EO: fileSize + defaultBlockSize})
		So(wal.sf.Len(), ShouldEqual, 2)

		err = wal.Compact(ctx, r.SO)
		So(err, ShouldBeNil)
		So(wal.sf.Len(), ShouldEqual, 1)

		err = wal.Compact(ctx, r.EO)
		So(err, ShouldBeNil)
		So(wal.sf.Len(), ShouldEqual, 1)

		ranges, err := Append(ctx, wal, [][]byte{
			data, data, data, data, data, data, data, data,
			data, data, data, data, data, data, data, data,
		})
		So(err, ShouldBeNil)
		So(ranges[len(ranges)-1].EO, ShouldEqual, fileSize*3+defaultBlockSize)
		So(wal.sf.Len(), ShouldEqual, 3)

		err = wal.Compact(ctx, fileSize*2)
		So(err, ShouldBeNil)
		So(wal.sf.Len(), ShouldEqual, 2)

		wal.Close()
		wal.Wait()
	})
}
