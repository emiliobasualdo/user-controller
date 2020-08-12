package models

import (
	"fmt"
)

type NoSuchAccountError struct {
	Query string
}

func (e *NoSuchAccountError) Error() string {
	return fmt.Sprintf("No such account with Query %s", e.Query)
}

type NoSuchInstrumentError struct {
	ID string
}

func (e *NoSuchInstrumentError) Error() string {
	return fmt.Sprintf("No such instrument with id %s", e.ID)
}

type UnauthorizedError struct {
	Message string
}

func (e *UnauthorizedError) Error() string {
	return fmt.Sprint(e.Message)
}

type NotEnoughCreditError struct {
}

func (e *NotEnoughCreditError) Error() string {
	return fmt.Sprint("Your account does not have the necessary credits ")
}

type DisabledAccountError struct {
	Message string
}

func (e *DisabledAccountError) Error() string {
	return fmt.Sprint(e.Message)
}

type IllegalTransactionError struct {
	Message string
}

func (e *IllegalTransactionError) Error() string {
	return fmt.Sprint(e.Message)
}
