package app

import "github.com/figment-networks/oasis-rpc-proxy/controllers"

func mapUrl() {
	router.GET("/block/:block_id", controllers.GetBlock)
	router.GET("/transactions", controllers.GetTransactions)
}