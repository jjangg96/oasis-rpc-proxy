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

* `APP_ENV` - application environment (development | production) 
* `SERVER_ADDR` - address to use for API
* `SERVER_PORT` - port to use for API
* `OASIS_SOCKET` - absolute path to oasis node socket
* `GENESIS_FILE_PATH` - absolute path to genesis.json (It is needed for decrypting transactions)
* `LOG_LEVEL` - level of log
* `LOG_OUTPUT` - log output (ie. stdout or /tmp/logs.json)
* `ROLLBAR_ACCESS_TOKEN` - Rollbar access token for error reporting
* `ROLLBAR_SERVER_ROOT` - Rollbar server root for error reporting


### Running the project

To run the project use:

```shell script
go run main.go -config path/to/config.json -cmd=server
```