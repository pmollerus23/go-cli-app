package main

import (
	"fmt"
	"os"
	"task-manager/internal/command"
)

func main() {
	if err := command.HandleArgs(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
