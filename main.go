package main

import "fmt"
import "github.com/gesquive/passforge/cmd"

var version = "v0.1.0"
var dirty = ""

func main() {
	displayVersion := fmt.Sprintf("passforge %s%s",
		version,
		dirty)
	cmd.Execute(displayVersion)
}
