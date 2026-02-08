package task

import "fmt"

type TaskManager struct {
	UserID   string
	TaskList []Task
}

func (tm *TaskManager) AddTask() error {
	task, err := NewTask()
	if err != nil {
		return err
	}
	fmt.Printf("Task created: %s\n", task.Name)
	tm.TaskList = append(tm.TaskList, *task)
	return nil
}

func (tm *TaskManager) ListTasks() {
	if len(tm.TaskList) > 0 {
		fmt.Println(tm.TaskList)
	}
}
