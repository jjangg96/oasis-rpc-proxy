# oasis-rpc-proxy

### Description
gRPC proxy over Oasis gRPC. Currently the only way to connect to Oasis node is using unix socket.
This proxy is responsible for exposing Oasis node's data to the outside work using gRPC. 
It also serves as a anti-corruption layer which is responsible for converting raw Oasis data to 
the format understood by the oasishub-indexer 

### Oasis Core version
```shell script
20.6
```
### Available gRPC Services

* `AccountService@GetByPublickKey` - get account by public key
* `ChainService@GetCurrent` - get current chain information
* `BlockService@GetByHeight` - get block data by height
* `StateService@GetByHeight` - get state data by height
* `ValidatorService@GetByHeight` - get validators data by height
* `TransactionsService@GetByHeight` - get transactions data by height

### Environmental variables

* `APP_PORT` - What port to use for API
* `GO_ENVIRONMENT` - set go environment (production | development)
* `OASIS_SOCKET` - absolute path to oasis node socket
* `GENESIS_FILE_PATH` - absolute path to genesis.json (It is needed for decrypting transactions)

### Running the project

To run the project use:

```shell script
go run main.go -config path/to/config.json -cmd=grpc
```