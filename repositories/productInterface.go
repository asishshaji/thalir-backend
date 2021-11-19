package repository

type PRInterface interface {
	CreateProduct(p interface{}) (interface{}, error)
	GetAllProducts() (interface{}, error)
	UpdateProduct(p interface{}) error
	DeleteProduct(pid int) error
}
