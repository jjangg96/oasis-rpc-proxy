module github.com/figment-networks/oasis-rpc-proxy

replace (
	// Updates the version used in spf13/cobra (dependency via tendermint) as
	// there is no release yet with the fix. Remove once an updated release of
	// spf13/cobra exists and tendermint is updated to include it.
	// https://github.com/spf13/cobra/issues/1091
	github.com/gorilla/websocket => github.com/gorilla/websocket v1.4.2

	github.com/tendermint/tendermint => github.com/oasislabs/tendermint v0.33.4-oasis1
	golang.org/x/crypto/curve25519 => github.com/oasislabs/ed25519/extra/x25519 v0.0.0-20191022155220-a426dcc8ad5f
	golang.org/x/crypto/ed25519 => github.com/oasislabs/ed25519 v0.0.0-20191109133925-b197a691e30d
)

go 1.13

require (
	github.com/golang/protobuf v1.4.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.0 // indirect
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/oasislabs/ed25519 v0.0.0-20200302143042-29f6767a7c3e // indirect
	github.com/oasislabs/oasis-core/go v0.0.0-20200522162332-defba01417d8
	github.com/pelletier/go-toml v1.6.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.5.1
	github.com/prometheus/procfs v0.0.10 // indirect
	github.com/rollbar/rollbar-go v1.2.0
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/tendermint/go-amino v0.15.1 // indirect
	github.com/tendermint/tendermint v0.33.4
	go.uber.org/zap v1.15.0
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/tools v0.0.0-20200305140159-d7d444866696 // indirect
	google.golang.org/genproto v0.0.0-20200305110556-506484158171 // indirect
	google.golang.org/grpc v1.29.1
	gopkg.in/ini.v1 v1.52.0 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
	honnef.co/go/tools v0.0.1-2020.1.3 // indirect
)
