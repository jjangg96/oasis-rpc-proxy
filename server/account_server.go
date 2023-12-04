package server

import (
	"context"
	"github.com/jjangg96/oasis-rpc-proxy/client"
	"github.com/jjangg96/oasis-rpc-proxy/grpc/account/accountpb"
	"github.com/jjangg96/oasis-rpc-proxy/mapper"
)

type AccountServer interface {
	GetByAddress(context.Context, *accountpb.GetByAddressRequest) (*accountpb.GetByAddressResponse, error)
}

type accountServer struct {
	client *client.Client
}

func NewAccountServer(c *client.Client) AccountServer {
	return &accountServer{
		client: c,
	}
}

func (s *accountServer) GetByAddress(ctx context.Context, req *accountpb.GetByAddressRequest) (*accountpb.GetByAddressResponse, error) {
	rawAccount, err := s.client.Staking.GetAccountByAddress(ctx, req.Address, req.Height)
	if err != nil {
		return nil, err
	}

	account := mapper.AccountToPb(*rawAccount)

	return &accountpb.GetByAddressResponse{Account: account}, nil
}
