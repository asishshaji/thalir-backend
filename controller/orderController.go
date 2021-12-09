package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/asishshaji/thalir-backend/models"
	"github.com/asishshaji/thalir-backend/services"
	"github.com/kr/pretty"
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

	o, err := oC.oS.CreateOrder(order)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseError{StatusCode: http.StatusBadRequest, Message: err.Error()})

	}

	return c.JSON(http.StatusCreated, models.ResponseSuccess{StatusCode: http.StatusCreated, Message: o})

}

func (oC OrderController) GetOrderById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseError{StatusCode: http.StatusBadRequest, Message: "error parsing id"})
	}
	o, _ := oC.oS.GetOrder(uint(id))
	pretty.Println(o)
	return c.JSON(http.StatusOK, models.ResponseSuccess{
		StatusCode: http.StatusOK,
		Message:    o,
	})
}

func (oC OrderController) GetOrders(c echo.Context) error {

	oM, err := oC.oS.GetOrders()
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseError{StatusCode: http.StatusBadRequest, Message: err.Error()})

	}
	return c.JSON(http.StatusCreated, models.ResponseSuccess{StatusCode: http.StatusCreated, Message: oM})

}

func (oC OrderController) GetOrdersByDateRange(c echo.Context) error {
	start_date := c.QueryParam("start_date")
	end_date := c.QueryParam("end_date")
	oM, err := oC.oS.GetOrdersByDateRange(start_date, end_date)
	fmt.Println(err)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseError{StatusCode: http.StatusBadRequest, Message: err.Error()})

	}
	return c.JSON(http.StatusCreated, models.ResponseSuccess{StatusCode: http.StatusCreated, Message: oM})

}
