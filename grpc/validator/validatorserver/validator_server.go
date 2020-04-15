package validatorserver

import (
	"context"
	"github.com/figment-networks/oasis-rpc-proxy/connections"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/validator/validatorpb"
	"github.com/figment-networks/oasis-rpc-proxy/log"
	"github.com/figment-networks/oasis-rpc-proxy/mappers/validatormapper"
	registryApi "github.com/oasislabs/oasis-core/go/registry/api"
	"github.com/oasislabs/oasis-core/go/scheduler/api"
)

type Server interface {
	GetByHeight(context.Context, *validatorpb.GetByHeightRequest) (*validatorpb.GetByHeightResponse, error)
}

type server struct{}

func New() Server {
	return &server{}
}

func (*server) GetByHeight(ctx context.Context, req *validatorpb.GetByHeightRequest) (*validatorpb.GetByHeightResponse, error) {
	conn, err := connections.GetOasisConn()
	if err != nil {
		log.Error("error connecting to gRPC server", err)
		return nil, err
	}
	defer conn.Close()

	schedulerClient := api.NewSchedulerClient(conn)
	registryClient := registryApi.NewRegistryClient(conn)

	rawValidators, err := schedulerClient.GetValidators(ctx, req.Height)
	if err != nil {
		log.Error("could not get list of validators", err)
		return nil, err
	}

	var validators []*validatorpb.Validator
	for _, rawValidator := range rawValidators {
		node, rErr := registryClient.GetNode(ctx, &registryApi.IDQuery{
			Height: req.Height,
			ID:     rawValidator.ID,
		})
		if rErr != nil {
			log.Error("could not get node details", rErr)
		}

		validators = append(validators, validatormapper.ToPb(rawValidator, node))
	}
	return &validatorpb.GetByHeightResponse{Validators: validators}, nil
}
