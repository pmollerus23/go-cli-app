package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"task-manager/internal/task"
)

type Repository struct {
	UserID   string       `json:"user_id"`
	TaskList []task.Task  `json:"task_list"`
}

func (tm *Repository) AddTask() error {
	task, err := task.NewTask()
	if err != nil {
		return err
	}
	fmt.Printf("Task created: %s\n", task.Name)
	tm.TaskList = append(tm.TaskList, *task)
	return nil
}

func (tm *Repository) ListTasks() {
	if len(tm.TaskList) > 0 {
		fmt.Println(tm.TaskList)
	}
}

func ParseTaskFileFromPath(filePath string) (bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}

	defer file.Close()
	return true, nil
}

func ValidateTaskFile(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if err != nil {
		fmt.Println("error checking file info")
	}
	// if fileInfo.Mode()
	// file, err := os.Open(filePath)
	// if err != nil {
	// 	return false, err
	// }
	// defer file.Close()

	data, err := os.ReadFile(filePath)
	if err != nil {
		return false, err
	}

	if !json.Valid(data) {
		fmt.Println("Task repository is corrupt. Use 'task init' to initialize a new repository.")
		return false, errors.New(".task file contains invalid json.")
	}

	// TODO - Add parsing logic to ensure all required fields exist and struct is intact

	return true, nil

}

func InitializeTaskRepository() (*Repository, error) {
	_, err := FindRepositoryFile()
	if err == nil {
		fmt.Println(`Repository already initialized.`)
		return nil, err

	} else {
		fmt.Println("Initializing repository ...")
		f, err := os.Create(".task.json")
		if err != nil {
			fmt.Println("Error creating .task file")
			return nil, err
		}

		repository := &Repository{}

		jsonData, err := json.MarshalIndent(repository, "", "  ") // Use 2 spaces for indentation
		if err != nil {
			fmt.Printf("Error marshalling JSON: %s\n", err)
			return nil, err
		}

		if err := os.WriteFile(f.Name(), jsonData, 0644); err != nil {
			return nil, err
		}

		fmt.Println("Repository initialized")
		return repository, nil
	}
}

func FindRepositoryFile() (string, error) {
	// Start from current directory
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Walk up until we find .task or hit root
	for {
		taskPath := filepath.Join(dir, ".task.json")

		if _, err := os.Stat(taskPath); err == nil {
			return taskPath, nil
		}

		// Get parent directory
		parent := filepath.Dir(dir)

		// If we've reached root, stop
		if parent == dir {
			return "", os.ErrNotExist
		}

		dir = parent
	}
}

func FindRepositoryPath() (string, error) {
	taskFilePath, err := FindRepositoryFile()
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println(`.task.json repository not found. Use 'task init' to initialize a repository`)
			return "", os.ErrNotExist
		}
	}
	return taskFilePath, nil
}
