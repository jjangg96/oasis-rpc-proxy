package server

import (
	"context"
	"github.com/figment-networks/oasis-rpc-proxy/client"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/account/accountpb"
	"github.com/figment-networks/oasis-rpc-proxy/mapper"
)

type AccountServer interface {
	GetByPublicKey(context.Context, *accountpb.GetByPublicKeyRequest) (*accountpb.GetByPublicKeyResponse, error)
}

type accountServer struct{
	client *client.Client
}

func NewAccountServer(c *client.Client) AccountServer {
	return &accountServer{
		client: c,
	}
}

func (s *accountServer) GetByPublicKey(ctx context.Context, req *accountpb.GetByPublicKeyRequest) (*accountpb.GetByPublicKeyResponse, error) {
	rawAccount, err := s.client.Staking.GetAccountByPublicKey(ctx, req.PublicKey, req.Height)
	if err != nil {
		return nil, err
	}

	account := mapper.AccountToPb(*rawAccount)

	return &accountpb.GetByPublicKeyResponse{Account: account}, nil
}
