package models

import "time"

type Transaction struct {
	ID            			string	`json:"id" gorm:"primary_key" example:"1asdfasdf"`
	Amount					float64	`json:"amount" gorm:"not null" example:"1504.56"`
	InstrumentId 			string	`json:"instrumentId" gorm:"not null" example:"asdfafasdfad5"`
	OriginAccountId			string 	`json:"originAccountId" gorm:"not null" example:"asdfasdfasdf3"`
	DestinationAccountId	string 	`json:"destinationAccountId" gorm:"not null" example:"19fdsasfads"`
	CreatedAt     			time.Time	`json:"createdAt" example:"2020-07-07 13:36:15.738848+02:00"`
}


type TransactionBuilder interface {
	Amount(float64)		TransactionBuilder
	Instrument(string)	TransactionBuilder
	Origin(string) 		TransactionBuilder
	Destination(string) 	TransactionBuilder
	Build() 			Transaction
}

type tBuilder struct {
	amount	 		float64
	instrumentId	string
	originId		string
	receptorId		string
}

func (ib *tBuilder) Amount(f float64) TransactionBuilder {
	ib.amount = f
	return ib
}

func (ib *tBuilder) Instrument(instId string) TransactionBuilder {
	ib.instrumentId = instId
	return ib
}

func (ib *tBuilder) Origin(origin string) TransactionBuilder {
	ib.originId = origin
	return ib
}

func (ib *tBuilder) Destination(receptor string) TransactionBuilder {
	ib.receptorId = receptor
	return ib
}

func (ib *tBuilder) Build() Transaction {
	return Transaction{
		Amount:           	ib.amount,
		InstrumentId:		ib.instrumentId,
		DestinationAccountId:    ib.receptorId,
	}
}

func NewTransactionBuilder() TransactionBuilder {
	return &tBuilder{}
}
