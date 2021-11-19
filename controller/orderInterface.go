package controller

import "github.com/labstack/echo"

type OInterface interface {
	CreateOrder(c echo.Context) error
}
