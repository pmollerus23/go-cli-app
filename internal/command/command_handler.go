package command

import (
	"fmt"
	"task-manager/internal/repository"
)

func HandleArgs(args []string) error {
	if len(args) < 2 {
		helpDisplay()
		return nil
	}

	cmd := args[1]

	if cmd == "init" {
		if err := repository.InitializeTaskRepository(); err != nil {
			return err
		}
		fmt.Println("Repository initialized")
		return nil
	}

	repo, err := repository.LoadRepositoryFromFile()
	if err != nil {
		return err
	}

	switch cmd {
	case "add":
		t, err := repo.AddTask()
		if err != nil {
			return fmt.Errorf("adding task: %w", err)
		}
		fmt.Printf("Task created: %s\n", t.Name)
	case "list":
		for _, t := range repo.Tasks {
			fmt.Println(t)
		}
	default:
		return fmt.Errorf("unrecognized command: %s", cmd)
	}

	return nil
}

func helpDisplay() {
	fmt.Println("To use this tool, pass an argument!")
}
