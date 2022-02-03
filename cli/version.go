package cli

import "fmt"

var (
	appName    = "oasis-rpc-proxy"
	appVersion = "0.9.6"
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
