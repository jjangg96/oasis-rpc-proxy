package controllers

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/figment-networks/oasis-rpc-proxy/connections"
	"github.com/figment-networks/oasis-rpc-proxy/utils"
	"github.com/figment-networks/oasis-rpc-proxy/utils/log"
	"github.com/gin-gonic/gin"
	"github.com/oasislabs/oasis-core/go/common/cbor"
	"github.com/oasislabs/oasis-core/go/consensus/api"
	"github.com/oasislabs/oasis-core/go/consensus/api/transaction"
	"math/big"
	"net/http"
	"reflect"
	"strconv"
)

type GetTransactionsResponse struct {
	Message string        `json:"message"`
	Data    []Transaction `json:"data"`
}

type Transaction struct {
	Success           bool        `json:"success"`
	SignatureVerified bool        `json:"signature_verified"`
	SanityChecked     bool        `json:"sanity_checked"`
	Raw               string      `json:"raw"`
	Hash              string      `json:"hash"`
	PublicKey         string      `json:"public_key"`
	Signature         string      `json:"signature"`
	Nonce             uint64      `json:"nonce"`
	Fee               big.Int     `json:"fee"`
	GasLimit          uint64      `json:"gas_limit"`
	GasPrice          big.Int     `json:"gas_price"`
	Method            string      `json:"method"`
	Body              interface{} `json:"body"`
}

func GetTransactions(c *gin.Context) {
	height, err := (strconv.ParseInt(c.Param("height"), 10, 64))
	if err != nil {
		log.Error("height must be a number", err)
		c.JSON(http.StatusBadRequest, utils.ApiError{Message: "height must be a number"})
		return
	}

	conn, err := connections.GetOasisConn()
	if err != nil {
		log.Error("error connecting to gRPC server", err)
		c.JSON(http.StatusBadRequest, utils.ApiError{Message: "error connecting to gRPC server"})
		return
	}
	defer conn.Close()

	client := api.NewConsensusClient(conn)

	transactions, err := client.GetTransactions(c, height)
	if err != nil {
		log.Error("could not get transactions", err)
		c.JSON(http.StatusBadRequest, utils.ApiError{Message: "could not get transactions"})
		return
	}

	resp := GetTransactionsResponse{Message: "success", Data: []Transaction{}}
	for _, rawTx := range transactions {
		stx := base64.StdEncoding.EncodeToString(rawTx)
		t := Transaction{
			Success:           false,
			SignatureVerified: false,
			SanityChecked:     false,
			Raw:               stx,
		}

		sigTx, err := getSignedTransaction(rawTx, stx, &t)
		if err != nil {
			resp.Data = append(resp.Data, t)
			continue
		}

		tx, err := getTransaction(sigTx, stx, &t)
		if err != nil {
			resp.Data = append(resp.Data, t)
			continue
		}

		if err := sanityCheck(tx, stx, &t); err != nil {
			resp.Data = append(resp.Data, t)
			continue
		}

		bodyType, err := getBodyType(tx, stx)
		if err != nil {
			resp.Data = append(resp.Data, t)
			continue
		}

		v, err := deserialize(tx, bodyType, stx)
		if err != nil {
			resp.Data = append(resp.Data, t)
			continue
		}

		t.Body = v
		resp.Data = append(resp.Data, t)
	}

	c.JSON(http.StatusOK, resp)
}

func getSignedTransaction(rawTx []byte, stx string, t *Transaction) (*transaction.SignedTransaction, error) {
	var sigTx transaction.SignedTransaction
	if err := cbor.Unmarshal(rawTx, &sigTx); err != nil {
		log.Error(fmt.Sprintf("failed to unmarshal signed transaction %s", stx), err)
		return nil, err
	}
	t.Hash = sigTx.Hash().String()
	t.Signature = sigTx.Signature.Signature.String()
	t.PublicKey = sigTx.Signature.PublicKey.String()
	return &sigTx, nil
}

func getTransaction(sigTx *transaction.SignedTransaction, stx string, t *Transaction) (*transaction.Transaction, error) {
	var tx transaction.Transaction
	if err := sigTx.Open(&tx); err != nil {
		log.Error(fmt.Sprintf("failed to verify transaction signature %s", stx), err)
		return nil, err
	}
	t.SignatureVerified = true
	t.Nonce = tx.Nonce
	t.Fee = *tx.Fee.Amount.ToBigInt()
	t.GasLimit = uint64(tx.Fee.Gas)
	t.GasPrice = *tx.Fee.GasPrice().ToBigInt()
	t.Method = string(tx.Method)
	return &tx, nil
}

func sanityCheck(tx *transaction.Transaction, stx string, t *Transaction) error {
	if err := tx.SanityCheck(); err != nil {
		log.Error(fmt.Sprintf("bad transaction %s", stx), err)
		return err
	}
	t.SanityChecked = true
	return nil
}

func getBodyType(tx *transaction.Transaction, stx string) (interface{}, error) {
	bodyType := tx.Method.BodyType()
	if bodyType == nil {
		msg := fmt.Sprintf("unknown method body %s for %s", bodyType, stx)
		err := errors.New(msg)
		log.Error(msg, err)
		return nil, err
	}
	return bodyType, nil
}

func deserialize(tx *transaction.Transaction, bodyType interface{}, stx string) (interface{}, error) {
	v := reflect.New(reflect.TypeOf(bodyType)).Interface()
	if err := cbor.Unmarshal(tx.Body, v); err != nil {
		log.Error(fmt.Sprintf("could not deserialize %s", stx), err)
		return nil, err
	}
	return v, nil
}