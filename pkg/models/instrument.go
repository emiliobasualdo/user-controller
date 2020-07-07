package models

import "time"

type Instrument struct {
	ID            	uint		`json:"id" gorm:"primary_key"`
	AccountID		uint		`json:"accountId"`
	Holder    		string		`json:"holder"`
	LastFourNumbers string		`json:"lastFourNumbers"`
	ValidThru 		string		`json:"validThru"`
	Issuer	 		string		`json:"issuer"` // banco itaú
	PPS 			string		`json:"pps"`// visa o amex o mc
	CreditType 		string		`json:"creditType"`
	CreatedAt     	time.Time	`json:"createdAt"`
	DisabledAt		time.Time	`json:"-"`// we dont export this field
}

type InstrumentDto struct {
	AccountID		uint		`json:"accountId"`
	Holder    		string		`json:"holder"`
	LastFourNumbers string		`json:"lastFourNumbers"`
	ValidThru 		string		`json:"validThru"`
	Issuer	 		string		`json:"issuer"`
	PPS 			string		`json:"pps"`
	CreditType 		string		`json:"creditType"`
}

func (dto InstrumentDto) Builder() InstrumentBuilder{
	return New().
		CreditType(dto.CreditType).
		Holder(dto.Holder).
		Issuer(dto.Issuer).
		LastFourNumbers(dto.LastFourNumbers).
		PPS(dto.PPS).
		ValidThru(dto.ValidThru)
}

type InstrumentBuilder interface {
	FromAccount(uint)	InstrumentBuilder
	Holder(string)    		InstrumentBuilder
	LastFourNumbers(string) InstrumentBuilder
	ValidThru(string) 		InstrumentBuilder
	Issuer(string)	 		InstrumentBuilder
	PPS(string) 			InstrumentBuilder
	CreditType(string) 		InstrumentBuilder
	Build() 				Instrument
}

type auxBuilder struct {
	account			uint
	holder    		string
	lastFourNumbers string
	validThru 		string
	issuer	 		string
	pps				string
	creditType 		string
}

func (ib *auxBuilder) FromAccount(acc uint) InstrumentBuilder {
	ib.account = acc
	return ib
}

func (ib *auxBuilder) Holder(s string) InstrumentBuilder {
	ib.holder = s
	return ib
}

func (ib *auxBuilder) LastFourNumbers(s string) InstrumentBuilder {
	ib.lastFourNumbers = s
	return ib
}

func (ib *auxBuilder) ValidThru(s string) InstrumentBuilder {
	ib.validThru = s
	return ib
}

func (ib *auxBuilder) Issuer(s string) InstrumentBuilder {
	ib.issuer = s
	return ib
}

func (ib *auxBuilder) PPS(s string) InstrumentBuilder {
	ib.pps = s
	return ib
}

func (ib *auxBuilder) CreditType(s string) InstrumentBuilder {
	ib.creditType = s
	return ib
}

func (ib *auxBuilder) Build() Instrument {
	return Instrument{
		AccountID:       ib.account,
		Holder:          ib.holder,
		LastFourNumbers: ib.lastFourNumbers,
		ValidThru:       ib.validThru,
		Issuer:          ib.issuer,
		PPS:             ib.issuer,
		CreditType:      ib.creditType,
	}
}



func New() InstrumentBuilder {
	return &auxBuilder{}
}

