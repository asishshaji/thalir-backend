package main

import (
	"github.com/asishshaji/thalir-backend/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type App struct {
	e    *echo.Echo
	port string
}

func NewApp(port string, pController controller.PCInterface) *App {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	pG := e.Group("/product")
	pG.POST("", pController.CreateProduct)
	pG.GET("", pController.GetAllProducts)
	pG.PUT("", pController.UpdateProduct)
	pG.DELETE("/:id", pController.DeleteProduct)

	return &App{
		port: port,
		e:    e,
	}
}

func (a *App) RunServer() {
	a.e.Logger.Fatal(a.e.Start(a.port))
}
