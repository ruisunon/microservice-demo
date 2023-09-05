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
	// standard libraries.
	"context"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	// third-party libraries.
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	// first-party libraries.
	"github.com/vanus-labs/vanus/observability/log"
	ctrlpb "github.com/vanus-labs/vanus/proto/pkg/controller"

	// this project.
	"github.com/vanus-labs/vanus/pkg/errors"
	errinterceptor "github.com/vanus-labs/vanus/pkg/grpc/interceptor/errors"
)

const (
	vanusConnBypass = "VANUS_CONN_BYPASS"
)

type Conn struct {
	mutex        sync.Mutex
	leader       string
	leaderClient *grpc.ClientConn
	endpoints    []string
	credentials  credentials.TransportCredentials
	grpcConn     map[string]*grpc.ClientConn
	bypass       bool
}

func NewConnection(endpoints []string, credentials credentials.TransportCredentials) *Conn {
	// TODO temporary implement
	v, _ := strconv.ParseBool(os.Getenv(vanusConnBypass))
	log.Info().Strs("endpoints", endpoints).Msg("init Conn")
	return &Conn{
		endpoints:   endpoints,
		grpcConn:    map[string]*grpc.ClientConn{},
		credentials: credentials,
		bypass:      v,
	}
}

func (c *Conn) invoke(ctx context.Context, method string, args, reply interface{}, opts ...grpc.CallOption) error {
	log.Debug(ctx).Str("method", method).
		Str("args", fmt.Sprintf("%v", args)).
		Msg("grpc invoke")
	conn, err := c.makeSureClient(ctx, false)
	if conn == nil || err != nil {
		log.Warn(ctx).Err(err).Msg("not get client for controller")
		return err
	}

	for idx := 1; idx <= 3; idx++ {
		err = conn.Invoke(ctx, method, args, reply, opts...)
		if err != nil {
			log.Debug(ctx).Err(err).Msg("invoke error, try to retry")
		}
		if errors.Is(err, errors.ErrNotReady) {
			time.Sleep(time.Duration(3*idx) * time.Second)
			continue
		} else if isNeedRetry(err) {
			conn, err = c.makeSureClient(ctx, true)
			if conn == nil {
				log.Warn(ctx).Err(err).Msg("not get client when try to renew client")
				return err
			}
		} else {
			return err
		}
	}
	return err
}

func (c *Conn) close() error {
	var err error
	for ip, conn := range c.grpcConn {
		if _err := conn.Close(); _err != nil {
			log.Info().Err(_err).
				Str("peer_address", ip).
				Msg("close grpc connection failed")
			err = errors.Chain(err, _err)
		}
	}
	return err
}

func (c *Conn) makeSureClient(ctx context.Context, renew bool) (*grpc.ClientConn, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.leaderClient == nil || renew {
		if c.bypass {
			c.leaderClient = c.getGRPCConn(ctx, c.endpoints[0])
			return c.leaderClient, nil
		}
		log.Debug(ctx).
			Bool("renew", renew).
			Strs("endpoints", c.endpoints).Msg("try to create connection")
		for _, v := range c.endpoints {
			conn := c.getGRPCConn(ctx, v)
			if conn == nil {
				continue
			}
			pingClient := ctrlpb.NewPingServerClient(conn)
			res, err := pingClient.Ping(context.Background(), &emptypb.Empty{})
			if err != nil {
				log.Info(ctx).Str("address", v).Err(err).Msg("failed to ping controller")
				return nil, errors.ErrNoControllerLeader
			}
			c.leader = res.LeaderAddr
			if v == res.LeaderAddr {
				c.leaderClient = conn
				return conn, nil
			}
			break
		}

		conn := c.getGRPCConn(ctx, c.leader)
		if conn == nil {
			log.Info(ctx).Msg("failed to get Conn")
			return nil, errors.ErrNoControllerLeader
		}
		log.Info(ctx).Str("leader", c.leader).Msg("success to get connection")
		c.leaderClient = conn
	}
	return c.leaderClient, nil
}

func (c *Conn) getGRPCConn(ctx context.Context, addr string) *grpc.ClientConn {
	if addr == "" {
		return nil
	}
	var err error
	conn := c.grpcConn[addr]
	if isConnectionOK(conn) {
		return conn
	} else if conn != nil {
		_ = conn.Close() // make sure it's closed
	}

	opts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTransportCredentials(c.credentials),
		grpc.WithUnaryInterceptor(errinterceptor.UnaryClientInterceptor()),
	}
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()
	conn, err = grpc.DialContext(ctx, addr, opts...)
	if err != nil {
		log.Error().Str("address", addr).Err(err).Msg("failed to dial to controller")
		return nil
	}
	c.grpcConn[addr] = conn
	return conn
}

func isNeedRetry(err error) bool {
	if err == nil {
		return false
	}
	if errors.Is(err, errors.ErrNotLeader) {
		return true
	}
	sts := status.Convert(err)
	if sts == nil {
		return false
	}
	if sts.Code() == codes.Unavailable {
		return true
	}
	return false
}

func isConnectionOK(conn *grpc.ClientConn) bool {
	if conn == nil {
		return false
	}
	return conn.GetState() == connectivity.Idle || conn.GetState() == connectivity.Ready
}
