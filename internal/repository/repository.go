package repository

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func ParseTaskFileFromPath(filePath string) (bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()
	return true, nil
}

func InitializeTaskRepository() {
	_, err := FindTaskFile()
	if err == nil {
		fmt.Println(`Repository already initialized.`)
		return

	} else {
		fmt.Println("Initializing repository ...")
		_, err := os.Create(".task")
		if err != nil {
			fmt.Println("Error creating .task file")
			return
		}
		fmt.Println("Repository initialized")
		return
	}
}

func FindTaskFile() (string, error) {
	// Start from current directory
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Walk up until we find .task or hit root
	for {
		taskPath := filepath.Join(dir, ".task")

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
	taskFilePath, err := FindTaskFile()
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println(`.task repository not found. Use 'task init' to initialize a repository`)
			return "", os.ErrNotExist
		}
	}
	return taskFilePath, nil
}
