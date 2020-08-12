package models

type Prepaid struct {
	ID uint 	`json:"id" example:"2010"`
	IDGP uint 	`json:"-" example:"4200"`
	Brand string `json:"QUILMES" example:"4200"`
}
