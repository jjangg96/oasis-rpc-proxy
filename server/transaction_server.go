package server

import (
	"context"
	"fmt"

	"github.com/figment-networks/oasis-rpc-proxy/client"
	"github.com/figment-networks/oasis-rpc-proxy/grpc/transaction/transactionpb"
	"github.com/figment-networks/oasis-rpc-proxy/mapper"
	"github.com/oasisprotocol/oasis-core/go/common/crypto/signature"
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
	reqPublicKey := req.GetSignature().GetPublicKey()
	if len(reqPublicKey) != signature.PublicKeySize {
		return nil, fmt.Errorf("publickey is invalid, must be type [%v]byte", signature.PublicKeySize)
	}

	reqSignature := req.GetSignature().GetRawSignature()
	if len(reqSignature) != signature.SignatureSize {
		return nil, fmt.Errorf("rawsignature is invalid, must be type [%v]byte", signature.SignatureSize)
	}

	var publicKey [signature.PublicKeySize]byte
	for i, val := range reqPublicKey {
		publicKey[i] = val
	}

	var rawSignature [signature.SignatureSize]byte
	for i, val := range reqSignature {
		rawSignature[i] = val
	}

	tx := &transaction.SignedTransaction{
		signature.Signed{
			Blob: req.GetUntrustedRawValue(),
			Signature: signature.Signature{
				PublicKey: publicKey,
				Signature: rawSignature,
			},
		},
	}

	err := s.client.Consensus.BroadcastTransaction(ctx, tx)
	if err != nil {
		return nil, err
	}

	return &transactionpb.BroadcastResponse{Success: true}, nil
}
