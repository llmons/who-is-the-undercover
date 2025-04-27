package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/llmons/havefun/internal/base/data"
	"github.com/llmons/havefun/internal/controller"
	"github.com/llmons/havefun/internal/repo"
	"github.com/llmons/havefun/internal/router"
	"github.com/llmons/havefun/internal/service"
)

func initApplication(
	debug bool,
	dbConfig *data.Database,
) (*gin.Engine, error) {
	panic(wire.Build(
		controller.ProviderSetController,
		service.ProviderSetService,
		repo.ProviderSetRepo,
		router.ProviderSetRouter,
	))
}
