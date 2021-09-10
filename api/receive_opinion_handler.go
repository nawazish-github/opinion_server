package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nawazish-github/opinion_server/database"
	"github.com/nawazish-github/opinion_server/io"
	"github.com/nawazish-github/opinion_server/model"
)

func ReceiveOpinion(c *gin.Context) {
	var opinion model.Opinion

	if err := c.ShouldBindJSON(&opinion); err != nil {
		fmt.Printf("Invalid opinion received: %v \n", err)
		io.ErrResponse(c, err)
		return
	}
	fmt.Printf("successfully received opinion: %v", opinion)
	ipAddr := getIPAddress(c)
	opinion.IPAddress = ipAddr
	database.SaveOpinion(opinion)
	io.Response(c)
}

func getIPAddress (c *gin.Context) string {
	return c.ClientIP()
}