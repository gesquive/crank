package main

import (
	"fmt"

	"github.com/gesquive/crank/cmd"
)

var version = "v0.1.2"
var dirty = ""

func main() {
	displayVersion := fmt.Sprintf("crank %s%s",
		version,
		dirty)
	cmd.Execute(displayVersion)
}
