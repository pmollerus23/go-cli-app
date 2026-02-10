package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"task-manager/internal/task"
)

type Repository struct {
	UserID string      `json:"user_id"`
	Tasks  []task.Task `json:"task_list"`
}

func LoadRepositoryFromFile() (*Repository, error) {
	path, err := FindRepositoryFile()
	if err != nil {
		return nil, fmt.Errorf(".task directory not found. Run 'task init' to initialize task folder")
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading repository: %w", err)
	}

	repo := &Repository{}
	if err := json.Unmarshal(data, repo); err != nil {
		return nil, fmt.Errorf("parsing repository: %w", err)
	}

	return repo, nil
}

func (r *Repository) AddTask() (*task.Task, error) {
	t, err := task.NewTask()
	if err != nil {
		return nil, err
	}
	r.Tasks = append(r.Tasks, *t)
	data, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("marshalling repository: %w", err)
	}

	if err := os.WriteFile(".task/repo.json", data, 0644); err != nil {
		return nil, fmt.Errorf("writing repository file: %w", err)
	}
	return t, nil
}

func InitializeTaskRepository() error {
	if _, err := FindRepositoryFile(); err == nil {
		return fmt.Errorf("repository already initialized")
	}

	repo := &Repository{}
	data, err := json.MarshalIndent(repo, "", "  ")
	if err != nil {
		return fmt.Errorf("marshalling repository: %w", err)
	}

	if err := os.Mkdir(".task", 0755); err != nil {
		return fmt.Errorf("creating task directory: %w", err)
	}

	if err := os.WriteFile(".task/repo.json", data, 0644); err != nil {
		return fmt.Errorf("writing repo file: %w", err)
	}

	return nil
}

func FindRepositoryFile() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	repoDir := ".task/"
	repoFile := "repo.json"

	for {
		taskPath := filepath.Join(dir, repoDir, repoFile)
		if _, err := os.Stat(taskPath); err == nil {
			return taskPath, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return "", os.ErrNotExist
		}
		dir = parent
	}
}
