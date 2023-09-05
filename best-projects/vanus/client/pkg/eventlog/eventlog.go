// Copyright 2022 Linkall Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:generate mockgen -source=eventlog.go -destination=mock_eventlog.go -package=eventlog
package eventlog

import (
	// standard libraries.
	"context"

	// first-party libraries.
	"github.com/vanus-labs/vanus/proto/pkg/cloudevents"
	segpb "github.com/vanus-labs/vanus/proto/pkg/segment"

	// this project.
	"github.com/vanus-labs/vanus/client/pkg/api"
)

const (
	XVanusLogOffset = segpb.XVanusLogOffset
)

type ReaderConfig struct {
	PollingTimeout int64
}

type Eventlog interface {
	api.Eventlog

	Close(ctx context.Context)
	Writer() LogWriter
	Reader(cfg ReaderConfig) LogReader
}

type LogWriter interface {
	Log() Eventlog

	Close(ctx context.Context)

	Append(ctx context.Context, events *cloudevents.CloudEventBatch) (offs []int64, err error)
}

type LogReader interface {
	Log() Eventlog

	Close(ctx context.Context)

	// TODO: async
	Read(ctx context.Context, size int16) (events *cloudevents.CloudEventBatch, err error)

	// Seek sets the offset for the next Read to offset,
	// interpreted according to whence.
	//
	// `Seek(context.Background(), 0, io.SeekCurrent)` will return current offset.
	//
	// Also see `io.Seeker`.
	Seek(ctx context.Context, offset int64, whence int) (off int64, err error)
}
