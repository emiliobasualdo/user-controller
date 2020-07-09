package models

import "time"

type Instrument struct {
	ID            	uint		`json:"id" gorm:"primary_key" example:"1"`
	AccountID		uint		`json:"accountId" example:"3"`
	Holder    		string		`json:"holder" example:"José Pepe Argento"`
	LastFourNumbers string		`json:"lastFourNumbers" example:"4930"`
	ValidThru 		string		`json:"validThru" example:"11/24"`
	Issuer	 		string		`json:"issuer" example:"Banco Itaú"`
	PPS 			string		`json:"pps" example:"VISA" enums:"VISA, AMEX, MC"`
	CreditType 		string		`json:"creditType" example:"DEBIT" enums:"DEBIT, CREDIT, PREPAID"`
	CreatedAt     	time.Time	`json:"createdAt" example:"2020-07-07 13:36:15.738848+02:00"`
	DisabledAt		time.Time	`json:"-"` // we dont export this field
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

type iBuilder struct {
	account			uint
	holder    		string
	lastFourNumbers string
	validThru 		string
	issuer	 		string
	pps				string
	creditType 		string
}

func (ib *iBuilder) FromAccount(acc uint) InstrumentBuilder {
	ib.account = acc
	return ib
}

func (ib *iBuilder) Holder(s string) InstrumentBuilder {
	ib.holder = s
	return ib
}

func (ib *iBuilder) LastFourNumbers(s string) InstrumentBuilder {
	ib.lastFourNumbers = s
	return ib
}

func (ib *iBuilder) ValidThru(s string) InstrumentBuilder {
	ib.validThru = s
	return ib
}

func (ib *iBuilder) Issuer(s string) InstrumentBuilder {
	ib.issuer = s
	return ib
}

func (ib *iBuilder) PPS(s string) InstrumentBuilder {
	ib.pps = s
	return ib
}

func (ib *iBuilder) CreditType(s string) InstrumentBuilder {
	ib.creditType = s
	return ib
}

func (ib *iBuilder) Build() Instrument {
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

func NewInstrumentBuilder() InstrumentBuilder {
	return &iBuilder{}
}

