package controller

import "github.com/labstack/echo"

type IDashboardController interface {
	GetDashboardData(c echo.Context) error
	SalesAndProfitGraph(c echo.Context) error
}
