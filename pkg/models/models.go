package models

import (
	"time"
)

type Account struct {
	ID            uint         	`json:"id" gorm:"primary_key" example:"5"`
	Name          string        `json:"name" example:"Mónica"`
	LastName      string        `json:"lastName" example:"Potrelli de Argento"`
	PhoneNumber   string        `json:"phoneNumber" gorm:"unique:not null" example:"+5491133071114"`
	Instruments   []Instrument  `json:"instruments" gorm:"foreignkey:AccountID"`
	Beneficiaries []Beneficiary `json:"beneficiaries"`
	Balance       float64       `json:"balance" example:"5430.54"`
	CreatedAt     time.Time     `json:"createdAt" example:"2020-07-07T11:38:09.157803072Z"`
	UpdatedAt     time.Time     `json:"-"`
}

type Beneficiary struct {
	AccountID	    int          	`json:"accountId" example:"8"`
	Name    		string          `json:"name" example:"Alfio Coqui"`
	LastName    	string          `json:"lastName" example:"Argento"`
}
