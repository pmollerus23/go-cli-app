package task

import "errors"

var (
	ErrReadTitle       = errors.New("failed to read task title")
	ErrReadDescription = errors.New("failed to read task description")
)
