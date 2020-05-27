package server

import (
	"context"
	"github.com/figment-networks/oasis-rpc-proxy/client"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/transaction/transactionpb"
	"github.com/figment-networks/oasis-rpc-proxy/mapper"
)

type TransactionServer interface {
	GetByHeight(context.Context, *transactionpb.GetByHeightRequest) (*transactionpb.GetByHeightResponse, error)
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
