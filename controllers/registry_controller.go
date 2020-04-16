package controllers

import (
	"fmt"
	"github.com/figment-networks/oasis-rpc-proxy/connections"
	"github.com/figment-networks/oasis-rpc-proxy/utils"
	"github.com/figment-networks/oasis-rpc-proxy/utils/log"
	"github.com/gin-gonic/gin"
	"github.com/oasislabs/oasis-core/go/common/crypto/signature"
	"github.com/oasislabs/oasis-core/go/common/entity"
	"github.com/oasislabs/oasis-core/go/common/node"
	registryApi "github.com/oasislabs/oasis-core/go/registry/api"
	"net/http"
)

type (
	GetEntityRequest struct {
		EntityId string `form:"entity_id" binding:"required"`
		Height   int64  `form:"height"`
	}
	GetEntityResponse struct {
		Message string         `json:"message"`
		Data    *entity.Entity `json:"data"`
	}
	GetNodeRequest struct {
		NodeId string `form:"node_id" binding:"required"`
		Height int64  `form:"height"`
	}
	GetNodeResponse struct {
		Message string     `json:"message"`
		Data    *node.Node `json:"data"`
	}
	GetStateDumpRequest struct {
		Height int64 `form:"height"`
	}
	GetStateDumpResponse struct {
		Message string     `json:"message"`
		Data    *registryApi.Genesis `json:"data"`
	}
)

func GetEntity(c *gin.Context) {
	var req GetEntityRequest
	if err := c.ShouldBind(&req); err != nil {
		msg := "invalid request parameters"
		log.Error(msg, err)
		c.JSON(http.StatusBadRequest, utils.ApiError{Message: msg})
		return
	}

	// Get entityId
	entityId := signature.PublicKey{}
	err := entityId.UnmarshalText([]byte(req.EntityId))
	if err != nil {
		msg := "entity_id must be a valid public key"
		log.Error(msg, err)
		c.JSON(http.StatusBadRequest, utils.ApiError{Message: msg})
		return
	}

	// Connect to node
	conn, err := connections.GetOasisConn()
	if err != nil {
		log.Error("error connecting to gRPC server", err)
		c.JSON(http.StatusBadRequest, utils.ApiError{Message: "error connecting to gRPC server"})
		return
	}
	defer conn.Close()

	fmt.Println("entityId decoded", entityId)

	// Make a request
	client := registryApi.NewRegistryClient(conn)
	n, err := client.GetEntity(c, &registryApi.IDQuery{
		Height: req.Height,
		ID:     entityId,
	})
	if err != nil {
		msg := err.Error()
		log.Error(msg, err)
		c.JSON(http.StatusUnprocessableEntity, utils.ApiError{Message: msg})
		return
	}

	c.JSON(http.StatusOK, GetEntityResponse{Message: "Success", Data: n})
}

func GetNode(c *gin.Context) {
	var req GetNodeRequest
	if err := c.ShouldBind(&req); err != nil {
		msg := "invalid request parameters"
		log.Error(msg, err)
		c.JSON(http.StatusBadRequest, utils.ApiError{Message: msg})
		return
	}
	// Get nodeId
	nodeId := signature.PublicKey{}
	err := nodeId.UnmarshalText([]byte(req.NodeId))
	if err != nil {
		msg := "entity_id must be a valid public key"
		log.Error(msg, err)
		c.JSON(http.StatusBadRequest, utils.ApiError{Message: msg})
		return
	}

	// Connect to node
	conn, err := connections.GetOasisConn()
	if err != nil {
		log.Error("error connecting to gRPC server", err)
		c.JSON(http.StatusBadRequest, utils.ApiError{Message: "error connecting to gRPC server"})
		return
	}
	defer conn.Close()

	// Make a request
	client := registryApi.NewRegistryClient(conn)
	e, err := client.GetNode(c, &registryApi.IDQuery{
		Height: req.Height,
		ID:     nodeId,
	})
	if err != nil {
		msg := err.Error()
		log.Error(msg, err)
		c.JSON(http.StatusUnprocessableEntity, utils.ApiError{Message: msg})
		return
	}

	c.JSON(http.StatusOK, GetNodeResponse{Message: "Success", Data: e})
}

func GetStateDump(c *gin.Context) {
	var req GetStateDumpRequest
	if err := c.ShouldBind(&req); err != nil {
		msg := "invalid request parameters"
		log.Error(msg, err)
		c.JSON(http.StatusBadRequest, utils.ApiError{Message: msg})
		return
	}

	// Connect to node
	conn, err := connections.GetOasisConn()
	if err != nil {
		log.Error("error connecting to gRPC server", err)
		c.JSON(http.StatusBadRequest, utils.ApiError{Message: "error connecting to gRPC server"})
		return
	}
	defer conn.Close()

	// Make a request
	client := registryApi.NewRegistryClient(conn)
	e, err := client.StateToGenesis(c, req.Height)
	if err != nil {
		msg := err.Error()
		log.Error(msg, err)
		c.JSON(http.StatusUnprocessableEntity, utils.ApiError{Message: msg})
		return
	}

	c.JSON(http.StatusOK, GetStateDumpResponse{Message: "Success", Data: e})
}