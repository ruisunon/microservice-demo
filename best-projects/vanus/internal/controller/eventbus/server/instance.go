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

//go:generate mockgen -source=instance.go -destination=mock_instance.go -package=server
package server

import (
	"context"
	"sync"

	"github.com/vanus-labs/vanus/internal/controller/eventbus/metadata"
	"github.com/vanus-labs/vanus/internal/primitive/vanus"
	"github.com/vanus-labs/vanus/observability/log"
	"github.com/vanus-labs/vanus/pkg/errors"
	segpb "github.com/vanus-labs/vanus/proto/pkg/segment"
)

type Instance interface {
	ID() vanus.ID
	Address() string
	Close() error
	GetMeta() *metadata.VolumeMetadata
	CreateBlock(context.Context, int64) (*metadata.Block, error)
	DeleteBlock(context.Context, vanus.ID) error
	GetServer() Server
	SetServer(Server)
}

func NewInstance(md *metadata.VolumeMetadata) Instance {
	return &volumeInstance{
		md: md,
	}
}

type volumeInstance struct {
	md        *metadata.VolumeMetadata
	metaMutex sync.Mutex
	srv       Server
	rwMutex   sync.RWMutex
}

func (ins *volumeInstance) GetMeta() *metadata.VolumeMetadata {
	return ins.md
}

func (ins *volumeInstance) CreateBlock(ctx context.Context, capacity int64) (*metadata.Block, error) {
	id, err := vanus.NewID()
	if err != nil {
		return nil, err
	}
	blk := &metadata.Block{
		ID:       id,
		Capacity: capacity,
		VolumeID: ins.md.ID,
	}
	if ins.srv == nil {
		return nil, errors.ErrVolumeInstanceNoServer
	}
	_, err = ins.srv.GetClient().CreateBlock(ctx, &segpb.CreateBlockRequest{
		Size: blk.Capacity,
		Id:   blk.ID.Uint64(),
	})
	if err != nil {
		log.Warn(ctx).
			Interface("id", id).
			Interface("volume_id", ins.md.ID).
			Msg("failed to create block")
		return nil, err
	}

	log.Info(ctx).Interface("id", id).Msg("success to create block")

	ins.metaMutex.Lock()
	defer ins.metaMutex.Unlock()
	ins.md.Used += capacity
	ins.md.Blocks[blk.ID.Uint64()] = blk
	return blk, nil
}

func (ins *volumeInstance) DeleteBlock(ctx context.Context, id vanus.ID) error {
	if ins.srv == nil {
		return errors.ErrVolumeInstanceNoServer
	}
	blk := ins.md.Blocks[id.Uint64()]
	if blk == nil {
		return nil
	}
	if ins.srv == nil {
		return nil
	}
	_, err := ins.srv.GetClient().RemoveBlock(ctx, &segpb.RemoveBlockRequest{Id: id.Uint64()})
	if err != nil {
		return err
	}
	ins.metaMutex.Lock()
	defer ins.metaMutex.Unlock()

	ins.md.Used -= blk.Capacity
	delete(ins.md.Blocks, blk.ID.Uint64())
	return nil
}

func (ins *volumeInstance) ID() vanus.ID {
	return ins.md.ID
}

func (ins *volumeInstance) Address() string {
	ins.rwMutex.RLock()
	defer ins.rwMutex.RUnlock()

	if ins.srv != nil {
		return ins.srv.Address()
	}
	return ""
}

func (ins *volumeInstance) Close() error {
	ins.rwMutex.RLock()
	defer ins.rwMutex.RUnlock()

	if ins.srv != nil {
		return ins.srv.Close()
	}
	return nil
}

func (ins *volumeInstance) SetServer(srv Server) {
	ins.rwMutex.Lock()
	defer ins.rwMutex.Unlock()
	if srv == nil || !srv.IsActive(context.Background()) {
		return
	}
	log.Info().Interface("srv", srv.VolumeID()).
		Str("address", srv.Address()).
		Time("uptime", srv.Uptime()).
		Msg("update server of volume")
	ins.srv = srv
}

func (ins *volumeInstance) GetServer() Server {
	ins.rwMutex.Lock()
	defer ins.rwMutex.Unlock()
	if ins.srv == nil {
		return nil
	}
	if !ins.srv.IsActive(context.Background()) {
		ins.srv = nil
		return nil
	}
	return ins.srv
}
