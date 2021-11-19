package controller

import (
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
	return nil
}
