package models

type Prepaid struct {
	ID    ID     `json:"id" `
	IDGP  IDGP   `json:"-" example:"4200"`
	Brand string `json:"QUILMES" example:"4200"`
}
