package app

import (
	"fmt"
	"log"
	mw "middleware/internal/app/MW"
	"middleware/internal/app/endpoint"
	"middleware/internal/app/service"

	"github.com/labstack/echo/v4"
)

type App struct {
	e    *endpoint.Endpoint
	s    *service.Service
	echo *echo.Echo
}

func New() (*App, error) {
	app := &App{}

	app.s = service.New()

	app.e = endpoint.New(app.s)

	app.echo = echo.New()

	app.echo.Use(mw.RoleCheck)

	app.echo.GET("/status", app.e.Status)

	return app, nil
}

func (app *App) Run() error {
	fmt.Println("server running")

	err := app.echo.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
