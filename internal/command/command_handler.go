package command

import (
	"fmt"
	"strconv"
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
	case "delete":
		if len(repo.Tasks) == 0 {
			fmt.Println("no active tasks")
			return nil
		}
		if len(args) < 3 {
			return fmt.Errorf("usage: task delete <number>")
		}
		num, err := strconv.Atoi(args[2])
		if err != nil {
			return fmt.Errorf("invalid task number: %s", args[2])
		}
		t, err := repo.DeleteTaskByNumber(num)
		if err != nil {
			return fmt.Errorf("deleting task: %w", err)
		}
		fmt.Printf("Deleted task: %s\n", t.Name)
	default:
		return fmt.Errorf("unrecognized command: %s", cmd)
	}

	return nil
}

func helpDisplay() {
	fmt.Println(`Usage: task <command>

Commands:
  init        Initialize a new task repository
  add         Create a new task
  list        List all tasks`)
}
