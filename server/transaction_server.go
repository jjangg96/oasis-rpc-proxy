package server

import (
	"context"
	"encoding/json"

	"github.com/figment-networks/oasis-rpc-proxy/client"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/transaction/transactionpb"
	"github.com/figment-networks/oasis-rpc-proxy/mapper"
	"github.com/oasisprotocol/oasis-core/go/consensus/api/transaction"
)

type TransactionServer interface {
	GetByHeight(context.Context, *transactionpb.GetByHeightRequest) (*transactionpb.GetByHeightResponse, error)
	Broadcast(ctx context.Context, req *transactionpb.BroadcastRequest) (*transactionpb.BroadcastResponse, error)
}

type transactionServer struct {
	client *client.Client
}

func NewTransactionServer(c *client.Client) TransactionServer {
	return &transactionServer{
		client: c,
	}
}

func (s *transactionServer) GetByHeight(ctx context.Context, req *transactionpb.GetByHeightRequest) (*transactionpb.GetByHeightResponse, error) {
	rawTxs, err := s.client.Consensus.GetTransactionsByHeight(ctx, req.Height)
	if err != nil {
		return nil, err
	}

	var transactions []*transactionpb.Transaction
	for _, rawTx := range rawTxs {
		transactions = append(transactions, mapper.TransactionToPb(rawTx))
	}
	return &transactionpb.GetByHeightResponse{Transactions: transactions}, nil
}

func (s *transactionServer) Broadcast(ctx context.Context, req *transactionpb.BroadcastRequest) (*transactionpb.BroadcastResponse, error) {
	var tx *transaction.SignedTransaction
	if err := json.Unmarshal([]byte(req.GetTxRaw()), &tx); err != nil {
		return nil, err
	}

	err := s.client.Consensus.BroadcastTransaction(ctx, tx)
	if err != nil {
		return nil, err
	}

	return &transactionpb.BroadcastResponse{Success: true}, nil
}
