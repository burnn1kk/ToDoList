package tasklist

import (
	"time"
)

type Task struct {
	name           string
	text           string
	creationTime   time.Time
	completionMark bool
	completionTime time.Time
}

func NewTask(name string, text string) (Task, error) {
	if name == "" {
		return Task{}, ErrEmptyName
	}
	if len([]rune(name)) > 1 {
		return Task{}, ErrLongName
	}
	if text == "" {
		return Task{}, ErrEmptyText
	}

	return Task{
		name:           name,
		text:           text,
		creationTime:   time.Now(),
		completionMark: false,
		completionTime: time.Time{},
	}, nil
}

func (t *Task) GetName() string {
	return t.name
}

func (t *Task) GetText() string {
	return t.text
}
func (t *Task) GetCreatinTime() time.Time {
	return t.creationTime
}
func (t *Task) GetCompletionMark() bool {
	return t.completionMark
}

func (t *Task) GetCompletionTime() time.Time {
	return t.completionTime
}

func (t *Task) CompleteTask() {
	t.completionMark = true
	t.completionTime = time.Now()
}
