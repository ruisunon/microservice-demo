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
	"context"
	"time"

	"github.com/vanus-labs/vanus/observability/log"
	"github.com/vanus-labs/vanus/pkg/errors"
)

type Heartbeat interface {
	Beat(ctx context.Context, req interface{}) error
}

func RegisterHeartbeat(ctx context.Context, interval time.Duration,
	i interface{}, reqFunc func() interface{},
) error {
	hb, ok := i.(Heartbeat)
	if !ok {
		return errors.ErrInvalidHeartBeat
	}
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				// TODO log
				return
			case <-ticker.C:
				if err := hb.Beat(ctx, reqFunc()); err != nil {
					log.Warn().Err(err).Msg("heartbeat error")
				}
			}
		}
	}()
	return nil
}
