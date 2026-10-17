package tasklist

import (
	"time"

	"github.com/k0kubun/pp"
)

type Task struct {
	name           string
	text           string
	creationTime   time.Time
	completionMark bool
	completionTime time.Time
}

type TaskList map[string]Task

func NewTask(name string, text string) (Task, error) {
	if err := validateTaskName(name); err != nil {
		return Task{}, err
	}
	if err := validateTaskText(text); err != nil {
		return Task{}, err
	}

	return Task{
		name:           name,
		text:           text,
		creationTime:   time.Now(),
		completionMark: false,
		completionTime: time.Time{},
	}, nil
}

func (t Task) GetName() string {
	return t.name
}

func (t Task) GetText() string {
	return t.text
}
func (t Task) GetCreatinTime() time.Time {
	return t.creationTime
}
func (t Task) GetCompletionMark() bool {
	return t.completionMark
}

func (t Task) GetCompletionTime() time.Time {
	return t.completionTime
}

func PrintList(list []Task) {
	for _, t := range list {
		pp.Println(struct {
			Name           string
			Text           string
			CreationTime   time.Time
			CompletionMark bool
			CompletionTime time.Time
		}{
			Name:           t.name,
			Text:           t.text,
			CreationTime:   t.creationTime,
			CompletionMark: t.completionMark,
			CompletionTime: t.completionTime,
		})
	}
}
