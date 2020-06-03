package server

import (
	"context"
	"github.com/figment-networks/oasis-rpc-proxy/client"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/delegation/delegationpb"
	"github.com/figment-networks/oasis-rpc-proxy/mapper"
)

type DelegationServer interface {
	GetByPublicKey(context.Context, *delegationpb.GetByPublicKeyRequest) (*delegationpb.GetByPublicKeyResponse, error)
}

type delegationServer struct {
	client *client.Client
}

func NewDelegationServer(c *client.Client) DelegationServer {
	return &delegationServer{
		client: c,
	}
}

func (s *delegationServer) GetByPublicKey(ctx context.Context, req *delegationpb.GetByPublicKeyRequest) (*delegationpb.GetByPublicKeyResponse, error) {
	rawDelegations, err := s.client.Staking.GetDelegations(ctx, req.GetPublicKey(), req.GetHeight())
	if err != nil {
		return nil, err
	}

	return &delegationpb.GetByPublicKeyResponse{Delegations: mapper.DelegationToPb(rawDelegations)}, nil
}

