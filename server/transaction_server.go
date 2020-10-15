package server

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/figment-networks/oasis-rpc-proxy/client"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/transaction/transactionpb"
	"github.com/figment-networks/oasis-rpc-proxy/mapper"

	"github.com/oasisprotocol/oasis-core/go/common/cbor"
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
	resp := &transactionpb.BroadcastResponse{Submitted: false}

	rawTx, err := base64.StdEncoding.DecodeString(req.GetTxRaw())
	if err != nil {
		return resp, fmt.Errorf("base64 decode failed: %w", err)
	}

	var tx *transaction.SignedTransaction
	if err := cbor.Unmarshal(rawTx, &tx); err != nil {
		return resp, fmt.Errorf("CBOR decode failed: %w", err)
	}

	err = s.client.Consensus.BroadcastTransaction(ctx, tx)
	if err != nil {
		return resp, err
	}

	resp.Submitted = true
	return resp, nil
}
