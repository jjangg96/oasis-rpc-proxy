package client

import (
	"context"
	"github.com/oasisprotocol/oasis-core/go/consensus/api"
	genesisApi "github.com/oasisprotocol/oasis-core/go/genesis/api"
	"google.golang.org/grpc"
	"time"
)

var (
	_ ConsensusClient = (*consensusClient)(nil)
)

type ConsensusClient interface {
	GetBlockByHeight(context.Context, int64) (*api.Block, error)
	GetStateByHeight(context.Context, int64) (*genesisApi.Document, error)
	GetTransactionsByHeight(context.Context, int64) ([][]byte, error)
}

func NewConsensusClient(conn *grpc.ClientConn) ConsensusClient {
	return &consensusClient{
		client: api.NewConsensusClient(conn),
	}
}

type consensusClient struct {
	client api.ClientBackend
}

func (c *consensusClient) GetBlockByHeight(ctx context.Context, h int64) (*api.Block, error) {
	defer logRequestDuration(time.Now(), "ConsensusClient_GetBlockByHeight")

	return c.client.GetBlock(ctx, h)
}

func (c *consensusClient) GetStateByHeight(ctx context.Context, h int64) (*genesisApi.Document, error) {
	defer logRequestDuration(time.Now(), "ConsensusClient_GetStateByHeight")

	return c.client.StateToGenesis(ctx, h)
}

func (c *consensusClient) GetTransactionsByHeight(ctx context.Context, h int64) ([][]byte, error) {
	defer logRequestDuration(time.Now(), "ConsensusClient_GetTransactionsByHeight")

	return c.client.GetTransactions(ctx, h)
}
