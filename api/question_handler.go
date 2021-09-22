package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nawazish-github/opinion_server/database"
	"github.com/nawazish-github/opinion_server/io"
	"net/http"
)

func GetQuestion (c *gin.Context) {
	fmt.Println("Get question with options handler initiated")
	date := c.Param("date")
	fmt.Println("Fetch question for the date: ", date)
	m, err := database.GetQuestion(c, date)
	if err != nil {
		io.ErrResponse(c, http.StatusInternalServerError, err)
		return
	}

	io.Response(c, m)
}
