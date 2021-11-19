package repository

type ORInterface interface {
	CreateProduct(o interface{}) (interface{}, error)
}
