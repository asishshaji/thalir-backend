package controller

import (
	"encoding/json"
	"net/http"

	"github.com/asishshaji/thalir-backend/models"
	"github.com/asishshaji/thalir-backend/services"
	"github.com/labstack/echo"
)

type ProductController struct {
	productService services.PSInterface
}

func NewProductController(pS services.PSInterface) PCInterface {
	return ProductController{
		productService: pS,
	}
}

func (pC ProductController) CreateProduct(c echo.Context) error {
	reqProduct := models.NewEmptyProductCreationRequest()

	err := json.NewDecoder(c.Request().Body).Decode(&reqProduct)
	if err != nil {
		c.Logger().Error(err)
		return echo.ErrBadRequest
	}

	c.Logger().Print("Creating product :> ", reqProduct)

	err = pC.productService.CreateProduct(reqProduct)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseError{StatusCode: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, reqProduct)
}

func (pC ProductController) GetAllProducts(c echo.Context) error {
	products, err := pC.productService.GetAllProducts()
	if err != nil {
		return c.JSON(http.StatusNotFound, models.ResponseError{StatusCode: http.StatusBadRequest, Message: err.Error()})

	}

	return c.JSON(http.StatusOK, products)
}
