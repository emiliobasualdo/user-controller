package dtos

import . "massimple.com/wallet-controller/internal/models"
type AccountDto struct {
	Name        string        `json:"name" binding:"required" example:"Mónica"`
	LastName    string        `json:"lastName" binding:"required" example:"Potrelli de Argento"`
	DNI			string        `json:"dni" binding:"required" example:"21399433"`
	Email		string        `json:"email" binding:"required" example:"moni.argento@hotmail.com"`
	Instruments []InstrumentDto `json:"instruments"`
	Prepaids	[]PrepaidDto	`json:"prepaids"`
}

func (td AccountDto) Builder() AccountBuilder {
	return NewAccountBuilder().
		Name(td.Name).
		LastName(td.LastName).
		DNI(td.DNI).
		Email(td.Email)
}

func (td AccountDto) Build() Account {
	return td.Builder().Build()
}

func AccountDtoFromAccount(acc Account) AccountDto {
	instruments := mapInstruments(acc.Instruments)
	prepaids := mapPrepaids(acc.Prepaids)
	return AccountDto{
		Name:     acc.Name,
		LastName: acc.LastName,
		DNI:      acc.DNI,
		Email:    acc.Email,
		Instruments: instruments,
		Prepaids: prepaids,
	}
}

func FillAccountFromDto(acc Account, dto AccountDto) Account {
	acc.Name = 		dto.Name
	acc.LastName = 	dto.LastName
	acc.DNI = 		dto.DNI
	acc.Email = 	dto.Email
	return acc
}

func mapInstruments(instruments []Instrument) []InstrumentDto {
	resp := make([]InstrumentDto, 0)
	for _, inst := range instruments {
		resp = append(resp, InstDtoFromDto(inst))
	}
	return resp
}

func mapPrepaids(prepaids []Prepaid) []PrepaidDto {
	resp := make([]PrepaidDto, 0)
	for _, inst := range prepaids {
		resp = append(resp, PrepaidDtoFromPrepaid(inst))
	}
	return resp
}

func PrepaidDtoFromPrepaid(prepaid Prepaid) PrepaidDto {
	return PrepaidDto{
		ID:                      prepaid.ID,
		Brand:                   prepaid.Brand,
		DisponibleCompra:        0,
		DisponibleAnticipo:      0,
		DisponibleCompraDolar:   0,
		DisponibleAnticipoDolar: 0,
		SaldoPesos:              0,
		SaldoDolar:              0,
	}
}
