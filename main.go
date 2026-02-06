package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

type Task struct {
	id          string
	name        string
	description string
	category    string
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
	// ErrTaskFileNotExist = errors.New(".task error does not exist in working directory tree")
)

func parseTaskFileFromPath(filePath string) (bool, error) {
	_, err := os.Open(filePath) // For read access.
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	return true, nil
}

func initializeTaskRepository() {

	_, err := findTaskFile()
	if err == nil {
		fmt.Println(`Repository already initialized.`)
		return

	} else {
		fmt.Println("Initializing repository ...")
	}

}

func main() {

	var arg_1 string

	if len(os.Args[1:]) > 0 {
		arg_1 = os.Args[1]
		if arg_1 == "init" {
			initializeTaskRepository()
			return
		}

	} else {
		_, err := findRepositoryPath()
		if err != nil {
			return
		}
		if len(os.Args[1:]) == 0 {
			helpDisplay()
			return
		} else {
			arg_1 = os.Args[1]
		}
	}

	taskFilePath, err := findRepositoryPath()
	if err != nil {
		return
	}

	_, err = parseTaskFileFromPath(taskFilePath)

	// fmt.Println(_)

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
	default:
		fmt.Println("Unrecognized command:", arg_1)
	}

}

func cleanUp() {
	fmt.Println("Exiting")
	os.Exit(0)

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
		category:    "",
		status:      Pending,
	}, nil
}

func (tm TaskManager) ListTasks() {
	if len(tm.taskList) > 0 {
		fmt.Println(tm.taskList)
	}
}

func findTaskFile() (string, error) {
	// Start from current directory
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Walk up until we find .task or hit root
	for {
		taskPath := filepath.Join(dir, ".task")

		if _, err := os.Stat(taskPath); err == nil {
			// fmt.Println(".task file found at path:", taskPath)
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

func findRepositoryPath() (string, error) {
	taskFilePath, err := findTaskFile()
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println(`.task repository not found. Use 'task init' to initialize a repository`)
			return "", os.ErrNotExist
		}
	}
	return taskFilePath, nil
}
