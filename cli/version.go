package cli

import "fmt"

var (
	appName    = "oasis-rpc-proxy"
	appVersion = "0.3.8"
	gitCommit  = "-"
	goVersion  = "-"
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
