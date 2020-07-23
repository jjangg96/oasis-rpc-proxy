package cli

import "fmt"

var (
	appName    = "oasis-rpc-proxy"
	appVersion = "0.3.11"
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
