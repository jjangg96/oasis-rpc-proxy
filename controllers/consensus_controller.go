package controllers

import (
	genesisApi "github.com/oasislabs/oasis-core/go/genesis/api"
	"net/http"
	"strconv"

	"github.com/figment-networks/oasis-rpc-proxy/connections"
	"github.com/figment-networks/oasis-rpc-proxy/utils"
	"github.com/figment-networks/oasis-rpc-proxy/utils/log"
	"github.com/gin-gonic/gin"
	consensusApi "github.com/oasislabs/oasis-core/go/consensus/api"
)

type GetConsensusStateResponse struct {
	Message string          `json:"message"`
	Data    *genesisApi.Document `json:"data"`
}

func GetConsensusState(c *gin.Context) {
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

	client := consensusApi.NewConsensusClient(conn)

	state, err := client.StateToGenesis(c, height)
	if err != nil {
		log.Error("could not get block", err)
		c.JSON(http.StatusBadRequest, utils.ApiError{Message: "could not get block"})
		return
	}

	c.JSON(http.StatusOK, GetConsensusStateResponse{Message: "Success", Data: state})
}