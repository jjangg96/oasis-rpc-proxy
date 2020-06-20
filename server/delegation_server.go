package server

import (
	"context"
	"github.com/figment-networks/oasis-rpc-proxy/client"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/delegation/delegationpb"
	"github.com/figment-networks/oasis-rpc-proxy/mapper"
)

type DelegationServer interface {
	GetByAddress(context.Context, *delegationpb.GetByAddressRequest) (*delegationpb.GetByAddressResponse, error)
}

type delegationServer struct {
	client *client.Client
}

func NewDelegationServer(c *client.Client) DelegationServer {
	return &delegationServer{
		client: c,
	}
}

func (s *delegationServer) GetByAddress(ctx context.Context, req *delegationpb.GetByAddressRequest) (*delegationpb.GetByAddressResponse, error) {
	rawDelegations, err := s.client.Staking.GetDelegations(ctx, req.GetAddress(), req.GetHeight())
	if err != nil {
		return nil, err
	}

	return &delegationpb.GetByAddressResponse{Delegations: mapper.DelegationToPb(rawDelegations)}, nil
}

