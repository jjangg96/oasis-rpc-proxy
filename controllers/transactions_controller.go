package controllers

import (
	"github.com/figment-networks/oasis-rpc-proxy/connections"
	"github.com/figment-networks/oasis-rpc-proxy/log"
	"github.com/figment-networks/oasis-rpc-proxy/utils"
	"github.com/gin-gonic/gin"
	"github.com/oasislabs/oasis-core/go/consensus/api"
	"net/http"
	"strconv"
)

type GetTransactionsResponse struct {
	Message string   `json:"message"`
	Data    [][]byte `json:"data"`
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

	c.JSON(http.StatusOK, GetTransactionsResponse{Message: "Success", Data: transactions})
}
