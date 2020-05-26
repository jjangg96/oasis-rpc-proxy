package stateserver

import (
	"context"
	"github.com/figment-networks/oasis-rpc-proxy/client"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/state/statepb"
	"github.com/figment-networks/oasis-rpc-proxy/mapper"
)

type Server interface {
	GetByHeight(context.Context, *statepb.GetByHeightRequest) (*statepb.GetByHeightResponse, error)
}

type server struct{
	client *client.Client
}

func New(c *client.Client) Server {
	return &server{
		client: c,
	}
}

func (s *server) GetByHeight(ctx context.Context, req *statepb.GetByHeightRequest) (*statepb.GetByHeightResponse, error) {
	rawState, err := s.client.Consensus.GetStateByHeight(ctx, req.Height)
	if err != nil {
		return nil, err
	}

	state, err := mapper.StateToPb(rawState)
	if err != nil {
		return nil, err
	}

	return &statepb.GetByHeightResponse{State: state}, nil
}
