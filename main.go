package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"
)

type Task struct {
	id          string
	name        string
	description string
	status      TaskStatus
}

type TaskStatus int

const (
	Pending TaskStatus = iota
	InProgress
	Completed
	Cancelled
)

type TaskManager struct {
	userId   string
	taskList []Task
}

var (
	ErrReadTitle       = errors.New("failed to read task title")
	ErrReadDescription = errors.New("failed to read task description")
)

func main() {
	defer cleanUp()
	fmt.Println("Starting program...")

	if len(os.Args[1:]) == 0 {
		helpDisplay()
		return
	}

	arg_1 := os.Args[1]

	taskManager := &TaskManager{}

	switch arg_1 {
	case "add":
		err := taskManager.addTask()
		if err != nil {
			if errors.Is(err, ErrReadTitle) {
				fmt.Println("Reader failure: could not read title input")
			} else if errors.Is(err, ErrReadDescription) {
				fmt.Println("Reader failure: could not read description input")
			} else {
				fmt.Printf("Reader failure: %v\n", err)
			}
		}

	case "list":
		taskManager.ListTasks()
	}

}

func cleanUp() {
	fmt.Println("Ending program.")

}

func helpDisplay() {
	fmt.Println("To use this tool, pass an argument!")
}

func (tm *TaskManager) addTask() error {

	task, err := NewTask()
	if err != nil {
		return err
	}
	fmt.Printf("Task created: %s\n", task.name)
	// return nil

	tm.taskList = append(tm.taskList, *task)
	return nil
}

func NewTask() (*Task, error) {
	reader := bufio.NewReader(os.Stdin)
	// defer reader.Discard()
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
		id:          uuid.New().String(),
		name:        title,
		description: description,
		status:      Pending,
	}, nil
}

func (tm TaskManager) ListTasks() {
	fmt.Println(tm.taskList)
	// return tm.taskList
}
