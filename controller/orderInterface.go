package controller

import "github.com/labstack/echo"

type OInterface interface {
	CreateOrder(c echo.Context) error
	GetDashboardData(c echo.Context) error
	GetProfitsBasedOnDateRange(c echo.Context) error
	GetOrderById(c echo.Context) error
}
