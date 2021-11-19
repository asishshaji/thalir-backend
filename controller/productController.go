package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

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

	p, err := pC.productService.CreateProduct(reqProduct)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseError{StatusCode: http.StatusBadRequest, Message: err.Error()})
	}
	// instead of reqProduct pass creaed dto from db
	return c.JSON(http.StatusCreated, p)
}

func (pC ProductController) GetAllProducts(c echo.Context) error {
	products, err := pC.productService.GetAllProducts()
	if err != nil {
		return c.JSON(http.StatusNotFound, models.ResponseError{StatusCode: http.StatusBadRequest, Message: err.Error()})

	}

	return c.JSON(http.StatusOK, products)
}

func (pC ProductController) UpdateProduct(c echo.Context) error {
	reqProduct := models.NewEmptyProductCreationRequest()

	err := json.NewDecoder(c.Request().Body).Decode(&reqProduct)
	if err != nil {
		c.Logger().Error(err)
		return echo.ErrBadRequest
	}
	err = pC.productService.UpdateProduct(reqProduct)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseError{StatusCode: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, models.ResponseSuccess{StatusCode: http.StatusOK, Message: "Updated product"})
}

func (pC ProductController) DeleteProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseError{StatusCode: http.StatusBadRequest, Message: err.Error()})

	}
	err = pC.productService.DeleteProduct(id)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, models.ResponseError{StatusCode: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusAccepted, models.ResponseSuccess{StatusCode: http.StatusAccepted, Message: "Product deleted"})

}
