package dtos

import (
	"massimple.com/wallet-controller/internal/models"
	"time"
)

type LoginDto struct {
	PhoneNumber models.PhoneNumber `json:"phoneNumber" binding:"required" example:"005491133071114"`
	Code 		string `json:"code" binding:"required" example:"123654"`
}

type TokenDto struct {
	Code	int 		`json:"code" example:"200"`
	Expire	time.Time 	`json:"expire" example:"2020-07-08T15:58:45+02:00"`
	Token	string 		`json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTQyMTY3MjUsIm9yaWdfaWF0IjoxNTk0MjEzMTI1fQ.tWsDdREGVc2dPW7ZrcsoastWqfZm0s0w-oy6w0jH7YI"`
}

type PrepaidDto struct {
	ID models.ID 	`json:"id" `
	Brand string `json:"QUILMES" example:"4200"`
	DisponibleCompra float64 `json:"availableLocal" example:"230.45"`
	DisponibleAnticipo float64 `json:"availableLocalInAdvanced" example:"230.45"`
	DisponibleCompraDolar float64 `json:"availableDollar" example:"12.34"`
	DisponibleAnticipoDolar float64 `json:"availableDollarInAdvanced" example:"12.34"`
	SaldoPesos float64 `json:"creditLocal" example:"-2850.63"`
	SaldoDolar float64 `json:"creditDollar" example:"-23.54"`
}
