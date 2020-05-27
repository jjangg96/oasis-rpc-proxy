package server

import (
	"context"
	"github.com/figment-networks/oasis-rpc-proxy/client"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/state/statepb"
	"github.com/figment-networks/oasis-rpc-proxy/mapper"
)

type StateServer interface {
	GetByHeight(context.Context, *statepb.GetByHeightRequest) (*statepb.GetByHeightResponse, error)
}

type stateServer struct{
	client *client.Client
}

func NewStateServer(c *client.Client) StateServer {
	return &stateServer{
		client: c,
	}
}

func (s *stateServer) GetByHeight(ctx context.Context, req *statepb.GetByHeightRequest) (*statepb.GetByHeightResponse, error) {
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
