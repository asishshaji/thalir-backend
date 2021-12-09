package controller

import (
	"log"
	"net/http"

	"github.com/asishshaji/thalir-backend/models"
	"github.com/asishshaji/thalir-backend/services"
	"github.com/labstack/echo"
)

type DashboardController struct {
	dS services.IDashboardService
}

func NewDashboardController(dS services.IDashboardService) IDashboardController {
	return DashboardController{
		dS: dS,
	}
}

func (dC DashboardController) GetDashboardData(c echo.Context) error {
	start_date := c.QueryParam("start_date")
	end_date := c.QueryParam("end_date")

	dM, err := dC.dS.GetDashboardData(start_date, end_date)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, models.ResponseError{StatusCode: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, models.ResponseSuccess{StatusCode: http.StatusOK, Message: dM})
}

func (dC DashboardController) SalesAndProfitGraph(c echo.Context) error {
	start_date := c.QueryParam("start_date")
	end_date := c.QueryParam("end_date")
	gM, err := dC.dS.SalesAndProfitGraph(start_date, end_date)

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, models.ResponseError{StatusCode: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, models.ResponseSuccess{StatusCode: http.StatusOK, Message: gM})

}
