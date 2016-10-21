package main

import "fmt"
import "github.com/gesquive/passmakr/cmd"

var version = "v0.1.0"
var dirty = ""

func main() {
	displayVersion := fmt.Sprintf("passmakr %s%s",
		version,
		dirty)
	cmd.Execute(displayVersion)
}
