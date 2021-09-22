package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nawazish-github/opinion_server/database"
	"github.com/nawazish-github/opinion_server/io"
	"github.com/nawazish-github/opinion_server/model"
	"net/http"
)

func ReceiveOpinion(c *gin.Context) {
	fmt.Println("/opinion endpoint called to receive opinion")
	var opinion model.Opinion

	if err := c.ShouldBindJSON(&opinion); err != nil {
		fmt.Printf("Invalid opinion received: %v \n", err)
		io.ErrResponse(c, http.StatusBadRequest, err)
		return
	}
	fmt.Printf("Successfully received opinion")

	ipAddr := getIPAddress(c)
	database.SaveOpinion(opinion, ipAddr)
	io.Response(c)
}

func getIPAddress(c *gin.Context) string {
	return c.ClientIP()
}
