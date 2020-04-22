# oasis-rpc-proxy

### Description
gRPC proxy over Oasis gRPC. Currently the only way to connect to Oasis node is using unix socket.
This proxy is responsible for exposing Oasis node's data to the outside work using gRPC. 
It also serves as a anti-corruption layer which is responsible for converting raw Oasis data to 
the format understood by oasishub-indexer 

### Available gRPC Services

* _ChainService@GetCurrent_ - get current chain information
* _BlockService@GetByHeight_ - get block data by height
* _StateService@GetByHeight_ - get state data by height
* _ValidatorService@GetByHeight_ - get validators data by height
* _TransactionsService@GetByHeight_ - get transactions data by height

### Environmental variables

* APP_PORT - What port to use for API
* GO_ENVIRONMENT - set go environment (production | development)
* OASIS_SOCKET - absolute path to oasis node socket
* GENESIS_FILE_PATH - absolute path to genesis.json (It is needed for decrypting transactions)

### Running the project

To run the project use:

```shell script
go run main.go
```