package services

import "github.com/asishshaji/thalir-backend/models"

type OServiceInterface interface {
	CreateOrder(order models.Order) (models.Order, error)
	GetOrder(oId uint) ([]models.OrderItemDisplay, error)
	GetOrders() ([]models.OrderDetails, error)
	GetOrdersByDateRange(startDate, endDate string) ([]models.OrderDetails, error)
}
