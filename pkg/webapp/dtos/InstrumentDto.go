package dtos

import "massimple.com/wallet-controller/pkg/models"

type InstrumentDto struct {
	Holder    		string		`json:"holder" example:"José Pepe Argento"`
	LastFourNumbers string		`json:"lastFourNumbers" example:"4930"`
	ValidThru 		string		`json:"validThru" example:"11/24"`
	Issuer	 		string		`json:"issuer" example:"Banco Itaú"`
	PPS 			string		`json:"pps" example:"VISA" enums:"VISA, AMEX, MC"`
	CreditType 		string		`json:"creditType" example:"DEBIT" enums:"DEBIT, CREDIT, PREPAID"`
}

func (dto InstrumentDto) Builder() models.InstrumentBuilder{
	return models.New().
		CreditType(dto.CreditType).
		Holder(dto.Holder).
		Issuer(dto.Issuer).
		LastFourNumbers(dto.LastFourNumbers).
		PPS(dto.PPS).
		ValidThru(dto.ValidThru)
}

func (dto InstrumentDto) Build() models.Instrument{
	return dto.Builder().Build()
}
