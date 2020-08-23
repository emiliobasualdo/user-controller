package models

import (
	"time"
)

type IDGP int

type GpNewAccountInput struct {
	Name           string
	Lastname       string
	DocumentNumber string
	BirthDate      string
	Cellphone      string
	ExternalId     ID
}

type GpRecharge struct {
	Amount      float64
}

type GpNewAccountOutput struct {
	ID         IDGP
	ExternalId ID
	Cards      []GPCard
}

type GPCard struct {
	CardNumber  string
	Cvc         string
}

type GPAccountMovements struct {
	Amount int
	DateFrom time.Time
	DateTo	time.Time
	TotalAmount int
	Movements []GpMovement
}

type GpMovement struct {
	ID           IDGP
	Type         int
	Date         time.Time
	Description  string
	Amount       float64
	Observations string
}

type GPAvailable struct {
	LocalAvailableBuy		float64
	LocalAvailableAdvance	float64
	DollarAvailableBuy   	float64
	DollarAvailableAdvance 	float64
	LocalBalance            float64
	DollarBalance            float64
}