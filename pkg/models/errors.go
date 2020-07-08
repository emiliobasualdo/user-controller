package models

import (
	"fmt"
)

type NoSuchAccountError struct {
	ID uint
}

func (e *NoSuchAccountError) Error() string {
	return fmt.Sprintf("No such account with id %d", e.ID)
}

type NoSuchInstrumentError struct {
	ID uint
}

func (e *NoSuchInstrumentError) Error() string {
	return fmt.Sprintf("No such instrument with id %d", e.ID)
}

type UnauthorizedError struct {
	Message string
}

func (e *UnauthorizedError) Error() string {
	return fmt.Sprint(e.Message)
}
