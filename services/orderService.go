package services

import (
	"errors"
	"fmt"

	"github.com/asishshaji/thalir-backend/models"
	repository "github.com/asishshaji/thalir-backend/repositories"
	"github.com/kr/pretty"
)

type OrderService struct {
	oR repository.ORInterface
	pR repository.PRInterface
}

func NewOrderService(oR repository.ORInterface, pR repository.PRInterface) OServiceInterface {
	return OrderService{
		oR: oR,
		pR: pR,
	}
}
func (oS OrderService) CreateOrder(order models.Order) (models.Order, error) {
	orderDetails := order.OrderItems
	ords := []models.OrderItem{}
	products := []models.Product{}
	for _, oD := range orderDetails {
		p, _ := oS.pR.GetProduct(oD.ProductId)
		if p.Units < oD.Units {
			return models.Order{}, errors.New("സ്റ്റോക്ക് ലഭ്യമല്ല : " + p.Name)
		}
		profit := oD.Units * (p.SellPrice - p.BuyPrice)
		oD.Profit = profit
		oD.BuyingPrice = p.BuyPrice
		oD.SellingPrice = p.SellPrice
		ords = append(ords, oD)
		pretty.Println(p.Units)
		p.Units = p.Units - oD.Units
		if p.Units == 0 {
			p.Units = 0.0
		}

		products = append(products, p)
	}

	order.OrderItems = ords
	oS.oR.CreateOrder(order)

	for _, prod := range products {
		oS.pR.UpdateProduct(prod)
		fmt.Println(prod)
	}

	return order, nil
}

func (oS OrderService) GetOrder(oId uint) ([]models.OrderItemDisplay, error) {
	return oS.oR.GetOrder(oId)
}

func (oS OrderService) GetOrders() ([]models.OrderDetails, error) {
	return oS.oR.GetOrders()
}

func (oS OrderService) GetOrdersByDateRange(startDate, endDate string) ([]models.OrderDetails, error) {
	return oS.oR.GetOrdersByDateRange(startDate, endDate)
}
