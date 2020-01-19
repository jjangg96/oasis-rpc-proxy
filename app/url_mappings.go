package app

import "github.com/silentlight/oasis-test/controllers"

func mapUrl() {
	router.GET("/block/:block_id", controllers.GetBlock)
	router.GET("/transactions", controllers.GetTransactions)
}