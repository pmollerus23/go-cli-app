package main

import (
	"os"
	"task-manager/internal/command"
)

func main() {
	command.HandleArgs(os.Args)
}
