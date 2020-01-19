package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/oasislabs/oasis-core/go/consensus/api"
	"github.com/silentlight/oasis-test/config"
	"google.golang.org/grpc"
	"net/http"
	"strconv"
)

func GetBlock(c *gin.Context) {
	height, err := (strconv.ParseInt(c.Param("height"), 10, 64))
	if err != nil {
		c.JSON(http.StatusBadRequest, "height must be a number")
		return
	}

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(config.GetOasisSocket(), opts...)
	if err != nil {
		c.JSON(http.StatusBadRequest, "error connecting to grpc server")
		return
	}
	defer conn.Close()

	client := api.NewConsensusClient(conn)

	block, err := client.GetBlock(c, height)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "could not get block")
		return
	}

	c.JSON(http.StatusOK, block)
}
