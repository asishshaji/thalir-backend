package controller

import "github.com/labstack/echo"

type OInterface interface {
	CreateOrder(c echo.Context) error

	GetOrderById(c echo.Context) error
	GetOrdersByDateRange(c echo.Context) error
	GetOrders(c echo.Context) error
}
