package repository

import (
	"fmt"

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

	return o, nil
}

func (oR OrderRepository) GetOrder(oId uint) (models.Order, error) {
	order := models.Order{}
	rows, err := oR.db.Table("orders").Joins("Join order_details on order_details.order_id = orders.id").Where("orders.id =? ", oId).Rows()
	if err != nil {
		return models.Order{}, err
	}

	defer rows.Close()

	// for rows.Next(){
	// 	oD := models.OrderDetail{}
	// 	o := mode
	// }

	fmt.Println(order.OrderDetails)
	return order, nil
}

func (oR OrderRepository) GetDashboardData() (models.DashboardData, error) {
	// execute raw query

	rows, err := oR.db.Table("orders").Joins("Join order_details on order_details.order_id = orders.id").
		Select("sum(order_details.units)", "sum(order_details.profit)", "avg(order_details.profit)", "avg(order_details.units)").Rows()

	if err != nil {
		return models.DashboardData{}, err
	}

	defer rows.Close()

	var profit, units, avgProfit, avgUnit, invested_money float32
	var pCount int64

	for rows.Next() {
		rows.Scan(&units, &profit, &avgProfit, &avgUnit)
	}

	rows, err = oR.db.Table("products").Select("count(*)", "sum(products.buy_price*TotalAvailableProducts.units)").Rows()
	if err != nil {
		return models.DashboardData{}, err
	}

	for rows.Next() {
		rows.Scan(&pCount, &invested_money)
	}

	mb := models.DashboardData{
		Profit:                 profit,
		UnitsSold:              units,
		AvgProfit:              avgProfit,
		AvgUnits:               avgUnit,
		TotalMoneyInvested:     invested_money,
		TotalAvailableProducts: pCount,
	}

	return mb, nil

}

func (oR OrderRepository) GetProfitsBasedOnDateRange(d1, d2 string) (float32, float32, error) {

	rows, err := oR.db.Table("orders").Joins("Join order_details on order_details.order_id = orders.id").
		Select("sum(order_details.units)", "sum(order_details.profit)").Where("orders.created_at BETWEEN ? AND ?", d1, d2).Rows()
	var profit, units float32

	if err != nil {
		return 0, 0, err
	}

	defer rows.Close()

	for rows.Next() {
		rows.Scan(&units, &profit)
	}

	return units, profit, nil
}
