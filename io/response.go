package io

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response (c *gin.Context) {
	c.Status(http.StatusOK)
}

func ErrResponse (c *gin.Context, err error) {
	c.AbortWithError(http.StatusBadRequest, err)
}
