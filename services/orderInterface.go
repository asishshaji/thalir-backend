package services

import "github.com/asishshaji/thalir-backend/models"

type OServiceInterface interface {
	CreateOrder(order models.Order) (models.Order, error)
}
