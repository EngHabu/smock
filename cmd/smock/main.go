package main

import (
	"fmt"
	"smock/cmd/smock/commands"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	fmt.Printf("%v, commit %v, built at %v", version, commit, date)
	commands.ExecuteRootCommand()
}
