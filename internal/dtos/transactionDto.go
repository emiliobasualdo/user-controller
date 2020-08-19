package dtos

import . "massimple.com/wallet-controller/internal/models"

type TransactionDto struct {
	Amount			float64 `json:"amount" binding:"required" example:"1504.56"`
	InstrumentId 	ID	`json:"instrumentId"  binding:"required"example:"5"`
	DestinationID	ID 	`json:"destinationAccountId" binding:"required" example:"3"`
}

func (td TransactionDto) Builder() TransactionBuilder {
	return NewTransactionBuilder().
		Amount(td.Amount).
		Instrument(td.InstrumentId).
		Destination(td.DestinationID)
}

func (td TransactionDto) Build() Transaction {
	return td.Builder().Build()
}