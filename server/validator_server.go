package server

import (
	"context"
	"github.com/jjangg96/oasis-rpc-proxy/client"
	"github.com/jjangg96/oasis-rpc-proxy/grpc/validator/validatorpb"
	"github.com/jjangg96/oasis-rpc-proxy/mapper"
	"github.com/jjangg96/oasis-rpc-proxy/utils/logger"
	registryApi "github.com/oasisprotocol/oasis-core/go/registry/api"
	stakingApi "github.com/oasisprotocol/oasis-core/go/staking/api"
)

type ValidatorServer interface {
	GetByHeight(context.Context, *validatorpb.GetByHeightRequest) (*validatorpb.GetByHeightResponse, error)
}

type validatorServer struct {
	client *client.Client
}

func NewValidatorServer(c *client.Client) ValidatorServer {
	return &validatorServer{
		client: c,
	}
}

func (s *validatorServer) GetByHeight(ctx context.Context, req *validatorpb.GetByHeightRequest) (*validatorpb.GetByHeightResponse, error) {
	rawValidators, err := s.client.Scheduler.GetValidatorsByHeight(ctx, req.Height)
	if err != nil {
		return nil, err
	}

	rawEpochTime, err := s.client.Consensus.GetEpochByHeight(ctx, req.Height)
	if err != nil {
		return nil, err
	}

	var validators []*validatorpb.Validator
	for _, rawValidator := range rawValidators {
		rawNode, err := s.client.Registry.GetNodeById(ctx, rawValidator.ID, req.Height)

		if err == registryApi.ErrNoSuchNode {
			// some validators are missing nodes after the network upgrade on august 6
			logger.Info("skipping validator...", logger.Field("rawValidator.ID", rawValidator.ID.String()))
			continue
		} else if err != nil {
			return nil, err
		}

		address := stakingApi.NewAddress(rawNode.EntityID).String()

		rawAccount, err := s.client.Staking.GetAccountByAddress(ctx, address, req.Height)
		if err != nil {
			return nil, err
		}

		validators = append(validators, mapper.ValidatorToPb(rawValidator, address, rawNode, rawAccount, rawEpochTime))
	}
	return &validatorpb.GetByHeightResponse{Validators: validators}, nil
}
