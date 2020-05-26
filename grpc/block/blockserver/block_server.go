package blockserver

import (
	"context"
	"github.com/figment-networks/oasis-rpc-proxy/client"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/block/blockpb"
	"github.com/figment-networks/oasis-rpc-proxy/mapper"
)

type Server interface {
	GetByHeight(context.Context, *blockpb.GetByHeightRequest) (*blockpb.GetByHeightResponse, error)
}

type server struct {
	client *client.Client
}

func New(c *client.Client) Server {
	return &server{
		client: c,
	}
}

func (s *server) GetByHeight(ctx context.Context, req *blockpb.GetByHeightRequest) (*blockpb.GetByHeightResponse, error) {
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
