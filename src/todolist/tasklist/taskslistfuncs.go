package tasklist

import (
	taskerrors "tdl/todolist/errors"
	"time"
)

func (tl *TaskList) Add(name string, text string) error {
	if taskExists(*tl, name) {
		return taskerrors.ErrTaskAlreadyExists
	}
	task, err := NewTask(name, text)
	if err != nil {
		return err
	}

	(*tl)[name] = task
	return nil
}

func (tl *TaskList) Delete(name string) error {
	if !taskExists(*tl, name) {
		return taskerrors.ErrTaskNotInList
	}
	delete(*tl, name)
	return nil
}

func (tl TaskList) List() []Task {
	list := make([]Task, 0, len(tl))
	for _, v := range tl {
		list = append(list, v)
	}
	return list
}

func (tl *TaskList) CompleteTask(taskname string) error {
	if !taskExists(*tl, taskname) {
		return taskerrors.ErrTaskNotInList
	}
	task := (*tl)[taskname]
	task.completionMark = true
	task.completionTime = time.Now()
	(*tl)[taskname] = task
	return nil
}

func taskExists(tl TaskList, taskname string) bool {
	_, exists := tl[taskname]
	return exists
}
