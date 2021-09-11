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
	var opinion model.Opinion

	if err := c.ShouldBindJSON(&opinion); err != nil {
		fmt.Printf("Invalid opinion received: %v \n", err)
		io.ErrResponse(c, http.StatusBadRequest, err)
		return
	}
	fmt.Printf("successfully received opinion: %v", opinion)

	poplOpinion := populateIPAddress(c, opinion)
	database.SaveOpinion(poplOpinion)
	io.Response(c)
}

func getIPAddress(c *gin.Context) string {
	return c.ClientIP()
}

func populateIPAddress(c *gin.Context, opinion model.Opinion) model.Opinion {
	ipAddr := getIPAddress(c)
	opinion.IPAddress = ipAddr
	return opinion
}
