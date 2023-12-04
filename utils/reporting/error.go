package reporting

import (
	"fmt"

	"github.com/jjangg96/oasis-rpc-proxy/config"
	"github.com/jjangg96/oasis-rpc-proxy/utils/logger"
)

func RecoverError() {
	if err := recover(); err != nil {
		logger.Info(fmt.Sprintf("recovering from error [ERR: %+v]", err))
	}
}

func Init(cfg *config.Config) {
}
