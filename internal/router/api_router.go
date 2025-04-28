package router

import (
	"github.com/gin-gonic/gin"
	"github.com/llmons/who-is-the-undercover/internal/controller"
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
	r.GET("/undercover/wordpairs", a.controller.GetAllWordPairs)
}
