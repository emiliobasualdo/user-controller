package utils

import (
	"github.com/gin-gonic/gin"
)

func Respond(c *gin.Context, statusCode int, obj interface{}, err error) {
	var resp interface{}
	if obj == nil && err == nil {
		c.Status(statusCode)
		return
	} else if err == nil {
		resp = obj
	} else if obj == nil {
		resp = err.Error()
	}
	c.JSON(statusCode, resp)
}
