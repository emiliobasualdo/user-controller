package models

import "time"

type Transaction struct {
	ID            			uint	`json:"id" gorm:"primary_key" example:"1"`
	Amount					float64	`json:"amount" gorm:"not null" example:"1504.56"`
	InstrumentId 			uint	`json:"instrumentId" gorm:"not null" example:"5"`
	OriginAccountId			uint 	`json:"originAccountId" gorm:"not null" example:"3"`
	DestinationAccountId	uint 	`json:"destinationAccountId" gorm:"not null" example:"19"`
	CreatedAt     			time.Time	`json:"createdAt" example:"2020-07-07 13:36:15.738848+02:00"`
}


type TransactionBuilder interface {
	Amount(float64)		TransactionBuilder
	Instrument(uint)	TransactionBuilder
	Origin(uint) 		TransactionBuilder
	Destination(uint) 	TransactionBuilder
	Build() 			Transaction
}

type tBuilder struct {
	amount	 		float64
	instrumentId	uint
	originId		uint
	receptorId		uint
}

func (ib *tBuilder) Amount(f float64) TransactionBuilder {
	ib.amount = f
	return ib
}

func (ib *tBuilder) Instrument(instId uint) TransactionBuilder {
	ib.instrumentId = instId
	return ib
}

func (ib *tBuilder) Origin(origin uint) TransactionBuilder {
	ib.originId = origin
	return ib
}

func (ib *tBuilder) Destination(receptor uint) TransactionBuilder {
	ib.receptorId = receptor
	return ib
}

func (ib *tBuilder) Build() Transaction {
	return Transaction{
		Amount:           	ib.amount,
		InstrumentId:		ib.instrumentId,
		DestinationAccountId:       	ib.receptorId,
	}
}

func NewTransactionBuilder() TransactionBuilder {
	return &tBuilder{}
}
