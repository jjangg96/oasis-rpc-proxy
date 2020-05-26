package validatorserver

import (
	"context"
	"github.com/figment-networks/oasis-rpc-proxy/client"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/validator/validatorpb"
	"github.com/figment-networks/oasis-rpc-proxy/mapper"
)

type Server interface {
	GetByHeight(context.Context, *validatorpb.GetByHeightRequest) (*validatorpb.GetByHeightResponse, error)
}

type server struct{
	client *client.Client
}

func New(c *client.Client) Server {
	return &server{
		client: c,
	}
}

func (s *server) GetByHeight(ctx context.Context, req *validatorpb.GetByHeightRequest) (*validatorpb.GetByHeightResponse, error) {
	rawValidators, err := s.client.Scheduler.GetValidatorsByHeight(ctx, req.Height)
	if err != nil {
		return nil, err
	}

	var validators []*validatorpb.Validator
	for _, rawValidator := range rawValidators {
		node, err := s.client.Registry.GeNodeById(ctx, rawValidator.ID.String(), req.Height)
		if err != nil {
			return nil, err
		}

		validators = append(validators, mapper.ValidatorToPb(rawValidator, node))
	}
	return &validatorpb.GetByHeightResponse{Validators: validators}, nil
}
