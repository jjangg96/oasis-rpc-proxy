package config

import "os"

const (
	oasisSocketVariable     = "OASIS_SOCKET"
	genesisFilePathVariable = "GENESIS_FILE_PATH"
	goEnvironmentVariable   = "GO_ENVIRONMENT"
	logLevelVariable        = "LOG_LEVEL"
	appPortVariable         = "PORT"

	production = "production"
)

var (
	oasisSocket     = os.Getenv(oasisSocketVariable)
	appPort         = "8080"
	genesisFilePath = "genesis.json"
	logLevel        = "info"
)

func OasisSocket() string {
	return oasisSocket
}

func AppPort() string {
	if port := os.Getenv(appPortVariable); port != "" {
		return port
	}
	return appPort
}

func GenesisFilePath() string {
	if filePath := os.Getenv(genesisFilePathVariable); filePath != "" {
		return filePath
	}
	return genesisFilePath
}

func LogLevel() string {
	if level := os.Getenv(logLevelVariable); level != "" {
		return level
	}
	return logLevel
}

func IsProduction() bool {
	return os.Getenv(goEnvironmentVariable) == production
}
