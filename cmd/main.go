package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/llmons/who-is-the-undercover/internal/base/conf"
	"github.com/llmons/who-is-the-undercover/internal/router"
)

const addr = "localhost:8080"

func main() {
	c, err := conf.ReadConfig("config.yaml")
	if err != nil {
		panic(err)
	}

	app, cleanup, err := initApplication(c.Debug, c.Server, c.Data.Database)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting server at", addr)

	defer cleanup()
	if err = app.Run(addr); err != nil {
		panic(err)
	}
}

func NewApp(router *router.APIRouter) *gin.Engine {
	r := gin.Default()
	group := r.Group("")
	router.RegisterRouter(group)
	return r
}
