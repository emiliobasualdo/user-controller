package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Account struct {
	ID     			string			 	`json:"id" bson:"id,omitempty" example:"5adf87asdfa7s8df6a9"`
	MongoID     	primitive.ObjectID 	`json:"-" bson:"_id,omitempty" example:"5adf87asdfa7s8df6a9"`
	GPID            string		  		`json:"-" example:"58756"`
	Name          	string        		`json:"name" example:"Mónica"`
	LastName      	string        		`json:"lastName" example:"Potrelli de Argento"`
	PhoneNumber   	string        		`json:"phoneNumber" gorm:"unique:not null" example:"+5491133071114"`
	DNI				string        		`json:"dni" example:"21399433"`
	Email			string        		`json:"email" example:"moni.argento@hotmail.com"`
	Instruments   	[]Instrument  		`json:"instruments" gorm:"foreignkey:AccountID"`
	Prepaids	 	[]Prepaid 			`json:"prepaids"`
	Transactions 	[]Transaction 		`json:"transactions"`
	Balance       	float64       		`json:"balance" example:"5430.54"`
	Disabled		bool				`json:"-"`
	CreatedAt     	time.Time     		`json:"createdAt" bson:"createdAt,omitempty" example:"2020-07-07T11:38:09.157803072Z"`
	UpdatedAt     	time.Time     		`json:"-" bson:"updatedAt,omitempty"`
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
