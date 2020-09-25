package models

import "time"

type SliderCard struct {
	Image string `json:"image" example:"www.s3.com/image"`
	Title string `json:"title" example:"30% de descuento con débito"`
	Action string `json:"action" example:"redirect_screen?screenId=123"`
}

type Movement struct {
	Amount float64 `json:"amount" example:"302.15"`
	Type string `json:"type" example:"OUT"`
	Action string `json:"action" example:"carga"`
	Date time.Time `json:"date" example:"2020-07-07T11:38:09.157803072Z"`
	Extra string `json:"extra" example:"+ $56 ahorro"`
	Commerce string `json:"commerce" example:"Quilmes"`
	Link string `json:"link" example:"http//www.quilmes.com"`
	StatusText string `json:"statusText" example:"Transacción confirmada"`
}

type Summary struct {
	Balance float64 `json:"balance" example:"3022.12"`
	SliderCards []SliderCard `json:"sliderCards"`
	LastMovements []Movement `json:"lastMovements"`
}
