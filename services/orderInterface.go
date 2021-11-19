package services

type OServiceInterface interface {
	CreateOrder(order interface{}) (interface{}, error)
}
