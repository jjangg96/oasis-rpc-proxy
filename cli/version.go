package cli

import "fmt"

var (
	appName    = "oasis-rpc-proxy"
	appVersion = "0.6.0"
	gitCommit  = "-"
	goVersion  = "1.14"
)

func versionString() string {
	return fmt.Sprintf(
		"%s %s (git: %s, %s)",
		appName,
		appVersion,
		gitCommit,
		goVersion,
	)
}
