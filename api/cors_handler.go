package api

import "github.com/gin-gonic/gin"

func HandleCORS(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Content-Type")
	c.Writer.Header().Set("Content-Type", "application/json")
}
