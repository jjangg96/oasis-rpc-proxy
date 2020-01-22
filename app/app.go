package app

import (
	"fmt"
	"github.com/figment-networks/oasis-rpc-proxy/config"
	"github.com/figment-networks/oasis-rpc-proxy/log"
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func StartApp() {
	mapUrl()

	port := config.GetAppPort()
	log.Info(fmt.Sprintf("Starting server at port %s...", port))
	if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
		panic(err)
	}
}
