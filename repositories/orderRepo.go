package repository

import (
	"context"

	"github.com/asishshaji/thalir-backend/models"
	"github.com/uptrace/bun"
)

type OrderRepository struct {
	db *bun.DB
}

func NewOrderRepository(db *bun.DB) ORInterface {
	return OrderRepository{
		db: db,
	}
}

func (oR OrderRepository) CreateProduct(o models.Order) (models.Order, error) {
	oR.db.NewInsert().Model(&o).Exec(context.Background())
	oR.db.NewInsert().Model(&o.OrderDetails).Exec(context.Background())

	return models.Order{}, nil
}
