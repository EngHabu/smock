package main

import (
	"fmt"
	"smock/cmd/smock/commands"
	"time"
)

var (
	version = "dev"
	date    = "unknown"
)

func main() {
	parsed, err := time.Parse(time.RFC3339, date)
	if err != nil {
		panic(err)
	}
	
	fmt.Printf("smock version: %v. built at %v", version, parsed.Format(time.RFC822Z))
	commands.ExecuteRootCommand()
}
