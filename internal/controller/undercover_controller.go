package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/llmons/who-is-the-undercover/internal/service"
)

type UndercoverController struct {
	service *service.UndercoverService
}

func NewUndercoverController(service *service.UndercoverService) *UndercoverController {
	return &UndercoverController{
		service: service,
	}
}

func (c *UndercoverController) GetAllWordPairs(ctx *gin.Context) {
	wordPairList, err := c.service.GetAllWordPairs()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	ctx.JSON(200, wordPairList)
}

func (c *UndercoverController) GetRandomWordPair(ctx *gin.Context) {
	wordPair, err := c.service.GetRandomWordPair()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	ctx.JSON(200, wordPair)
}
