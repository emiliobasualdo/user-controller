package dtos

import "time"

type LoginDto struct {
	PhoneNumber string `json:"phoneNumber" binding:"required" example:"+5491133071114"`
	Code 		string `json:"code" binding:"required" example:"123654"`
}

type PhoneNumberDto struct {
	PhoneNumber string `json:"phoneNumber" binding:"required" example:"+5491133071114"`
}

type TokenDto struct {
	Code	int 		`json:"code" example:"200"`
	Expire	time.Time 	`json:"expire" example:"2020-07-08T15:58:45+02:00"`
	Token	string 		`json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTQyMTY3MjUsIm9yaWdfaWF0IjoxNTk0MjEzMTI1fQ.tWsDdREGVc2dPW7ZrcsoastWqfZm0s0w-oy6w0jH7YI"`
}
