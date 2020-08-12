package dtos

import "massimple.com/wallet-controller/internal/models"

type InstrumentDto struct {
	Holder    		string		`json:"holder" binding:"required" example:"José Pepe Argento"`
	LastFourNumbers string		`json:"lastFourNumbers" binding:"required" example:"4930"`
	ValidThru 		string		`json:"validThru"binding:"required"  example:"11/24"`
	Issuer	 		string		`json:"issuer" binding:"required" example:"Banco Itaú"`
	PPS 			string		`json:"pps" binding:"required" example:"VISA" enums:"VISA, AMEX, MC"`
	CreditType 		string		`json:"creditType" binding:"required" example:"DEBIT" enums:"DEBIT, CREDIT, PREPAID"`
}

func (dto InstrumentDto) Builder() models.InstrumentBuilder{
	return models.NewInstrumentBuilder().
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

func InstDtoFromDto(inst models.Instrument) InstrumentDto {
	return InstrumentDto{
		Holder:          inst.Holder,
		LastFourNumbers: inst.LastFourNumbers,
		ValidThru:       inst.ValidThru,
		Issuer:          inst.Issuer,
		PPS:             inst.PPS,
		CreditType:      inst.CreditType,
	}
}