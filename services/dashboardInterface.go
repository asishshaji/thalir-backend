package services

import "github.com/asishshaji/thalir-backend/models"

type IDashboardService interface {
	GetDashboardData(start_date, end_date string) (models.DashboardData, error)
	SalesAndProfitGraph(start_date, end_date string) (models.SalesAndProfitGraph, error)
}
