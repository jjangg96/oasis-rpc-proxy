package chainserver

import (
	"context"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/chain/chainpb"
	"github.com/figment-networks/oasis-rpc-proxy/mappers/chainmapper"
	"github.com/oasislabs/oasis-core/go/genesis/api"
)

type Server interface {
	GetCurrent(context.Context, *chainpb.GetCurrentRequest) (*chainpb.GetCurrentResponse, error)
}

type server struct {
	doc api.Document
}

func New(doc api.Document) Server {
	return &server{
		doc: doc,
	}
}

func (s *server) GetCurrent(ctx context.Context, req *chainpb.GetCurrentRequest) (*chainpb.GetCurrentResponse, error) {
	chain, err := chainmapper.ToPb(s.doc)
	if err != nil {
		return nil, err
	}

	return &chainpb.GetCurrentResponse{
		Chain: chain,
	}, nil
}
