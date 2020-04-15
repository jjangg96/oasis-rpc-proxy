package stateserver

import (
	"context"
	"github.com/figment-networks/oasis-rpc-proxy/connections"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/state/statepb"
	"github.com/figment-networks/oasis-rpc-proxy/log"
	"github.com/figment-networks/oasis-rpc-proxy/mappers/statemapper"
	consensusApi "github.com/oasislabs/oasis-core/go/consensus/api"
)

type Server interface {
	GetByHeight(context.Context, *statepb.GetByHeightRequest) (*statepb.GetByHeightResponse, error)
}

type server struct{}

func New() Server {
	return &server{}
}

func (*server) GetByHeight(ctx context.Context, req *statepb.GetByHeightRequest) (*statepb.GetByHeightResponse, error) {
	conn, err := connections.GetOasisConn()
	if err != nil {
		log.Error("error connecting to gRPC server", err)
		return nil, err
	}
	defer conn.Close()

	client := consensusApi.NewConsensusClient(conn)

	rawState, err := client.StateToGenesis(ctx, req.Height)
	if err != nil {
		log.Error("could not get block", err)
		return nil, err
	}

	state, err := statemapper.ToPb(rawState)
	if err != nil {
		return nil, err
	}

	return &statepb.GetByHeightResponse{State: state}, nil
}
