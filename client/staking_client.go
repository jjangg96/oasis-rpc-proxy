package client

import (
	"context"
	"github.com/oasislabs/oasis-core/go/staking/api"
	"google.golang.org/grpc"
)

var (
	_ StakingClient = (*stakingClient)(nil)
)

type StakingClient interface {
	GetAccountByPublicKey(context.Context, string, int64) (*api.Account, error)
}

func NewStakingClient(conn *grpc.ClientConn) *stakingClient {
	return &stakingClient{
		client: api.NewStakingClient(conn),
	}
}

type stakingClient struct {
	client api.Backend
}

func (c *stakingClient) GetAccountByPublicKey(ctx context.Context, key string, height int64) (*api.Account, error) {
	pKey, err := getPublicKey(key)
	if err != nil {
		return nil, err
	}
	q := &api.OwnerQuery{
		Height: height,
		Owner:  *pKey,
	}
	return c.client.AccountInfo(ctx, q)
}
