package router

import (
	"github.com/gin-gonic/gin"
	"github.com/llmons/havefun/internal/controller"
)

type APIRouter struct {
	controller *controller.UndercoverController
}

func NewAPIRouter(controller *controller.UndercoverController) *APIRouter {
	return &APIRouter{
		controller: controller,
	}
}

func (a *APIRouter) RegisterRouter(r *gin.RouterGroup) {
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "test"})
	})
	r.GET("/undercover/wordpairs", a.controller.GetAllWordPairs)
}
