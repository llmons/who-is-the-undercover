package controller

import (
	"github.com/gin-gonic/gin"
)

func GetKey(ctx *gin.Context) {
	ctx.Set("data", map[string]interface{}{
		"who": "spy",
	})
}
