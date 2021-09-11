package io

import (
	"github.com/gin-gonic/gin"
	"github.com/nawazish-github/opinion_server/model"
	"net/http"
)

func Response(c *gin.Context, args ...interface{}) {
	if args != nil {
		data := args[0].(model.QuestionAndOptions)//Todo: remove this smelly code - type casting
		c.JSON(http.StatusOK, data)
		return
	}
	c.Status(http.StatusOK)
}

func ErrResponse(c *gin.Context, status int, err error) {
	c.AbortWithError(status, err)
}
