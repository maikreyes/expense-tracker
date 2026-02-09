package main

import (
	"expense-tracker/cmd/cli"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: expense-tracker <command> [flags]")
		return
	}

	cmd := os.Args[1]
	cli.Run(cmd)
}
