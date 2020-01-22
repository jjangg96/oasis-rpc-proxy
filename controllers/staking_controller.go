package controllers

import (
	"github.com/figment-networks/oasis-rpc-proxy/connections"
	"github.com/figment-networks/oasis-rpc-proxy/log"
	"github.com/figment-networks/oasis-rpc-proxy/utils"
	"github.com/gin-gonic/gin"
	"github.com/oasislabs/oasis-core/go/common/crypto/signature"
	"github.com/oasislabs/oasis-core/go/staking/api"
	"net/http"
	"strconv"
)

func GetTotalSupply(c *gin.Context) {
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

	client := api.NewStakingClient(conn)

	totalSupply, err := client.TotalSupply(c, height)
	if err != nil {
		log.Error("could not get total supply", err)
		c.JSON(http.StatusBadRequest, utils.ApiError{Message: "could not get total supply"})
		return
	}

	c.JSON(http.StatusOK, utils.ApiResponse{Message: "Success", Data: totalSupply})
}

func GetCommonPool(c *gin.Context) {
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

	client := api.NewStakingClient(conn)

	commonPool, err := client.CommonPool(c, height)
	if err != nil {
		log.Error("could not get common pool", err)
		c.JSON(http.StatusBadRequest, utils.ApiError{Message: "could not get common pool"})
		return
	}

	c.JSON(http.StatusOK, utils.ApiResponse{Message: "Success", Data: commonPool})
}

func GetAccounts(c *gin.Context) {
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

	client := api.NewStakingClient(conn)

	accounts, err := client.Accounts(c, height)
	if err != nil {
		log.Error("could not get accounts", err)
		c.JSON(http.StatusBadRequest, utils.ApiError{Message: "could not get accounts"})
		return
	}

	c.JSON(http.StatusOK, utils.ApiResponse{Message: "Success", Data: accounts})
}

func GetAccountDetails(c *gin.Context) {
	height, err := (strconv.ParseInt(c.Param("height"), 10, 64))
	if err != nil {
		log.Error("height must be a number", err)
		c.JSON(http.StatusBadRequest, utils.ApiError{Message: "height must be a number"})
		return
	}

	var address signature.PublicKey
	if err := address.UnmarshalHex(c.Param("account_address")); err != nil {
		log.Error("failed to parse account address", err)
		c.JSON(http.StatusBadRequest, utils.ApiError{Message: "failed to parse account address"})
		return
	}

	conn, err := connections.GetOasisConn()
	if err != nil {
		log.Error("error connecting to gRPC server", err)
		c.JSON(http.StatusBadRequest, utils.ApiError{Message: "error connecting to gRPC server"})
		return
	}
	defer conn.Close()

	client := api.NewStakingClient(conn)

	accountInfo, err := client.AccountInfo(c, &api.OwnerQuery{
		Height: height,
		Owner:  address,
	})
	if err != nil {
		log.Error("could not get account details", err)
		c.JSON(http.StatusBadRequest, utils.ApiError{Message: "could not get account details"})
		return
	}

	c.JSON(http.StatusOK, utils.ApiResponse{Message: "Success", Data: accountInfo})
}

func GetThreshold(c *gin.Context) {
	height, err := (strconv.ParseInt(c.Param("height"), 10, 64))
	if err != nil {
		log.Error("height must be a number", err)
		c.JSON(http.StatusBadRequest, utils.ApiError{Message: "height must be a number"})
		return
	}

	kind, err := (strconv.ParseInt(c.Param("kind"), 10, 64))
	if err != nil {
		log.Error("kind must be a number", err)
		c.JSON(http.StatusBadRequest, utils.ApiError{Message: "kind must be a number"})
		return
	}

	conn, err := connections.GetOasisConn()
	if err != nil {
		log.Error("error connecting to gRPC server", err)
		c.JSON(http.StatusBadRequest, utils.ApiError{Message: "error connecting to gRPC server"})
		return
	}
	defer conn.Close()

	client := api.NewStakingClient(conn)

	accountInfo, err := client.Threshold(c, &api.ThresholdQuery{
		Height: height,
		Kind:  api.ThresholdKind(kind),
	})
	if err != nil {
		log.Error("could not get account details", err)
		c.JSON(http.StatusBadRequest, utils.ApiError{Message: "could not get account details"})
		return
	}

	c.JSON(http.StatusOK, utils.ApiResponse{Message: "Success", Data: accountInfo})
}

func GetDebondingDelegations(c *gin.Context) {
	height, err := (strconv.ParseInt(c.Param("height"), 10, 64))
	if err != nil {
		log.Error("height must be a number", err)
		c.JSON(http.StatusBadRequest, utils.ApiError{Message: "height must be a number"})
		return
	}

	var address signature.PublicKey
	if err := address.UnmarshalHex(c.Param("account_address")); err != nil {
		log.Error("failed to parse account address", err)
		c.JSON(http.StatusBadRequest, utils.ApiError{Message: "failed to parse account address"})
		return
	}

	conn, err := connections.GetOasisConn()
	if err != nil {
		log.Error("error connecting to gRPC server", err)
		c.JSON(http.StatusBadRequest, utils.ApiError{Message: "error connecting to gRPC server"})
		return
	}
	defer conn.Close()

	client := api.NewStakingClient(conn)

	accountInfo, err := client.DebondingDelegations(c, &api.OwnerQuery{
		Height: height,
		Owner:  address,
	})
	if err != nil {
		log.Error("could not get account details", err)
		c.JSON(http.StatusBadRequest, utils.ApiError{Message: "could not get account details"})
		return
	}

	c.JSON(http.StatusOK, utils.ApiResponse{Message: "Success", Data: accountInfo})
}

func GetStateToGenesis(c *gin.Context) {
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

	client := api.NewStakingClient(conn)

	genesis, err := client.StateToGenesis(c, height)
	if err != nil {
		log.Error("could not get state to genesis", err)
		c.JSON(http.StatusBadRequest, utils.ApiError{Message: "could not get state to genesis"})
		return
	}

	c.JSON(http.StatusOK, utils.ApiResponse{Message: "Success", Data: genesis})
}