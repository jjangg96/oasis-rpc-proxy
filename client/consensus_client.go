package client

import (
	"context"
	"time"

	"github.com/oasisprotocol/oasis-core/go/consensus/api"
	"github.com/oasisprotocol/oasis-core/go/consensus/api/transaction"
	beaconApi "github.com/oasisprotocol/oasis-core/go/beacon/api"
	genesisApi "github.com/oasisprotocol/oasis-core/go/genesis/api"
	"google.golang.org/grpc"
)

var (
	_ ConsensusClient = (*consensusClient)(nil)
)

type ConsensusClient interface {
	BroadcastTransaction(ctx context.Context, tx *transaction.SignedTransaction) error
	GetBlockByHeight(context.Context, int64) (*api.Block, error)
	GetEpochByHeight(context.Context, int64) (beaconApi.EpochTime, error)
	GetStateByHeight(context.Context, int64) (*genesisApi.Document, error)
	GetTransactionsByHeight(context.Context, int64) ([][]byte, error)
}

func NewConsensusClient(conn *grpc.ClientConn) *consensusClient {
	return &consensusClient{
		client: api.NewConsensusClient(conn),
	}
}

type consensusClient struct {
	client api.ClientBackend
}

func (c *consensusClient) BroadcastTransaction(ctx context.Context, tx *transaction.SignedTransaction) error {
	defer logRequestDuration(time.Now(), "ConsensusClient_BroadcastTransaction")

	return c.client.SubmitTxNoWait(ctx, tx)
}

func (c *consensusClient) GetBlockByHeight(ctx context.Context, h int64) (*api.Block, error) {
	defer logRequestDuration(time.Now(), "ConsensusClient_GetBlockByHeight")

	return c.client.GetBlock(ctx, h)
}

func (c *consensusClient) GetEpochByHeight(ctx context.Context, h int64) (beaconApi.EpochTime, error) {
	defer logRequestDuration(time.Now(), "ConsensusClient_GetEpochByHeight")

	return c.client.Beacon().GetEpoch(ctx, h)
}

func (c *consensusClient) GetStateByHeight(ctx context.Context, h int64) (*genesisApi.Document, error) {
	defer logRequestDuration(time.Now(), "ConsensusClient_GetStateByHeight")

	return c.client.StateToGenesis(ctx, h)
}

func (c *consensusClient) GetTransactionsByHeight(ctx context.Context, h int64) ([][]byte, error) {
	defer logRequestDuration(time.Now(), "ConsensusClient_GetTransactionsByHeight")

	return c.client.GetTransactions(ctx, h)
}
