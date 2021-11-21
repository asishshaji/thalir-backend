package services

import "github.com/asishshaji/thalir-backend/models"

type PSInterface interface {
	CreateProduct(product models.Product) (interface{}, error)
	GetAllProducts() (interface{}, error)
	UpdateProduct(product models.Product) error
	DeleteProduct(pid int) error
	GetProduct(pid int) (models.Product, error)
}
