package repository

import "github.com/asishshaji/thalir-backend/models"

type ORInterface interface {
	CreateOrder(o models.Order) (models.Order, error)
	// GetDashboardData() (models.DashboardData, error)
	// GetDashboardDataBasedOnDateRange(d1, d2 string) (float32, float32, error)
	GetOrder(oId uint) (models.OrderDetails, error)
	GetOrders() ([]models.OrderDetails, error)
	GetOrdersByDateRange(startDate, endDate string) ([]models.OrderDetails, error)
}
