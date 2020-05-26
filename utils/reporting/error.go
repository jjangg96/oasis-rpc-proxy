package reporting

import (
	"fmt"
	"github.com/figment-networks/oasis-rpc-proxy/config"
	"github.com/figment-networks/oasis-rpc-proxy/utils/logger"
	"github.com/rollbar/rollbar-go"
)

func RecoverError() {
	if err := recover(); err!= nil {
		logger.Info(fmt.Sprintf("recovering from error [ERR: %+v]", err))
		rollbar.LogPanic(err, true)
	}
}

func Init(cfg *config.Config) {
	rollbar.SetToken(cfg.RollbarAccessToken)
	rollbar.SetEnvironment(cfg.AppEnv)
	rollbar.SetServerRoot("github.com/figment-networks/oasishub-indexer")
}
