package app

import "github.com/silentlight/oasis-test/controllers"

func mapUrl() {
	router.GET("/block/:height", controllers.GetBlock)
	router.GET("/transactions", controllers.GetTransactions)
}