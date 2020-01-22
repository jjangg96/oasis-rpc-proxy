package app

import "github.com/figment-networks/oasis-rpc-proxy/controllers"

func mapUrl() {
	router.GET("/block/:height", controllers.GetBlock)
	router.GET("/transactions/:height", controllers.GetTransactions)
	router.GET("/validators/:height", controllers.GetValidators)
	router.GET("/staking/:height/total_supply", controllers.GetTotalSupply)
	router.GET("/staking/:height/common_pool", controllers.GetCommonPool)
	router.GET("/staking/:height/accounts", controllers.GetAccounts)
	router.GET("/staking/:height/state_to_genesis", controllers.GetStateToGenesis)
	router.GET("/staking/:height/accounts/:account_address", controllers.GetAccountDetails)
	router.GET("/staking/:height/debonding_delegations/:account_address", controllers.GetDebondingDelegations)
	router.GET("/staking/:height/threshold/:kind", controllers.GetThreshold)
}