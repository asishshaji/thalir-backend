package controller

import "github.com/labstack/echo"

type PCInterface interface {
	CreateProduct(c echo.Context) error
	GetAllProducts(c echo.Context) error
	UpdateProduct(c echo.Context) error
	DeleteProduct(c echo.Context) error
}
