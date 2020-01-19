package controllers

import (
	"github.com/gin-gonic/gin"
	//"github.com/oasislabs/oasis-core/go/consensus/api"
	"github.com/oasislabs/oasis-core/go/scheduler/api"
	"github.com/silentlight/oasis-test/config"
	"google.golang.org/grpc"
	"net/http"
	"strconv"
)

type Response struct {
	Message string `json:"message"`
}

func GetBlock(c *gin.Context) {
	height, err := (strconv.ParseInt(c.Param("height"), 10, 64))
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Message: "height must be a number"})
		return
	}

	conn, err := grpc.Dial(
		config.GetOasisSocket(),
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.WaitForReady(true)),
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Message: "error connecting to grpc server"})
		return
	}
	defer conn.Close()

	client := api.NewSchedulerClient(conn)

	_, err = client.GetValidators(c, height)
	if err != nil {
		c.JSON(http.StatusUnauthorized, Response{Message: "could not get block"})
		return
	}

	c.JSON(http.StatusOK, Response{Message: "All good"})
}
