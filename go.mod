module github.com/figment-networks/oasis-rpc-proxy

go 1.14

replace (
	// Updates the version used in spf13/cobra (dependency via tendermint) as
	// there is no release yet with the fix. Remove once an updated release of
	// spf13/cobra exists and tendermint is updated to include it.
	// https://github.com/spf13/cobra/issues/1091
	github.com/gorilla/websocket => github.com/gorilla/websocket v1.4.2

	github.com/tendermint/tendermint => github.com/oasisprotocol/tendermint v0.34.0-rc4-oasis2
	golang.org/x/crypto/curve25519 => github.com/oasisprotocol/ed25519/extra/x25519 v0.0.0-20200819094954-65138ca6ec7c
	golang.org/x/crypto/ed25519 => github.com/oasisprotocol/ed25519 v0.0.0-20200819094954-65138ca6ec7c
)

require (
	github.com/golang/protobuf v1.4.2
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/oasisprotocol/oasis-core/go v0.2011.0
	github.com/pelletier/go-toml v1.6.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.7.1
	github.com/rollbar/rollbar-go v1.2.0
	github.com/smartystreets/assertions v1.0.0 // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/tendermint/tendermint v0.33.6
	go.uber.org/zap v1.15.0
	golang.org/x/crypto v0.0.0-20200709230013-948cd5f35899 // indirect
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/tools v0.0.0-20200305140159-d7d444866696 // indirect
	google.golang.org/grpc v1.32.0
	google.golang.org/protobuf v1.25.0
	gopkg.in/ini.v1 v1.52.0 // indirect
	honnef.co/go/tools v0.0.1-2020.1.3 // indirect
)
