package cluster

import (
	"context"
	"time"

	ctrlpb "github.com/vanus-labs/vanus/proto/pkg/controller"

	"github.com/vanus-labs/vanus/pkg/cluster/raw_client"
)

type segmentService struct {
	client ctrlpb.SegmentControllerClient
}

func newSegmentService(cc *raw_client.Conn) SegmentService {
	return &segmentService{client: raw_client.NewSegmentClient(cc)}
}

func (es *segmentService) RawClient() ctrlpb.SegmentControllerClient {
	return es.client
}

func (es *segmentService) RegisterHeartbeat(ctx context.Context, interval time.Duration, reqFunc func() interface{}) error {
	return raw_client.RegisterHeartbeat(ctx, interval, es.client, reqFunc)
}
