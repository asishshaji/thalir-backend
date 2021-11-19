package repository

type PRInterface interface {
	CreateProduct(p interface{}) error
	GetAllProducts() (interface{}, error)
}
