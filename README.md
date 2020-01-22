# oasis-rpc-proxy

Gin-based proxy server over Oasis gRPC.

## 1. Available endpoints

* _/block/:height_ - get block meta
* _/transactions/:height_ - get transactions at given height
* _/validators/:height_ - get validators at given height
* _/staking/:height/total_supply_ - get total supply at given height
* _/staking/:height/accounts_ - get accounts at given height
* _/staking/:height/state_to_genesis_ - get state to genesis at given height
* _/staking/:height/accounts/:account_address_ - get account details at given height for give public key
* _/staking/:height/debonding_delegators/:account_address_ - get debonding delegators for give account public key
* _/staking/:height/threshold/:kind_ - get threshold for given height and type

## 2 Environmental variables

* APP_PORT - What port to use for API
* GIN_MODE - set this to release on production
* GO_ENVIRONMENT - set go environment (production | development)
* OASIS_SOCKET - absolute path to oasis node socket