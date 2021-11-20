package services

import (
	"github.com/asishshaji/thalir-backend/models"
	repository "github.com/asishshaji/thalir-backend/repositories"
)

type OrderService struct {
	oR repository.ORInterface
	pR repository.PRInterface
}

func NewOrderService(oR repository.ORInterface, pR repository.PRInterface) OServiceInterface {
	return OrderService{
		oR: oR,
	}
}
func (oS OrderService) CreateOrder(order models.Order) (models.Order, error) {

	// calculate profit here

	orderDetails := order.OrderDetails

	for _, elem := range orderDetails {
		elem.OrderId = order.OrderId

		pI, _ := oS.pR.GetProduct(elem.ProductId)
		productModel := models.InterfaceToProduct(pI)
		profit := (productModel.SellPrice - productModel.BuyPrice) * elem.Units
		elem.Profit = profit
	}

	oS.oR.CreateProduct(order)
	return models.Order{}, nil
}
