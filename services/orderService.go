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
		elem.OrderId = int(order.ID)

		productModel, _ := oS.pR.GetProduct(elem.ProductId)
		profit := (productModel.SellPrice - productModel.BuyPrice) * elem.Units
		elem.Profit = profit
	}

	order, err := oS.oR.CreateOrder(order)
	if err != nil {
		return models.Order{}, err

	}
	return order, nil
}

func (oS OrderService) GetDashboardData() (models.DashboardData, error) {

	return oS.oR.GetDashboardData()
}

func (oS OrderService) GetProfitsBasedOnDateRange(d1, d2 string) (float32, float32, error) {
	return oS.oR.GetProfitsBasedOnDateRange(d1, d2)
}

func (oS OrderService) GetOrder(oId uint) (models.Order, error) {
	return oS.oR.GetOrder(oId)
}
