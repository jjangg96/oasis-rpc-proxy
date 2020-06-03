package server

import (
	"context"
	"github.com/figment-networks/oasis-rpc-proxy/client"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/debondingdelegation/debondingdelegationpb"
	"github.com/figment-networks/oasis-rpc-proxy/mapper"
)

type DebondingDelegationServer interface {
	GetByPublicKey(context.Context, *debondingdelegationpb.GetByPublicKeyRequest) (*debondingdelegationpb.GetByPublicKeyResponse, error)
}

type debondingDelegationServer struct {
	client *client.Client
}

func NewDebondingDelegationServer(c *client.Client) DebondingDelegationServer {
	return &debondingDelegationServer{
		client: c,
	}
}

func (s *debondingDelegationServer) GetByPublicKey(ctx context.Context, req *debondingdelegationpb.GetByPublicKeyRequest) (*debondingdelegationpb.GetByPublicKeyResponse, error) {
	rawDelegations, err := s.client.Staking.GetDebondingDelegations(ctx, req.GetPublicKey(), req.GetHeight())
	if err != nil {
		return nil, err
	}

	return &debondingdelegationpb.GetByPublicKeyResponse{DebondingDelegations: mapper.DebondingDelegationToPb(rawDelegations)}, nil
}
