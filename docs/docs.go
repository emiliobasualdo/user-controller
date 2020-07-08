// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "email": "ebasualdo@itba.edu.ar"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/login": {
            "post": {
                "description": "Returns a jwt token to use as authentication",
                "produces": [
                    "application/json"
                ],
                "summary": "Generate a jwt",
                "operationId": "Get User",
                "parameters": [
                    {
                        "description": "user's phone number and the received sms code",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.LoginDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.TokenDto"
                        }
                    },
                    "400": {
                        "description": "The phone number does not match the code",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": " \"id does not exist",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/sms-code": {
            "post": {
                "description": "Sends an sms to the specified phonenumber",
                "summary": "SMS auth",
                "operationId": "Get User",
                "parameters": [
                    {
                        "description": "user's phone number",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.PhoneNumberDto"
                        }
                    }
                ],
                "responses": {
                    "200": {},
                    "400": {
                        "description": "Something went wrong",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/me": {
            "get": {
                "description": "Returns a list of the available instruments uploaded by the client",
                "produces": [
                    "application/json"
                ],
                "summary": "Get available Instruments",
                "operationId": "Get Instruments",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Account"
                            }
                        }
                    },
                    "400": {
                        "description": "Illegal token",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": " \"no such user",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/me/instruments": {
            "get": {
                "description": "Returns a list of the available instruments uploaded by the client",
                "produces": [
                    "application/json"
                ],
                "summary": "Get available Instruments",
                "operationId": "Get Instruments",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Instrument"
                            }
                        }
                    },
                    "400": {
                        "description": "Illegal token",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": " \"no such user",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Inserts and instrument to the list of available user instruments\nReturn the instrument object with its id",
                "produces": [
                    "application/json"
                ],
                "summary": "Insert instrument",
                "operationId": "Insert instrument",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id of the user that requests the instruments",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Instrument to insert",
                        "name": "instrument",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.InstrumentDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Instrument"
                            }
                        }
                    },
                    "400": {
                        "description": "The id provided is illegal",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "id does not exist",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/me/instruments/{id}": {
            "delete": {
                "description": "Deletes one of the instruments available to the user",
                "produces": [
                    "text/plain"
                ],
                "summary": "Delete an Instrument",
                "operationId": "Delete Instruments",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id of the instrument to delete",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Card deleted",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": " \"The id provided is illegal",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": " \"id does not exist",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dtos.InstrumentDto": {
            "type": "object",
            "properties": {
                "creditType": {
                    "type": "string",
                    "enum": [
                        "DEBIT",
                        " CREDIT",
                        " PREPAID"
                    ],
                    "example": "DEBIT"
                },
                "holder": {
                    "type": "string",
                    "example": "José Pepe Argento"
                },
                "issuer": {
                    "type": "string",
                    "example": "Banco Itaú"
                },
                "lastFourNumbers": {
                    "type": "string",
                    "example": "4930"
                },
                "pps": {
                    "type": "string",
                    "enum": [
                        "VISA",
                        " AMEX",
                        " MC"
                    ],
                    "example": "VISA"
                },
                "validThru": {
                    "type": "string",
                    "example": "11/24"
                }
            }
        },
        "dtos.LoginDto": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string",
                    "example": "123654"
                },
                "phoneNumber": {
                    "type": "string",
                    "example": "+5491133071114"
                }
            }
        },
        "dtos.PhoneNumberDto": {
            "type": "object",
            "properties": {
                "phoneNumber": {
                    "type": "string",
                    "example": "+5491133071114"
                }
            }
        },
        "dtos.TokenDto": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 200
                },
                "expire": {
                    "type": "string",
                    "example": "2020-07-08T15:58:45+02:00"
                },
                "token": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTQyMTY3MjUsIm9yaWdfaWF0IjoxNTk0MjEzMTI1fQ.tWsDdREGVc2dPW7ZrcsoastWqfZm0s0w-oy6w0jH7YI"
                }
            }
        },
        "models.Account": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "number",
                    "example": 5430.54
                },
                "beneficiaries": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Beneficiary"
                    }
                },
                "createdAt": {
                    "type": "string",
                    "example": "2020-07-07T11:38:09.157803072Z"
                },
                "id": {
                    "type": "integer",
                    "example": 5
                },
                "instruments": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Instrument"
                    }
                },
                "lastName": {
                    "type": "string",
                    "example": "Potrelli de Argento"
                },
                "name": {
                    "type": "string",
                    "example": "Mónica"
                },
                "phoneNumber": {
                    "type": "string",
                    "example": "+5491133071114"
                }
            }
        },
        "models.Beneficiary": {
            "type": "object",
            "properties": {
                "accountId": {
                    "type": "integer",
                    "example": 8
                },
                "lastName": {
                    "type": "string",
                    "example": "Argento"
                },
                "name": {
                    "type": "string",
                    "example": "Alfio Coqui"
                }
            }
        },
        "models.Instrument": {
            "type": "object",
            "properties": {
                "accountId": {
                    "type": "integer",
                    "example": 3
                },
                "createdAt": {
                    "type": "string",
                    "example": "2020-07-07 13:36:15.738848+02:00"
                },
                "creditType": {
                    "type": "string",
                    "enum": [
                        "DEBIT",
                        " CREDIT",
                        " PREPAID"
                    ],
                    "example": "DEBIT"
                },
                "holder": {
                    "type": "string",
                    "example": "José Pepe Argento"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "issuer": {
                    "type": "string",
                    "example": "Banco Itaú"
                },
                "lastFourNumbers": {
                    "type": "string",
                    "example": "4930"
                },
                "pps": {
                    "type": "string",
                    "enum": [
                        "VISA",
                        " AMEX",
                        " MC"
                    ],
                    "example": "VISA"
                },
                "validThru": {
                    "type": "string",
                    "example": "11/24"
                }
            }
        }
    },
    "securityDefinitions": {
        "JWT-Bearer": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost:5000",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Más Simple Wallet-Controller API",
	Description: "This is the main server where wallet operations for the Más Simple Wallet will be received",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
