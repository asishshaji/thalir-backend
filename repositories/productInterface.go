package repository

import "github.com/asishshaji/thalir-backend/models"

type PRInterface interface {
	CreateProduct(p models.Product) (interface{}, error)
	GetAllProducts() (interface{}, error)
	UpdateProduct(p models.Product) error
	DeleteProduct(pid int) error
	GetProduct(pid int) (models.Product, error)
}
