package controllers

import (
	"github.com/figment-networks/oasis-rpc-proxy/connections"
	"github.com/figment-networks/oasis-rpc-proxy/log"
	"github.com/figment-networks/oasis-rpc-proxy/utils"
	"github.com/gin-gonic/gin"
	"github.com/oasislabs/oasis-core/go/common/cbor"
	"github.com/oasislabs/oasis-core/go/consensus/api"
	tmApi "github.com/oasislabs/oasis-core/go/consensus/tendermint/api"
	"net/http"
	"strconv"
)

type GetBlockResponse struct {
	Message string          `json:"message"`
	Data    tmApi.BlockMeta `json:"data"`
}

func GetBlock(c *gin.Context) {
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

	block, err := client.GetBlock(c, height)
	if err != nil {
		log.Error("could not get block", err)
		c.JSON(http.StatusBadRequest, utils.ApiError{Message: "could not get block"})
		return
	}

	var tmBlockMeta tmApi.BlockMeta
	if err := cbor.Unmarshal(block.Meta, &tmBlockMeta); err != nil {
		log.Error("could not unmarshal block meta", err)
		c.JSON(http.StatusBadRequest, utils.ApiError{Message: "could not unmarshal block meta"})
		return
	}

	c.JSON(http.StatusOK, GetBlockResponse{Message: "Success", Data: tmBlockMeta})
}
