package server

import (
	"context"

	"github.com/figment-networks/oasis-rpc-proxy/client"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/event/eventpb"
	"github.com/figment-networks/oasis-rpc-proxy/mapper"
)

type EventServer interface {
	GetEscrowEventsByHeight(context.Context, *eventpb.GetEscrowEventsByHeightRequest) (*eventpb.GetEscrowEventsByHeightResponse, error)
	GetTransferEventsByHeight(context.Context, *eventpb.GetEscrowEventsByHeightRequest) (*eventpb.GetEscrowEventsByHeightResponse, error)
}

type eventServer struct {
	client *client.Client
}

func NewEventServer(c *client.Client) EventServer {
	return &eventServer{
		client: c,
	}
}

func (s *eventServer) GetEscrowEventsByHeight(ctx context.Context, req *eventpb.GetEscrowEventsByHeightRequest) (*eventpb.GetEscrowEventsByHeightResponse, error) {
	rawEvents, err := s.client.Staking.GetEvents(ctx, req.Height)
	if err != nil {
		return nil, err
	}

	var add []*eventpb.AddEscrowEvent
	var take []*eventpb.TakeEscrowEvent

	for _, rawEvent := range rawEvents {
		if rawEvent.Escrow == nil {
			continue
		}

		if rawEvent.Escrow.Add != nil {
			add = append(add, mapper.AddEscrowEventToPb(rawEvent.Escrow.Add))
		} else if rawEvent.Escrow.Take != nil {
			take = append(take, mapper.TakeEscrowEventToPb(rawEvent.Escrow.Take))
		}
	}

	return &eventpb.GetEscrowEventsByHeightResponse{
		Events: &eventpb.EscrowEvents{
			Add:  add,
			Take: take,
		},
	}, nil
}

func (s *eventServer) GetTransferEventsByHeight(ctx context.Context, req *eventpb.GetTransferEventsByHeightRequest) (*eventpb.GetTransferEventsByHeightResponse, error) {
	rawEvents, err := s.client.Staking.GetEvents(ctx, req.Height)
	if err != nil {
		return nil, err
	}

	var events []*eventpb.TransferEvent

	for _, rawEvent := range rawEvents {
		if rawEvent.Transfer != nil {
			events = append(events, mapper.TransferEventToPb(rawEvent.Transfer))
		}
	}

	return &eventpb.GetTransferEventsByHeightResponse{
		Events: events,
	}, nil
}
