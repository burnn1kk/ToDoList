package taskslisterrors

import (
	"errors"
)

var (
	ErrEmptyName         = errors.New("task name cannot be empty")
	ErrLongName          = errors.New("task name cannot be more than one word")
	ErrEmptyText         = errors.New("task text cannot be empty")
	ErrTaskAlreadyExists = errors.New("task with such name already exists")
	ErrTaskNotInList     = errors.New("task with such name does not exist")
)
