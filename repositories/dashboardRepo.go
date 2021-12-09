package repository

import (
	"fmt"
	"time"

	"github.com/asishshaji/thalir-backend/models"
	"gorm.io/gorm"
)

type DashboardRepo struct {
	db *gorm.DB
}

func NewDashboardRepo(db *gorm.DB) IDashboardRepo {
	return DashboardRepo{db: db}
}

func (dR DashboardRepo) GetDashboardData(start_date, end_date string) (models.DashboardData, error) {

	var sales, profit, orders float32

	rows, err := dR.db.Model(&models.OrderItem{}).Select("sum(profit),sum(selling_price * units)").Where("created_at >= ? and created_at < ?", start_date, end_date).Rows()

	if err != nil {
		return models.DashboardData{}, err
	}
	for rows.Next() {
		rows.Scan(&profit, &sales)
	}

	rows, err = dR.db.Model(&models.Order{}).Select("count(*)").Where("created_at >= ? and created_at < ?", start_date, end_date).Rows()

	if err != nil {
		return models.DashboardData{}, err
	}
	for rows.Next() {
		rows.Scan(&orders)
	}
	mDashboard := models.DashboardData{
		SalesToday:  sales,
		ProfitToday: profit,
		OrdersToday: orders,
	}

	return mDashboard, nil
}

func (dR DashboardRepo) SalesAndProfitGraph(start_date, end_date string) (models.SalesAndProfitGraph, error) {
	var sales, profit []float32
	var dates []time.Time
	m := models.SalesAndProfitGraph{}
	m.SalesData = []float32{}
	m.ProfitData = []float32{}
	m.Dates = []time.Time{}

	fmt.Println(m)

	rows, err := dR.db.Model(&models.OrderItem{}).Select("sum(profit),sum(selling_price * units) as sales,DATE(created_at) AS created_date").Where("created_at >= ? and created_at < ?", start_date, end_date).Group("created_date").Rows()
	if err != nil {
		fmt.Println(err)
		return models.SalesAndProfitGraph{}, err
	}

	var p, s float32
	fmt.Println(rows)
	rows.Scan(&p, &s)
	for rows.Next() {
		var p, s float32
		var d time.Time
		rows.Scan(&p, &s, &d)
		fmt.Println(p, s)
		sales = append(sales, s)
		profit = append(profit, p)
		dates = append(dates, d)
	}

	if len(sales) != 0 {
		m.Dates = dates
		m.SalesData = sales
		m.ProfitData = profit
	}

	return m, nil
}
