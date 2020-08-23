package dtos

import . "massimple.com/wallet-controller/internal/models"

type TransactionDto struct {
	Amount float64 `json:"amount" binding:"required" example:"1504.56"`
	InstrumentId ID `json:"instrumentId"  binding:"required" `
	DestinationID ID `json:"destinationAccountId" binding:"required" `
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
