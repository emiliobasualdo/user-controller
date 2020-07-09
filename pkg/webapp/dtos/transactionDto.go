package dtos

import . "massimple.com/wallet-controller/pkg/models"

type TransactionDto struct {
	Amount			float64 `json:"amount" example:"1504.56"`
	InstrumentId 	uint	`json:"instrumentId" example:"5"`
	DestinationID	uint 	`json:"destinationAccountId" example:"3"`
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