package server

import (
	"context"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/chain/chainpb"
	"github.com/figment-networks/oasis-rpc-proxy/mapper"
	"github.com/oasislabs/oasis-core/go/genesis/api"
)

type ChainServer interface {
	GetCurrent(context.Context, *chainpb.GetCurrentRequest) (*chainpb.GetCurrentResponse, error)
}

type chainServer struct {
	doc *api.Document
}

func NewChainServer(doc *api.Document) ChainServer {
	return &chainServer{
		doc: doc,
	}
}

func (s *chainServer) GetCurrent(ctx context.Context, req *chainpb.GetCurrentRequest) (*chainpb.GetCurrentResponse, error) {
	chain, err := mapper.ChainToPb(s.doc)
	if err != nil {
		return nil, err
	}

	return &chainpb.GetCurrentResponse{
		Chain: chain,
	}, nil
}
