package config

import "os"

const  (
	oasisSocketVariable = "OASIS_SOCKET"
)

var (
	oasisSocket = os.Getenv(oasisSocketVariable)
)

func GetOasisSocket() string {
	return oasisSocket
}
