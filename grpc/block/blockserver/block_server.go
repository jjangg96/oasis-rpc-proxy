package blockserver

import (
	"context"
	"github.com/figment-networks/oasis-rpc-proxy/connections"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/block/blockpb"
	"github.com/figment-networks/oasis-rpc-proxy/mappers/blockmapper"
	"github.com/figment-networks/oasis-rpc-proxy/utils/log"
	"github.com/oasislabs/oasis-core/go/consensus/api"
)

type Server interface {
	GetByHeight(context.Context, *blockpb.GetByHeightRequest) (*blockpb.GetByHeightResponse, error)
}

type server struct{}

func New() Server {
	return &server{}
}

func (*server) GetByHeight(ctx context.Context, req *blockpb.GetByHeightRequest) (*blockpb.GetByHeightResponse, error) {
	conn, err := connections.GetOasisConn()
	if err != nil {
		log.Error("error connecting to gRPC server", err)
		return nil, err
	}
	defer conn.Close()

	client := api.NewConsensusClient(conn)

	rawBlock, err := client.GetBlock(ctx, req.Height)
	if err != nil {
		log.Error("could not get block", err)
		return nil, err
	}

	block, err := blockmapper.ToPb(*rawBlock)
	if err != nil {
		return nil, err
	}

	return &blockpb.GetByHeightResponse{Block: block}, nil
}
