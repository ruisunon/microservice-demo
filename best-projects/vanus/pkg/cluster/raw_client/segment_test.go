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

package raw_client

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/vanus-labs/vanus/pkg/errors"
)

func TestCtrlClientIsNeedRetry(t *testing.T) {
	Convey("test isNeedRetry", t, func() {
		// cli := NewSegmentClient([]string{"127.0.0.1:2048", "127.0.0.1:2148", "127.0.0.1:2248"}, insecure.NewCredentials())
		Convey("test error internal", func() {
			err := error(errors.ErrNoControllerLeader)
			So(isNeedRetry(err), ShouldBeTrue)

			err = fmt.Errorf("test error")
			So(isNeedRetry(err), ShouldBeFalse)
		})

		Convey("test error returned from gRPC", func() {
			err := fmt.Errorf("xxxxx, please connect to: 127.a.0.1 ")
			So(isNeedRetry(err), ShouldBeFalse)

			err = errors.New("xxxxx: 1111111111 ")
			So(isNeedRetry(err), ShouldBeFalse)

			err = errors.New("balabala, please connect to: 127.0.0.1:2048 ").WithGRPCCode(errors.ErrorCode_NOT_LEADER)
			So(isNeedRetry(err), ShouldBeTrue)
		})
	})
}
