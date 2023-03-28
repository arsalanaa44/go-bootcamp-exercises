package inmemory

import (
	"todo_cli/entity"
)

type Task struct {
	tasks []entity.Task
}

func NewTaskStore() *Task {
	return &Task{
		[]entity.Task{},
	}
}

func (t *Task) CreateNewTask(task entity.Task) (entity.Task, error) {

	task.ID = len(t.tasks) + 1
	t.tasks = append(t.tasks, task)

	return task, nil
}

func (t *Task) ListUserTasks(userID int) ([]entity.Task, error) {
	var userTasks []entity.Task
	for _, task := range t.tasks {
		if task.UserID == userID {
			userTasks = append(userTasks, task)
		}
	}
	return userTasks, nil
}
