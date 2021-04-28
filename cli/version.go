package cli

import "fmt"

var (
	appName    = "oasis-rpc-proxy"
	appVersion = "0.8.1"
	gitCommit  = "-"
	goVersion  = "1.16"
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
