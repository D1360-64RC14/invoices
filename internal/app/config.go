package app

import (
	"text/template"

	"github.com/D1360-64RC14/invoices/internal/functions"
	"github.com/D1360-64RC14/invoices/internal/routers"
)

func (a *App) configureTemplating() {
	a.engine.SetFuncMap(template.FuncMap{
		"dateFormat":  functions.DateFormat,
		"priceFormat": functions.PriceFormat,
		"reverse":     functions.Reverse,
	})

	a.engine.LoadHTMLGlob("web/templates/*.html")
}

func (a *App) configureGin() {
	a.engine.Static("/static", "web/static")
}

func (a *App) configureRouters() {
	routers.NewDefaultRouter().AttachTo(&a.engine.RouterGroup)
}
