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
		if len(repo.Tasks) == 0 {
			fmt.Println("no active tasks")
			return nil
		}
		for i, t := range repo.Tasks {
			fmt.Printf("[%d] %s\n    Description: %s\n    Status:      %s\n", i+1, t.Name, t.Description, t.Status)
		}
	default:
		return fmt.Errorf("unrecognized command: %s", cmd)
	}

	return nil
}

func helpDisplay() {
	fmt.Println("To use this tool, pass an argument!")
}
