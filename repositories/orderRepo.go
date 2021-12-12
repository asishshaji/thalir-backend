package repository

import (
	"github.com/asishshaji/thalir-backend/models"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) ORInterface {
	return OrderRepository{
		db: db,
	}
}

func (oR OrderRepository) CreateOrder(o models.Order) (models.Order, error) {

	err := oR.db.Create(&o).Error
	if err != nil {
		return models.Order{}, err
	}
	return o, nil
}

func (oR OrderRepository) GetOrder(oId uint) ([]models.OrderItemDisplay, error) {
	odM := []models.OrderItemDisplay{}

	oR.db.Debug().Model(&models.Order{}).
		Select("p.name as product_name", "oi.units as units_bought", "oi.selling_price as price", "oi.profit as profit").
		Joins("Join order_items oi on oi.order_id = orders.id").
		Joins("Join products p on p.id = oi.product_id").
		Where("orders.id = ? ", oId).
		Find(&odM)

	return odM, nil
}

func (oR OrderRepository) GetOrders() ([]models.OrderDetails, error) {
	oM := []models.OrderDetails{}

	oR.db.Debug().Model(&models.Order{}).
		Select("orders.id as order_id", "orders.phone_number", "sum(profit) as profit", "orders.created_at").
		Joins("Join order_items oi on oi.order_id =  orders.id").
		Group("orders.id,orders.phone_number").
		Order("orders.id DESC").
		Limit(10).
		Find(&oM)

	return oM, nil
}

func (oR OrderRepository) GetOrdersByDateRange(startDate, endDate string) ([]models.OrderDetails, error) {
	oM := []models.OrderDetails{}

	oR.db.Debug().Model(&models.Order{}).
		Select("orders.id as order_id", "orders.phone_number", "sum(profit) as profit", "orders.created_at").
		Joins("Join order_items oi on oi.order_id =  orders.id").
		Group("orders.id,orders.phone_number").
		Order("orders.id DESC").
		Where("oi.created_at >= ? and oi.created_at < ?", startDate, endDate).
		Find(&oM)

	return oM, nil
}
