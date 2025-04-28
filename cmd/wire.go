// go:build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/llmons/who-is-the-undercover/internal/base/conf"
	"github.com/llmons/who-is-the-undercover/internal/base/data"
	"github.com/llmons/who-is-the-undercover/internal/controller"
	"github.com/llmons/who-is-the-undercover/internal/repo"
	"github.com/llmons/who-is-the-undercover/internal/router"
	"github.com/llmons/who-is-the-undercover/internal/service"
)

func initApplication(
	debug bool,
	serverConf *conf.Server,
	dbConfig *data.Database,
) (*gin.Engine, func(), error) {
	panic(wire.Build(
		controller.ProviderSetController,
		service.ProviderSetService,
		repo.ProviderSetRepo,
		router.ProviderSetRouter,
		NewApp,
	))
}
