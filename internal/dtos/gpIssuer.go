package dtos

import (
	"massimple.com/wallet-controller/internal/models"
	"time"
)
type IDGP int

type GpNewAccountInputDto struct {
	Name           string
	Lastname       string
	DocumentNumber string
	BirthDate      string
	Cellphone      string
	ExternalId     models.ID
}

type GpRechargeDto struct {
	Amount      float64
}

type GpNewAccountOutputDto struct {
	ID         IDGP
	ExternalId models.ID
	Cards      []GPCardDto
}

type GPCardDto struct {
	CardNumber  string
	Cvc         string
}

type GPAccountMovementsDto struct {
	Amount int
	DateFrom time.Time
	DateTo	time.Time
	TotalAmount int
	Movements []GpMovementDto
}

type GpMovementDto struct {
	ID 				IDGP
	Type 			int
	Date 			time.Time
	Description	 	string
	Amount 			float64
	Observations 	string
}

type GPAvailableDto struct {
	LocalAvailableBuy		float64
	LocalAvailableAdvance	float64
	DollarAvailableBuy   	float64
	DollarAvailableAdvance 	float64
	LocalBalance            float64
	DollarBalance            float64
}