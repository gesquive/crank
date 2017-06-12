package main

import "fmt"
import "github.com/gesquive/forge/cmd"

var version = "v0.1.2"
var dirty = ""

func main() {
	displayVersion := fmt.Sprintf("forge %s%s",
		version,
		dirty)
	cmd.Execute(displayVersion)
}
