package models

import (
	"time"
)

type Login struct {
	PhoneNumber	string `json:"phoneNumber" binding:"required"`
	Code		string `json:"code" binding:"required"`
}

type Account struct {
	ID            uint         	`json:"id" gorm:"primary_key"`
	Name          string        `json:"name"`
	LastName      string        `json:"lastName"`
	PhoneNumber   string        `json:"phoneNumber" gorm:"unique:not null"`
	Instruments   []Instrument  `json:"instruments" gorm:"foreignkey:AccountID"`
	Beneficiaries []Beneficiary `json:"beneficiaries"`
	Balance       float64       `json:"balance"`
	CreatedAt     time.Time     `json:"createdAt"`
	UpdatedAt     time.Time     `json:"updatedAt"`
}

type Beneficiary struct {
	AccountID	    int          	`json:"accountId"`
	Name    		string          `json:"name"`
	LastName    	string          `json:"lastName"`
}

const (
	DEBIT   = "DEBIT"
	CREDIT  = "CREDIT"
	PREPAID = "PREPAID"
)
