package server

import (
	"context"
	"github.com/figment-networks/oasis-rpc-proxy/client"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/block/blockpb"
	"github.com/figment-networks/oasis-rpc-proxy/mapper"
)

type BlockServer interface {
	GetByHeight(context.Context, *blockpb.GetByHeightRequest) (*blockpb.GetByHeightResponse, error)
}

type blockServer struct {
	client *client.Client
}

func NewBlockServer(c *client.Client) BlockServer {
	return &blockServer{
		client: c,
	}
}

func (s *blockServer) GetByHeight(ctx context.Context, req *blockpb.GetByHeightRequest) (*blockpb.GetByHeightResponse, error) {
	rawBlock, err := s.client.Consensus.GetBlockByHeight(ctx, req.Height)
	if err != nil {
		return nil, err
	}

	block, err := mapper.BlockToPb(*rawBlock)
	if err != nil {
		return nil, err
	}

	return &blockpb.GetByHeightResponse{Block: block}, nil
}
