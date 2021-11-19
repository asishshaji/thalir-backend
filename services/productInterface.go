package services

type PSInterface interface {
	CreateProduct(product interface{}) (interface{}, error)
	GetAllProducts() (interface{}, error)
	UpdateProduct(product interface{}) error
	DeleteProduct(pid int) error
}
