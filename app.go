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

func NewApp(port string, pController controller.PCInterface, oController controller.OInterface, dController controller.IDashboardController) *App {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	pG := e.Group("/product")
	pG.POST("", pController.CreateProduct)
	pG.GET("", pController.GetAllProducts)
	pG.PUT("", pController.UpdateProduct)
	pG.DELETE("/:id", pController.DeleteProduct)
	pG.GET("/:id", pController.GetProduct)

	oG := e.Group("/order")
	oG.POST("", oController.CreateOrder)
	oG.GET("/all", oController.GetOrdersByDateRange)
	oG.GET("/:id", oController.GetOrderById)
	oG.GET("", oController.GetOrders)

	dG := e.Group("/dashboard")
	dG.GET("", dController.GetDashboardData)
	dG.GET("/chart", dController.SalesAndProfitGraph)

	return &App{
		port: port,
		e:    e,
	}
}

func (a *App) RunServer() {
	a.e.Logger.Fatal(a.e.Start(a.port))
}
