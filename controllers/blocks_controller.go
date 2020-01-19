package controllers

import (
	"fmt"
	"github.com/silentlight/oasis-test/config"
	"github.com/gin-gonic/gin"
	"github.com/oasislabs/oasis-core/go/consensus/api"
	"google.golang.org/grpc"
	"net/http"
)

func GetBlock(c *gin.Context) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(config.GetOasisSocket(), opts...)
	if err != nil {
		fmt.Println("error connecting to grpc", err.Error())
		return
	}
	defer conn.Close()

	client := api.NewConsensusClient(conn)

	height := int64(0)

	block, err := client.GetBlock(c, height)
	if err != nil {
		fmt.Println("could not get block", err.Error())
		return
	}

	c.JSON(http.StatusOK, block)
}
