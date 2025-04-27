package main

import (
	"github.com/gin-gonic/gin"
	"github.com/llmons/havefun/internal/controller"
)

const addr = "localhost:8080"

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.Set("data", map[string]interface{}{
			"hello": "world",
		})
	})
	r.GET("/whoisthespy", controller.GetKey)

	r.Run(addr)
}
