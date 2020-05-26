package transactionserver

import (
	"context"
	"github.com/figment-networks/oasis-rpc-proxy/client"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/transaction/transactionpb"
	"github.com/figment-networks/oasis-rpc-proxy/mapper"
)

type Server interface {
	GetByHeight(context.Context, *transactionpb.GetByHeightRequest) (*transactionpb.GetByHeightResponse, error)
}

type server struct {
	client *client.Client
}

func New(c *client.Client) Server {
	return &server{
		client: c,
	}
}

func (s *server) GetByHeight(ctx context.Context, req *transactionpb.GetByHeightRequest) (*transactionpb.GetByHeightResponse, error) {
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
