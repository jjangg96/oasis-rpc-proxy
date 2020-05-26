package transactionserver

import (
	"context"
	"github.com/figment-networks/oasis-rpc-proxy/connections"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/transaction/transactionpb"
	"github.com/figment-networks/oasis-rpc-proxy/mapper"
	"github.com/figment-networks/oasis-rpc-proxy/utils/log"
	"github.com/oasislabs/oasis-core/go/consensus/api"
)

type Server interface {
	GetByHeight(context.Context, *transactionpb.GetByHeightRequest) (*transactionpb.GetByHeightResponse, error)
}

type server struct{}

func New() Server {
	return &server{}
}

func (*server) GetByHeight(ctx context.Context, req *transactionpb.GetByHeightRequest) (*transactionpb.GetByHeightResponse, error) {
	conn, err := connections.GetOasisConn()
	if err != nil {
		log.Error("error connecting to gRPC server", err)
		return nil, err
	}
	defer conn.Close()

	client := api.NewConsensusClient(conn)

	rawTxs, err := client.GetTransactions(ctx, req.Height)
	if err != nil {
		log.Error("could not get transactions", err)
		return nil, err
	}

	var transactions []*transactionpb.Transaction
	for _, rawTx := range rawTxs {
		transactions = append(transactions, mapper.TransactionToPb(rawTx))
	}
	return &transactionpb.GetByHeightResponse{Transactions: transactions}, nil
}
