package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const addr = "localhost:8080"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(addr)
}
