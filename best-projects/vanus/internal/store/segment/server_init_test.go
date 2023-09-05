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

package segment

import (
	// standard libraries.
	"context"
	"testing"

	// third-party libraries.
	. "github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"

	// this project.
	"github.com/vanus-labs/vanus/internal/store"
)

func TestServer_recover(t *testing.T) {
	ctrl := NewController(t)
	defer ctrl.Finish()

	Convey("recover", t, func() {
		dir := t.TempDir()

		srv := &server{
			cfg: store.Config{
				Volume: store.VolumeInfo{
					Dir: dir,
				},
			},
		}
		err := srv.loadVSBEngine(context.Background(), srv.cfg.VSB)
		So(err, ShouldBeNil)

		srv.initRaftEngine(context.Background(), srv.cfg.Raft)
		// So(srv.metaStore, ShouldNotBeNil)
		// So(srv.offsetStore, ShouldNotBeNil)
		// So(srv.wal, ShouldNotBeNil)

		err = srv.recover(context.Background())
		So(err, ShouldBeNil)
	})
}
