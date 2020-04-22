package config

import "os"

const (
	// Path to Oasis socket
	oasisSocketVariable     = "OASIS_SOCKET"
	// Path to genesis.json file
	genesisFilePathVariable = "GENESIS_FILE_PATH"
	goEnvironmentVariable   = "GO_ENVIRONMENT"
	logLevelVariable        = "LOG_LEVEL"
	// Port of the gRPC server
	appPortVariable         = "PORT"

	production = "production"
)

var (
	oasisSocket     = "internal.sock"
	appPort         = "50051"
	genesisFilePath = "genesis.json"
	logLevel        = "info"
)

func OasisSocket() string {
	if socketPath := os.Getenv(oasisSocketVariable); socketPath != "" {
		return socketPath
	}
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
