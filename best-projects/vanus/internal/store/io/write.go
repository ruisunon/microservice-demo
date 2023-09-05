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

package io

type WriteCallback = func(n int, err error)

type WriterAt interface {
	WriteAt(b []byte, off int64, so, eo int, cb WriteCallback)
}

type WriteAtFunc func(b []byte, off int64, so, eo int, cb WriteCallback)

// Make sure WriteAtFunc implements WriterAt.
var _ WriterAt = (WriteAtFunc)(nil)

func (f WriteAtFunc) WriteAt(b []byte, off int64, so, eo int, cb WriteCallback) {
	f(b, off, so, eo, cb)
}
