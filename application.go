package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UserProfile struct {
	Name    	string	`json:"name"`
	LastName    string	`json:"lastName"`
	PhoneNumber string	`json:"phoneNumber"`
}

type Card struct {
	Holder    string	`json:"holder"`
	Number    string	`json:"number"`
	ValidThru string	`json:"validThru"`
}

func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()
	root := router.Group("/")
	{
		root.GET("/", statusHandler)
	}
	auth := router.Group("/auth")
	{
		auth.POST("/login", loginHandler)
	}
	instruments := router.Group("/instruments")
	{
		instruments.GET("/", instrumentsHandler)
	}
	log.Fatal(router.Run(":5000"))
}

func statusHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"version": "v1",
	})
}

func loginHandler(c *gin.Context) {
	profile := UserProfile{"Adolfo", "Olivera", "5491133071114"}
	c.JSON(http.StatusOK, profile)
	return
}

func instrumentsHandler(c *gin.Context) {
	card1 := Card{"Adolfo Olivera","5555666677778888","07/25"}
	card2 := Card{"Adolfo Olivera","1111222233334444","06/22"}
	cards := []Card{card1, card2}
	c.JSON(http.StatusOK, cards)
}


