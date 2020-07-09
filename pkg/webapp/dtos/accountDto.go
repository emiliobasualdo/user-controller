package dtos

import . "massimple.com/wallet-controller/pkg/models"
type AccountDto struct {
	Name        string        `json:"name" binding:"required" example:"Mónica"`
	LastName    string        `json:"lastName" binding:"required" example:"Potrelli de Argento"`
	DNI			string        `json:"dni" binding:"required" example:"21399433"`
	Email		string        `json:"email" binding:"required" example:"moni.argento@hotmail.com"`
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