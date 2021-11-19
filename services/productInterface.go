package services

type PSInterface interface {
	CreateProduct(product interface{}) error
	GetAllProducts() (interface{}, error)
}
