package server

import (
	"context"

	"github.com/figment-networks/oasis-rpc-proxy/client"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/event/eventpb"
	"github.com/figment-networks/oasis-rpc-proxy/mapper"
	"github.com/oasisprotocol/oasis-core/go/staking/api"
)

type EventServer interface {
	GetAddEscrowEventsByHeight(context.Context, *eventpb.GetAddEscrowEventsByHeightRequest) (*eventpb.GetAddEscrowEventsByHeightResponse, error)
}

type eventServer struct {
	client *client.Client
}

func NewEventServer(c *client.Client) EventServer {
	return &eventServer{
		client: c,
	}
}

func (s *eventServer) GetAddEscrowEventsByHeight(ctx context.Context, req *eventpb.GetAddEscrowEventsByHeightRequest) (*eventpb.GetAddEscrowEventsByHeightResponse, error) {
	rawEvents, err := s.client.Staking.GetEvents(ctx, req.Height)
	if err != nil {
		return nil, err
	}

	var events []*eventpb.AddEscrowEvent
	for _, rawEvent := range rawEvents {
		events = append(events, mapper.EventToPb(rawEvent))
	}

	return &eventpb.GetAddEscrowEventsByHeightResponse{
		Events:            events,
		CommonPoolAddress: api.CommonPoolAddress.String(),
	}, nil
}
