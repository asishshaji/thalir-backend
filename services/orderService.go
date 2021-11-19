package services

import repository "github.com/asishshaji/thalir-backend/repositories"

type OrderService struct {
	oR repository.ORInterface
}

func NewOrderService(oR repository.ORInterface) OServiceInterface {
	return OrderService{
		oR: oR,
	}
}
func (oS OrderService) CreateOrder(order interface{}) (interface{}, error) {
	return nil, nil
}
