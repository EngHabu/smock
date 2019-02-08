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
		parsed = time.Unix(0, 0)
	}

	fmt.Printf("smock version: %v. built at %v", version, parsed.Format(time.RFC822Z))
	commands.ExecuteRootCommand()
}
