package main

import (
	"errors"
	"fmt"
	"os"

	"task-manager/internal/repository"
	"task-manager/internal/task"
)

func main() {
	if len(os.Args) < 2 {
		helpDisplay()
		return
	}

	cmd := os.Args[1]

	if cmd == "init" {
		repository.InitializeTaskRepository()
		return
	}

	taskFilePath, err := repository.FindRepositoryPath()
	if err != nil {
		return
	}

	if _, err := repository.ParseTaskFileFromPath(taskFilePath); err != nil {
		return
	}

	taskManager := &task.TaskManager{}

	switch cmd {
	case "add":
		err = taskManager.AddTask()
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
		taskManager.ListTasks()
	default:
		fmt.Println("Unrecognized command:", cmd)
	}
}

func helpDisplay() {
	fmt.Println("To use this tool, pass an argument!")
}
