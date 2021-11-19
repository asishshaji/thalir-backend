package repository

import "github.com/uptrace/bun"

type OrderRepository struct {
	db *bun.DB
}

func NewOrderRepository(db *bun.DB) ORInterface {
	return OrderRepository{
		db: db,
	}
}

func (oR OrderRepository) CreateProduct(o interface{}) (interface{}, error) {
	return nil, nil
}
