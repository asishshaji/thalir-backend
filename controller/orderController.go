package controller

import (
	"encoding/json"
	"fmt"
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

func (oC OrderController) GetDashboardData(c echo.Context) error {
	mb, err := oC.oS.GetDashboardData()
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseError{StatusCode: http.StatusBadRequest, Message: err.Error()})

	}
	return c.JSON(http.StatusOK, models.ResponseSuccess{StatusCode: http.StatusCreated, Message: mb})

}

func (oC OrderController) GetProfitsBasedOnDateRange(c echo.Context) error {
	d1 := c.QueryParam("start_date")
	d2 := c.QueryParam("end_date")
	units, profit, err := oC.oS.GetProfitsBasedOnDateRange(d1, d2)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseError{StatusCode: http.StatusBadRequest, Message: err.Error()})

	}
	return c.JSON(http.StatusOK, models.ResponseSuccess{StatusCode: http.StatusCreated, Message: map[string]float32{
		"profit": profit,
		"units":  units,
	}})
}

func (oC OrderController) GetOrderById(c echo.Context) error {
	o, _ := oC.oS.GetOrder(1)
	fmt.Println(o)
	return nil
}
