package api

import (
	"github.com/gin-gonic/gin"
	"github.com/nawazish-github/opinion_server/database"
	"github.com/nawazish-github/opinion_server/io"
	"net/http"
)

func GetQuestion (c *gin.Context) {
	m, err := database.GetQuestion(c)
	if err != nil {
		io.ErrResponse(c, http.StatusInternalServerError, err)
		return
	}

	io.Response(c, m)
}
