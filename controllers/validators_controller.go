package controllers

import (
	"github.com/figment-networks/oasis-rpc-proxy/connections"
	"github.com/figment-networks/oasis-rpc-proxy/utils"
	"github.com/figment-networks/oasis-rpc-proxy/utils/log"
	"github.com/gin-gonic/gin"
	"github.com/oasislabs/oasis-core/go/common/crypto/signature"
	"github.com/oasislabs/oasis-core/go/common/node"
	tmcrypto "github.com/oasislabs/oasis-core/go/consensus/tendermint/crypto"
	registryApi "github.com/oasislabs/oasis-core/go/registry/api"
	"github.com/oasislabs/oasis-core/go/scheduler/api"
	"net/http"
	"strconv"
)

type GetValidatorsResponse struct {
	Message string      `json:"message"`
	Data    []Validator `json:"data"`
}

type Validator struct {
	ID          signature.PublicKey `json:"id"`
	VotingPower int64               `json:"voting_power"`
	Address     string              `json:"address"`
	Node        node.Node           `json:"node"`
}

func GetValidators(c *gin.Context) {
	height, err := strconv.ParseInt(c.Param("height"), 10, 64)
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

	client := api.NewSchedulerClient(conn)

	rawValidators, err := client.GetValidators(c, height)
	if err != nil {
		log.Error("could not get list of validators", err)
		c.JSON(http.StatusBadRequest, utils.ApiError{Message: "could not get list of validators"})
		return
	}

	var validators []Validator
	for _, validator := range rawValidators {
		// Get validator (node) details
		client2 := registryApi.NewRegistryClient(conn)
		node, err2 := client2.GetNode(c, &registryApi.IDQuery{
			Height: height,
			ID:     validator.ID,
		})
		if err2 != nil {
			log.Error("could not get node details", err2)
		}

		cID := node.Consensus.ID
		tmAddr := tmcrypto.PublicKeyToTendermint(&cID).Address().String()

		validators = append(validators, Validator{
			ID:          validator.ID,
			Address:     tmAddr,
			VotingPower: validator.VotingPower,
			Node:        *node,
		})
	}

	c.JSON(http.StatusOK, GetValidatorsResponse{Message: "Success", Data: validators})
}
