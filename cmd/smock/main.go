package main

import (
	"fmt"
	"github.com/fatih/color"
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

	color.Green(fmt.Sprintf("SMock version: %v. built at %v", version, parsed.Format(time.RFC822Z)))
	commands.ExecuteRootCommand()
}
