package main

import (
	"github.com/gesquive/crank/cmd"
)

var (
	buildVersion = "v1.0.0-dev"
	buildCommit  = ""
	buildDate    = ""
)

func main() {
	cmd.BuildVersion = buildVersion
	cmd.BuildCommit = buildCommit
	cmd.BuildDate = buildDate
	cmd.Execute()
}
