package mapper

import (
	"encoding/base64"
	"errors"
	"fmt"
	"reflect"

	"github.com/jjangg96/oasis-rpc-proxy/grpc/transaction/transactionpb"
	"github.com/jjangg96/oasis-rpc-proxy/utils/logger"
	"github.com/oasisprotocol/oasis-core/go/common/cbor"
	"github.com/oasisprotocol/oasis-core/go/consensus/api/transaction"
)

func TransactionToPb(rawTx []byte) *transactionpb.Transaction {
	stx := base64.StdEncoding.EncodeToString(rawTx)
	t := transactionpb.Transaction{
		Success:           false,
		SignatureVerified: false,
		SanityChecked:     false,
	}

	sigTx, err := ToSignedTransaction(rawTx, stx, &t)
	if err != nil {
		logger.Error(err)
		return &t
	}

	tx, err := ToTransaction(sigTx, stx, &t)
	if err != nil {
		logger.Error(err)
		return &t
	}

	if err := sanityCheck(tx, stx, &t); err != nil {
		logger.Error(err)
		return &t
	}

	return &t
}

func ToSignedTransaction(rawTx []byte, stx string, t *transactionpb.Transaction) (*transaction.SignedTransaction, error) {
	var sigTx transaction.SignedTransaction
	if err := cbor.Unmarshal(rawTx, &sigTx); err != nil {
		return nil, err
	}
	t.Hash = sigTx.Hash().String()
	t.Signature = sigTx.Signature.Signature.String()
	t.PublicKey = sigTx.Signature.PublicKey.String()
	return &sigTx, nil
}

func ToTransaction(sigTx *transaction.SignedTransaction, stx string, t *transactionpb.Transaction) (*transaction.Transaction, error) {
	var tx transaction.Transaction
	t.SignatureVerified = true

	if err := sigTx.Open(&tx); err != nil {
		t.SignatureVerified = false

		// Unmarshall blob if signature verification failed
		if err := cbor.Unmarshal(sigTx.Blob, &tx); err != nil {
			return nil, err
		}
	}

	t.Nonce = tx.Nonce
	t.Fee = tx.Fee.Amount.ToBigInt().Bytes()
	t.GasLimit = uint64(tx.Fee.Gas)
	t.GasPrice = tx.Fee.GasPrice().ToBigInt().Bytes()
	t.Method = string(tx.Method)
	return &tx, nil
}

func sanityCheck(tx *transaction.Transaction, stx string, t *transactionpb.Transaction) error {
	if err := tx.SanityCheck(); err != nil {
		return err
	}
	t.SanityChecked = true
	return nil
}

func getBodyType(tx *transaction.Transaction, stx string) (interface{}, error) {
	bodyType := tx.Method.BodyType()
	if bodyType == nil {
		return nil, errors.New(fmt.Sprintf("unknown method body %s for %s", bodyType, stx))
	}
	return bodyType, nil
}

func deserialize(tx *transaction.Transaction, bodyType interface{}, stx string) (interface{}, error) {
	v := reflect.New(reflect.TypeOf(bodyType)).Interface()
	if err := cbor.Unmarshal(tx.Body, v); err != nil {
		return nil, err
	}
	return v, nil
}
