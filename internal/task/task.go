package task

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"
)

type Task struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Category    string     `json:"category"`
	Status      TaskStatus `json:"status"`
}

func NewTask() (*Task, error) {
	reader := bufio.NewReader(os.Stdin)
	// Read title
	fmt.Print("Enter task title: ")
	title, err := reader.ReadString('\n')
	if err != nil {
		return nil, ErrReadTitle
	}
	title = strings.TrimSpace(title)

	// Read description
	fmt.Print("Enter task description: ")
	description, err := reader.ReadString('\n')
	if err != nil {
		return nil, ErrReadDescription
	}
	description = strings.TrimSpace(description)

	return &Task{
		ID:          uuid.New().String(),
		Name:        title,
		Description: description,
		Category:    "",
		Status:      Pending,
	}, nil
}
