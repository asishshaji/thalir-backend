package repository

import "github.com/asishshaji/thalir-backend/models"

type ORInterface interface {
	CreateProduct(o models.Order) (models.Order, error)
}
