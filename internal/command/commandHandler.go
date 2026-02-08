package command

import (
	"errors"
	"fmt"
	"task-manager/internal/repository"
	"task-manager/internal/task"
)

func HandleArgs(args []string) {
	if len(args) < 2 {
		helpDisplay()
		return
	}

	cmd := args[1]

	if cmd == "init" {
		err := repository.InitializeTaskRepository()
		if err != nil {
			return
		}
		return
	}

	repo, err := repository.LoadRepositoryFromFile()
	if err != nil {
		return
	}

	switch cmd {
	case "add":
		err = repo.AddTask()
		if err != nil {
			if errors.Is(err, task.ErrReadTitle) {
				fmt.Println("Reader failure: could not read title input")
			} else if errors.Is(err, task.ErrReadDescription) {
				fmt.Println("Reader failure: could not read description input")
			} else {
				fmt.Printf("Reader failure: %v\n", err)
			}
		}
	case "list":
		repo.ListTasks()
	default:
		fmt.Println("Unrecognized command:", cmd)
	}

}

func helpDisplay() {
	fmt.Println("To use this tool, pass an argument!")
}
