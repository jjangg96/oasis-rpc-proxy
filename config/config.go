package config

import "os"

const (
	oasisSocketVariable = "OASIS_SOCKET"
	LogLevel            = "info"
	goEnvironment       = "GO_ENVIRONMENT"
	production          = "production"
	appPortVariable     = "PORT"
)

var (
	oasisSocket = os.Getenv(oasisSocketVariable)
	appPort     = "8080"
)

func GetOasisSocket() string {
	return oasisSocket
}

func GetAppPort () string {
	if port := os.Getenv(appPortVariable); port != "" {
		return port
	}

	return appPort
}

func IsProduction() bool {
	return os.Getenv(goEnvironment) == production
}
