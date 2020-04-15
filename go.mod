module github.com/figment-networks/oasis-rpc-proxy

replace (
	github.com/tendermint/iavl => github.com/oasislabs/iavl v0.12.0-ekiden3
	github.com/tendermint/tendermint => github.com/oasislabs/tendermint v0.32.8-oasis1
	golang.org/x/crypto/curve25519 => github.com/oasislabs/ed25519/extra/x25519 v0.0.0-20191022155220-a426dcc8ad5f
	golang.org/x/crypto/ed25519 => github.com/oasislabs/ed25519 v0.0.0-20191109133925-b197a691e30d
)

go 1.13

require (
	github.com/btcsuite/btcd v0.20.1-beta // indirect
	github.com/dvyukov/go-fuzz v0.0.0-20191022152526-8cb203812681 // indirect
	github.com/fxamacker/cbor v1.5.1 // indirect
	github.com/gin-gonic/gin v1.5.0
	github.com/go-kit/kit v0.10.0 // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/golang/protobuf v1.3.4
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.0 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/oasislabs/ed25519 v0.0.0-20200302143042-29f6767a7c3e // indirect
	github.com/oasislabs/oasis-core/go v0.0.0-20200304114707-807935769a93
	github.com/pelletier/go-toml v1.6.0 // indirect
	github.com/prometheus/client_golang v1.5.0 // indirect
	github.com/prometheus/procfs v0.0.10 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/cobra v0.0.6 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.6.2 // indirect
	github.com/tendermint/go-amino v0.15.1 // indirect
	github.com/tendermint/tendermint v0.33.1 // indirect
	github.com/tendermint/tm-db v0.4.1 // indirect
	github.com/whyrusleeping/go-logging v0.0.1 // indirect
	go.uber.org/multierr v1.5.0 // indirect
	go.uber.org/zap v1.14.0
	golang.org/x/crypto v0.0.0-20200302210943-78000ba7a073 // indirect
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/net v0.0.0-20200301022130-244492dfa37a // indirect
	golang.org/x/sys v0.0.0-20200302150141-5c8b2ff67527 // indirect
	golang.org/x/tools v0.0.0-20200305140159-d7d444866696 // indirect
	google.golang.org/genproto v0.0.0-20200305110556-506484158171 // indirect
	google.golang.org/grpc v1.27.1
	gopkg.in/go-playground/validator.v9 v9.31.0 // indirect
	gopkg.in/ini.v1 v1.52.0 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
	honnef.co/go/tools v0.0.1-2020.1.3 // indirect
)
