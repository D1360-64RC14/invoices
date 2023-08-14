package app

import (
	"github.com/gin-gonic/gin"
)

type App struct {
	engine *gin.Engine
}

func New(address ...string) *App {
	app := &App{
		engine: gin.Default(),
	}

	app.configureTemplating()
	app.configureGin()
	app.configureRouters()

	return app
}

func (a *App) Run(addr ...string) {
	a.engine.Run(addr...)
}
