package accountserver

import (
	"context"
	"github.com/figment-networks/oasis-rpc-proxy/client"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/account/accountpb"
	"github.com/figment-networks/oasis-rpc-proxy/mapper"
)

type Server interface {
	GetByPublicKey(context.Context, *accountpb.GetByPublicKeyRequest) (*accountpb.GetByPublicKeyResponse, error)
}

type server struct{
	client *client.Client
}

func New(c *client.Client) Server {
	return &server{
		client: c,
	}
}

func (s *server) GetByPublicKey(ctx context.Context, req *accountpb.GetByPublicKeyRequest) (*accountpb.GetByPublicKeyResponse, error) {
	rawAccount, err := s.client.Staking.GetAccountByPublicKey(ctx, req.PublicKey, req.Height)
	if err != nil {
		return nil, err
	}

	account := mapper.AccountToPb(*rawAccount)

	return &accountpb.GetByPublicKeyResponse{Account: account}, nil
}
