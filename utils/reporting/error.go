package reporting

import (
	"fmt"

	"github.com/figment-networks/oasis-rpc-proxy/config"
	"github.com/figment-networks/oasis-rpc-proxy/utils/logger"
)

func RecoverError() {
	if err := recover(); err != nil {
		logger.Info(fmt.Sprintf("recovering from error [ERR: %+v]", err))
	}
}

func Init(cfg *config.Config) {
}
