package server

import (
	"context"

	"github.com/jjangg96/oasis-rpc-proxy/client"
	"github.com/jjangg96/oasis-rpc-proxy/grpc/state/statepb"
	"github.com/jjangg96/oasis-rpc-proxy/mapper"
)

type StateServer interface {
	GetByHeight(context.Context, *statepb.GetByHeightRequest) (*statepb.GetByHeightResponse, error)
	GetStakingByHeight(context.Context, *statepb.GetStakingByHeightRequest) (*statepb.GetStakingByHeightResponse, error)
}

type stateServer struct {
	client *client.Client
}

func NewStateServer(c *client.Client) StateServer {
	return &stateServer{
		client: c,
	}
}

func (s *stateServer) GetByHeight(ctx context.Context, req *statepb.GetByHeightRequest) (*statepb.GetByHeightResponse, error) {
	rawState, err := s.client.Consensus.GetStateByHeight(ctx, req.Height)
	if err != nil {
		return nil, err
	}

	state, err := mapper.StateToPb(rawState)
	if err != nil {
		return nil, err
	}

	return &statepb.GetByHeightResponse{State: state}, nil
}

func (s *stateServer) GetStakingByHeight(ctx context.Context, req *statepb.GetStakingByHeightRequest) (*statepb.GetStakingByHeightResponse, error) {
	if req.OmitAccountsAndDelegations {
		totalSupply, commonPool, params, err := s.client.Staking.GetStatus(ctx, req.GetHeight())
		if err != nil {
			return nil, err
		}

		return &statepb.GetStakingByHeightResponse{Staking: &statepb.Staking{
			TotalSupply: totalSupply.ToBigInt().Bytes(),
			CommonPool:  commonPool.ToBigInt().Bytes(),
			Parameters:  mapper.ConsensusParametersToStakingParameters(*params),
		}}, nil
	} else {
		rawState, err := s.client.Staking.GetState(ctx, req.GetHeight())
		if err != nil {
			return nil, err
		}

		return &statepb.GetStakingByHeightResponse{Staking: mapper.StakingToPb(*rawState)}, nil
	}
}
