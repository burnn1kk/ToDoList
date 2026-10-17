package tasklist

import (
	taskerrors "tdl/todolist/errors"
)

func validateTaskName(name string) error {
	if name == "" {
		return taskerrors.ErrEmptyName
	}
	return nil
}

func validateTaskText(text string) error {
	if text == "" {
		return taskerrors.ErrEmptyText
	}
	return nil
}
