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
	"context"
	"fmt"
	"os"
	"path/filepath"

	// first-party libraries.
	"github.com/vanus-labs/vanus/pkg/errors"

	// this project.
	"github.com/vanus-labs/vanus/internal/primitive/vanus"
	"github.com/vanus-labs/vanus/internal/store/block"
	"github.com/vanus-labs/vanus/internal/store/io"
	"github.com/vanus-labs/vanus/internal/store/io/zone/file"
	"github.com/vanus-labs/vanus/internal/store/vsb/codec"
)

const (
	vsbExt          = ".vsb"
	defaultFilePerm = 0o644
)

func (e *engine) Create(ctx context.Context, id vanus.ID, capacity int64) (block.Raw, error) {
	path := e.resolvePath(id)

	f, err := io.CreateFile(path, capacity, os.O_RDWR, true, false)
	if err != nil {
		return nil, err
	}

	dec, _ := codec.NewDecoder(false, codec.IndexSize)
	b := &vsBlock{
		id:         id,
		path:       path,
		capacity:   capacity,
		dataOffset: headerBlockSize,
		indexSize:  codec.IndexSize,
		fm: meta{
			writeOffset: headerBlockSize,
		},
		actx: appendContext{
			offset: headerBlockSize,
		},
		enc: codec.NewEncoder(),
		dec: dec,
		lis: e.lis,
		f:   f,
	}

	if err := b.persistHeader(ctx, b.fm); err != nil {
		return nil, processError(err, f, path)
	}

	if z, err := file.New(f); err == nil {
		b.z = z
	} else {
		return nil, processError(err, f, path)
	}

	b.s = e.s.Register(b.z, b.actx.offset, false)

	return b, nil
}

func processError(err error, f *os.File, path string) error {
	if err2 := f.Close(); err2 != nil {
		return errors.Chain(err, err2)
	}
	if err2 := os.Remove(path); err2 != nil {
		return errors.Chain(err, err2)
	}
	return err
}

func (e *engine) Open(ctx context.Context, id vanus.ID) (block.Raw, error) {
	path := e.resolvePath(id)

	b := &vsBlock{
		id:   id,
		path: path,
		lis:  e.lis,
	}

	if err := b.Open(ctx); err != nil {
		return nil, err
	}

	if z, err := file.New(b.f); err == nil {
		b.z = z
	} else {
		return nil, err
	}

	b.s = e.s.Register(b.z, b.actx.offset, false)

	return b, nil
}

func (e *engine) resolvePath(id vanus.ID) string {
	return filepath.Join(e.dir, fmt.Sprintf("%s%s", id.String(), vsbExt))
}
