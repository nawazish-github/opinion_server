package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nawazish-github/opinion_server/api"
)

func InitRoutes() {
	router := gin.Default()
	router.POST("/opinion", api.ReceiveOpinion)
	router.OPTIONS("/opinion", api.HandleCORS)
	router.GET("/question/:date", api.GetQuestion)
	fmt.Println("server started on 8080")
	router.Run("localhost:8080")
}
