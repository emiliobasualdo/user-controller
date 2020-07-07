package webapp

import (
	"crypto/sha256"
	b64 "encoding/base64"
	"encoding/json"
	"github.com/apsdehal/go-logger"
	"github.com/gin-gonic/gin"
	. "massimple.com/wallet-controller/pkg/models"
	"massimple.com/wallet-controller/pkg/services"
	"net/http"
)

type Response struct {
	Data		interface{}	`json:"data"`
	Signature	string		`json:"signature"`
}

type Message struct {
	Message string `json:"message"`
}


const PORT = "5000"

var log *logger.Logger

func Serve(_log *logger.Logger) {
	log = _log
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()
	root := router.Group("/")
	{
		root.GET("", statusHandler)
	}
	auth := router.Group("/auth")
	{
		auth.POST("/login", loginHandler)
	}
	me := router.Group("/me")
	{
		me.GET("/instruments", getInstrumentsHandler)
		me.POST("/instruments", insertInstrumentsHandler)
		me.DELETE("/instruments", deleteInstrumentsHandler)
		me.GET("/transactions", executionHandler)
		me.POST("/transactions", executionHandler)
	}
	if err := router.Run(":"+PORT); err != nil {
		panic(err)
	}
	log.NoticeF("Server started on port %s", PORT)
}

func respond(c *gin.Context, statusCode int, obj interface{}, err error) {
	var resp interface{}
	if obj == nil && err == nil {
		c.Status(statusCode)
		return
	} else if err == nil {
		resp = obj
	} else if obj == nil {
		resp = err.Error()
	}
	respByte,_ := json.Marshal(resp)
	signature := sha256.Sum256(respByte)
	c.JSON(statusCode, Response{resp, b64.StdEncoding.EncodeToString(signature[:])})
}

func statusHandler(c *gin.Context) {
	respond(c, http.StatusOK, gin.H{
		"status":  "ok",
		"version": "v1",
		"host": c.Request.Host,
	}, nil)
}

func loginHandler(c *gin.Context) {
	var login Login
	if err := c.BindJSON(&login); err != nil {
		respond(c, http.StatusBadRequest, nil, nil)
		return
	}
	account := services.GetAccount(login)
	respond(c, http.StatusOK, account, nil)
}

func getInstrumentsHandler(c *gin.Context) {
	body := struct {
		ID uint `json:"id"`
	}{}
	if err := c.BindJSON(&body); err != nil {
		respond(c, http.StatusBadRequest, "You must provide an ID", nil)
		return
	}
	instruments, err := services.GetInstrumentsById(body.ID)
	if err != nil {
		respond(c, http.StatusNotFound, nil, err)
		return
	}
	respond(c, http.StatusOK, instruments, nil)
}

func deleteInstrumentsHandler(c *gin.Context) {
	body := struct {
		ID uint `json:"id"`
	}{}
	if err := c.BindJSON(&body); err != nil {
		respond(c, http.StatusBadRequest, "You must provide an instrumentId", nil)
		return
	}
	err := services.DeleteInstrumentById(body.ID)
	if err != nil {
		respond(c, http.StatusNotFound, nil, err)
		return
	}
	respond(c, http.StatusOK, "Card deleted", nil)
}

func insertInstrumentsHandler(c *gin.Context) {
	var instrumentDto InstrumentDto
	if err := c.BindJSON(&instrumentDto); err != nil {
		respond(c, http.StatusBadRequest, Message{"You must provide an complete instrument"}, nil)
		return
	}
	instrument, err := services.InsertInstrumentById(instrumentDto)
	if err != nil {
		respond(c, http.StatusNotFound, nil, err)
		return
	}
	respond(c, http.StatusOK, instrument, nil)
}

func executionHandler(c *gin.Context) {
	resp := Message{"No implementado"}
	respond(c, http.StatusNotImplemented, resp, nil)
}