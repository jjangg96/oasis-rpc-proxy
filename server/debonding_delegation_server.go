package server

import (
	"context"
	"github.com/jjangg96/oasis-rpc-proxy/client"
	"github.com/jjangg96/oasis-rpc-proxy/grpc/debondingdelegation/debondingdelegationpb"
	"github.com/jjangg96/oasis-rpc-proxy/mapper"
)

type DebondingDelegationServer interface {
	GetByAddress(context.Context, *debondingdelegationpb.GetByAddressRequest) (*debondingdelegationpb.GetByAddressResponse, error)
}

type debondingDelegationServer struct {
	client *client.Client
}

func NewDebondingDelegationServer(c *client.Client) DebondingDelegationServer {
	return &debondingDelegationServer{
		client: c,
	}
}

func (s *debondingDelegationServer) GetByAddress(ctx context.Context, req *debondingdelegationpb.GetByAddressRequest) (*debondingdelegationpb.GetByAddressResponse, error) {
	rawDelegations, err := s.client.Staking.GetDebondingDelegations(ctx, req.GetAddress(), req.GetHeight())
	if err != nil {
		return nil, err
	}

	return &debondingdelegationpb.GetByAddressResponse{DebondingDelegations: mapper.DebondingDelegationToPb(rawDelegations)}, nil
}
