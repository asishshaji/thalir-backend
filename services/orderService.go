package services

import (
	"github.com/asishshaji/thalir-backend/models"
	repository "github.com/asishshaji/thalir-backend/repositories"
)

type OrderService struct {
	oR repository.ORInterface
}

func NewOrderService(oR repository.ORInterface) OServiceInterface {
	return OrderService{
		oR: oR,
	}
}
func (oS OrderService) CreateOrder(order models.Order) (models.Order, error) {

	for _, elem := range order.OrderHistory {
		elem.OrderId = order.OrderId
	}

	oS.oR.CreateProduct(order)
	return models.Order{}, nil
}
