package models


type Beneficiary struct {
	AccountID	    int          	`json:"accountId" example:"8"`
	Name    		string          `json:"name" example:"Alfio Coqui"`
	LastName    	string          `json:"lastName" example:"Argento"`
}

