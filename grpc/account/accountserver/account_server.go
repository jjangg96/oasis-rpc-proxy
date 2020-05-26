package accountserver

import (
	"context"
	"github.com/figment-networks/oasis-rpc-proxy/connections"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/account/accountpb"
	"github.com/figment-networks/oasis-rpc-proxy/mapper"
	"github.com/figment-networks/oasis-rpc-proxy/utils/log"
	"github.com/oasislabs/oasis-core/go/common/crypto/signature"
	"github.com/oasislabs/oasis-core/go/staking/api"
)

type Server interface {
	GetByPublicKey(context.Context, *accountpb.GetByPublicKeyRequest) (*accountpb.GetByPublicKeyResponse, error)
}

type server struct{}

func New() Server {
	return &server{}
}

func (*server) GetByPublicKey(ctx context.Context, req *accountpb.GetByPublicKeyRequest) (*accountpb.GetByPublicKeyResponse, error) {
	conn, err := connections.GetOasisConn()
	if err != nil {
		log.Error("error connecting to gRPC server", err)
		return nil, err
	}
	defer conn.Close()

	client := api.NewStakingClient(conn)

	var pKey signature.PublicKey
	err = pKey.UnmarshalText([]byte(req.PublicKey))
	if err != nil {
		return nil, err
	}
	q := &api.OwnerQuery{
		Height: req.Height,
		Owner:  pKey,
	}
	rawAccount, err := client.AccountInfo(ctx, q)
	if err != nil {
		log.Error("could not get block", err)
		return nil, err
	}

	account := mapper.AccountToPb(*rawAccount)

	return &accountpb.GetByPublicKeyResponse{Account: account}, nil
}
