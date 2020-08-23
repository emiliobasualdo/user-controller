package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ID string

type ToStringInterface interface {
	String()	string
}

func (id *ID) String() string {
	return string(*id)
}

type PhoneNumber string

func (p PhoneNumber) String() string {
	return string(p)
}

type Account struct {
	ID           ID                 `json:"id" bson:"id"`
	MongoID      primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	GPID         IDGP	            `json:"-" example:"58756"`
	Name         string             `json:"name" example:"Mónica"`
	LastName     string             `json:"lastName" example:"Potrelli de Argento"`
	PhoneNumber  PhoneNumber        `json:"phoneNumber" example:"005491133071114" bson:"phoneNumber"`
	DNI          string             `json:"dni" example:"21399433"`
	Email        string             `json:"email" example:"moni.argento@hotmail.com"`
	Instruments  []Instrument       `json:"instruments"`
	Prepaids     []Prepaid          `json:"prepaids"`
	Transactions []Transaction      `json:"transactions"`
	CreatedAt    time.Time     		`json:"createdAt" bson:"createdAt,omitempty" example:"2020-07-07T11:38:09.157803072Z"`
	UpdatedAt    time.Time     		`json:"-" bson:"updatedAt,omitempty"`
}

type AccountBuilder interface {
	Name(string)		AccountBuilder
	LastName(string)	AccountBuilder
	DNI(string)			AccountBuilder
	Email(string)		AccountBuilder
	Build() 			Account
}

type aBuilder struct {
	name          	string
	lastName      	string
	dni				string
	email			string
}

func (ab *aBuilder) Name(s string) AccountBuilder {
	ab.name = s
	return ab
}

func (ab *aBuilder) LastName(s string) AccountBuilder {
	ab.lastName = s
	return ab
}

func (ab *aBuilder) DNI(s string) AccountBuilder {
	ab.dni = s
	return ab
}

func (ab *aBuilder) Email(s string) AccountBuilder {
	ab.email = s
	return ab
}

func (ab *aBuilder) Build() Account {
	return Account{
		Name:     ab.name,
		LastName: ab.lastName,
		DNI:      ab.dni,
		Email:    ab.email,
	}
}

func NewAccountBuilder() AccountBuilder {
	return &aBuilder{}
}
