package main

import (
	"github.com/gin-gonic/gin"
)

const addr = "localhost:8080"

func main() {
	r := gin.Default()

	r.Run(addr)
}
