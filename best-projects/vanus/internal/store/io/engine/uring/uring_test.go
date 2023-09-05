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

//go:build linux
// +build linux

package uring

import (
	// standard libraries.
	"os"
	"testing"

	// third-party libraries.
	. "github.com/smartystreets/goconvey/convey"

	// this project.
	enginetest "github.com/vanus-labs/vanus/internal/store/io/engine/testing"
)

func TestURing(t *testing.T) {
	f, err := os.CreateTemp("", "wal-engine-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(f.Name())

	e := New()
	defer e.Close()

	Convey("uRing", t, func() {
		enginetest.DoEngineTest(e, f)
	})
}
