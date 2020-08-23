package models

import "time"


type Transaction struct {

	ID            			ID	`json:"id" `
	Amount					float64	`json:"amount" example:"1504.56"`
	InstrumentId 			ID	`json:"instrumentId"  example:"asdfafasdfad5"`
	OriginAccountId			ID 	`json:"originAccountId" example:"asdfasdfasdf3"`
	DestinationAccountId	ID 	`json:"destinationAccountId"  example:"19fdsasfads"`
	CreatedAt     			time.Time	`json:"createdAt" example:"2020-07-07 13:36:15.738848+02:00"`
}


type TransactionBuilder interface {
	Amount(float64)		TransactionBuilder
	Instrument(ID)	TransactionBuilder
	Origin(ID) 		TransactionBuilder
	Destination(ID) 	TransactionBuilder
	Build() 			Transaction
}

type tBuilder struct {
	amount	 		float64
	instrumentId	ID
	receptorId		ID
	originId		ID
}

func (ib *tBuilder) Amount(f float64) TransactionBuilder {
	ib.amount = f
	return ib
}

func (ib *tBuilder) Instrument(id ID) TransactionBuilder {
	ib.instrumentId = id
	return ib
}

func (ib *tBuilder) Origin(origin ID) TransactionBuilder {
	ib.originId = origin
	return ib
}

func (ib *tBuilder) Destination(receptor ID) TransactionBuilder {
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
