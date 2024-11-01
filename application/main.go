package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/sample", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "rust こそしこう"})
	})
	r.Run() // Start the server on default port 8080
}
