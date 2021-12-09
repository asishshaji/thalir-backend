package services

import (
	"github.com/asishshaji/thalir-backend/models"
	repository "github.com/asishshaji/thalir-backend/repositories"
)

type DashboardService struct {
	dR repository.IDashboardRepo
}

func NewDashboardService(dR repository.IDashboardRepo) IDashboardService {
	return DashboardService{
		dR: dR,
	}
}

func (dS DashboardService) GetDashboardData(start_date, end_date string) (models.DashboardData, error) {
	return dS.dR.GetDashboardData(start_date, end_date)
}

func (dS DashboardService) SalesAndProfitGraph(start_date, end_date string) (models.SalesAndProfitGraph, error) {
	return dS.dR.SalesAndProfitGraph(start_date, end_date)
}
