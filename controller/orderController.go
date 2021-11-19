package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/asishshaji/thalir-backend/models"
	"github.com/asishshaji/thalir-backend/services"
	"github.com/labstack/echo"
)

type OrderController struct {
	oS services.OServiceInterface
}

func NewOrderController(oS services.OServiceInterface) OInterface {
	return OrderController{
		oS: oS,
	}
}

func (oC OrderController) CreateOrder(c echo.Context) error {
	order := models.Order{}
	err := json.NewDecoder(c.Request().Body).Decode(&order)

	if err != nil {
		c.Logger().Error(err)
		return echo.ErrBadRequest
	}

	log.Println(order)

	o, err := oC.oS.CreateOrder(order)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseError{StatusCode: http.StatusBadRequest, Message: err.Error()})

	}

	return c.JSON(http.StatusCreated, models.ResponseSuccess{StatusCode: http.StatusCreated, Message: o})

}
