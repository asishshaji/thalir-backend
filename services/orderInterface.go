package services

import "github.com/asishshaji/thalir-backend/models"

type OServiceInterface interface {
	CreateOrder(order models.Order) (models.Order, error)
	GetDashboardData() (models.DashboardData, error)
	GetProfitsBasedOnDateRange(d1 string, d2 string) (float32, float32, error)
	GetOrder(oId uint) (models.Order, error)
}
