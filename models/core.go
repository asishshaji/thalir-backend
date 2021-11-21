package models

import (
	"time"
)

type Product struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Name      string    `json:"name" gorm:"unique"`
	Type      string    `json:"type"`
	BuyPrice  float32   `json:"buy_price"`
	SellPrice float32   `json:"sell_price"`
	Units     float32   `json:"units"`
}

type Order struct {
	ID           uint          `gorm:"primarykey" json:"id"`
	CreatedAt    time.Time     `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time     `json:"updated_at" gorm:"autoUpdateTime"`
	PhoneNumber  string        `json:"phone_number"`
	OrderDetails []OrderDetail `json:"order_details"`
}

type OrderDetail struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	OrderId   int       `json:"order_id"`
	ProductId int       `json:"p_id"`
	Units     float32   `json:"units"`
	Profit    float32   `json:"profit"`
}

type DashboardData struct {
	Profit                 float32 `json:"profit"`
	UnitsSold              float32 `json:"units"`
	AvgProfit              float32 `json:"avg_profit"`
	AvgUnits               float32 `json:"avg_units"`
	TotalMoneyInvested     float32 `json:"invested_money"`
	TotalAvailableProducts int64   `json:"product_count"`
}
