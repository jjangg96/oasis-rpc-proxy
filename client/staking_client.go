package client

import (
	"context"
	"time"

	"github.com/oasisprotocol/oasis-core/go/staking/api"
	"google.golang.org/grpc"
)

var (
	_ StakingClient = (*stakingClient)(nil)
)

type StakingClient interface {
	GetAccountByAddress(context.Context, string, int64) (*api.Account, error)
	GetDelegations(context.Context, string, int64) (map[api.Address]*api.Delegation, error)
	GetDebondingDelegations(context.Context, string, int64) (map[api.Address][]*api.DebondingDelegation, error)
	GetState(context.Context, int64) (*api.Genesis, error)
	GetEvents(ctx context.Context, height int64) ([]*api.Event, error)
}

func NewStakingClient(conn *grpc.ClientConn) *stakingClient {
	return &stakingClient{
		client: api.NewStakingClient(conn),
	}
}

type stakingClient struct {
	client api.Backend
}

func (c *stakingClient) GetAccountByAddress(ctx context.Context, key string, height int64) (*api.Account, error) {
	defer logRequestDuration(time.Now(), "StakingClient_GetAccountByPublicKey")

	q, err := c.buildOwnerQuery(key, height)
	if err != nil {
		return nil, err
	}
	return c.client.Account(ctx, q)
}

func (c *stakingClient) GetDelegations(ctx context.Context, key string, height int64) (map[api.Address]*api.Delegation, error) {
	defer logRequestDuration(time.Now(), "StakingClient_GetDelegations")

	q, err := c.buildOwnerQuery(key, height)
	if err != nil {
		return nil, err
	}
	return c.client.DelegationsFor(ctx, q)
}

func (c *stakingClient) GetEvents(ctx context.Context, height int64) ([]*api.Event, error) {
	defer logRequestDuration(time.Now(), "StakingClient_GetEvents")

	return c.client.GetEvents(ctx, height)
}

func (c *stakingClient) GetDebondingDelegations(ctx context.Context, key string, height int64) (map[api.Address][]*api.DebondingDelegation, error) {
	defer logRequestDuration(time.Now(), "StakingClient_GetDebondingDelegations")

	q, err := c.buildOwnerQuery(key, height)
	if err != nil {
		return nil, err
	}
	return c.client.DebondingDelegationsFor(ctx, q)
}

func (c *stakingClient) GetState(ctx context.Context, height int64) (*api.Genesis, error) {
	defer logRequestDuration(time.Now(), "StakingClient_GetState")

	return c.client.StateToGenesis(ctx, height)
}

func (c *stakingClient) buildOwnerQuery(key string, height int64) (*api.OwnerQuery, error) {
	address, err := getAddress(key)
	if err != nil {
		return nil, err
	}
	q := &api.OwnerQuery{
		Height: height,
		Owner:  *address,
	}
	return q, nil
}