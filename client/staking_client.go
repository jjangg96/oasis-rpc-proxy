package client

import (
	"context"
	"github.com/oasislabs/oasis-core/go/common/crypto/signature"
	"github.com/oasislabs/oasis-core/go/staking/api"
	"google.golang.org/grpc"
	"time"
)

var (
	_ StakingClient = (*stakingClient)(nil)
)

type StakingClient interface {
	GetAccountByPublicKey(context.Context, string, int64) (*api.Account, error)
	GetDelegations(context.Context, string, int64) (map[signature.PublicKey]*api.Delegation, error)
	GetDebondingDelegations(context.Context, string, int64) (map[signature.PublicKey][]*api.DebondingDelegation, error)
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
	defer logRequestDuration(time.Now(), "StakingClient_GetAccountByPublicKey")

	q, err := c.buildOwnerQuery(key, height)
	if err != nil {
		return nil, err
	}
	return c.client.AccountInfo(ctx, q)
}

func (c *stakingClient) GetDelegations(ctx context.Context, key string, height int64) (map[signature.PublicKey]*api.Delegation, error) {
	defer logRequestDuration(time.Now(), "StakingClient_GetDelegations")

	q, err := c.buildOwnerQuery(key, height)
	if err != nil {
		return nil, err
	}
	return c.client.Delegations(ctx, q)
}

func (c *stakingClient) GetDebondingDelegations(ctx context.Context, key string, height int64) (map[signature.PublicKey][]*api.DebondingDelegation, error) {
	defer logRequestDuration(time.Now(), "StakingClient_GetDebondingDelegations")

	q, err := c.buildOwnerQuery(key, height)
	if err != nil {
		return nil, err
	}
	return c.client.DebondingDelegations(ctx, q)
}

func (c *stakingClient) buildOwnerQuery(key string, height int64) (*api.OwnerQuery, error) {
	pKey, err := getPublicKey(key)
	if err != nil {
		return nil, err
	}
	q := &api.OwnerQuery{
		Height: height,
		Owner:  *pKey,
	}
	return q, nil
}