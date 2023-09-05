package cluster

import (
	ctrlpb "github.com/vanus-labs/vanus/proto/pkg/controller"

	"github.com/vanus-labs/vanus/pkg/cluster/raw_client"
)

type eventlogService struct {
	client ctrlpb.EventlogControllerClient
}

func newEventlogService(cc *raw_client.Conn) EventlogService {
	return &eventlogService{client: raw_client.NewEventlogClient(cc)}
}

func (es *eventlogService) RawClient() ctrlpb.EventlogControllerClient {
	return es.client
}
