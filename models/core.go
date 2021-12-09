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

func (p *Product) ToBackup() ProductBackup {
	return ProductBackup{
		Name:      p.Name,
		Type:      p.Type,
		BuyPrice:  p.BuyPrice,
		SellPrice: p.SellPrice,
		Units:     p.Units,
	}
}

type ProductBackup struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	BuyPrice  float32   `json:"buy_price"`
	SellPrice float32   `json:"sell_price"`
	Units     float32   `json:"units"`
}

type Order struct {
	ID          uint        `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time   `json:"updated_at" gorm:"autoUpdateTime"`
	PhoneNumber string      `json:"phone_number"`
	OrderItems  []OrderItem `json:"order_items"`
}

func (o *Order) ConvertOrder() Order {
	orderDetails := o.OrderItems
	var ords []OrderItem
	for _, de := range orderDetails {
		ord := OrderItem{
			OrderId:   de.OrderId,
			ProductId: de.ProductId,
			Units:     de.Units,
			Profit:    de.Profit,
		}
		ords = append(ords, ord)
	}
	return Order{
		PhoneNumber: o.PhoneNumber,
		OrderItems:  ords,
	}
}

type OrderItem struct {
	ID           uint      `gorm:"primarykey,default:uuid_generate_v3()" json:"id"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	OrderId      uint      `json:"order_id"`
	ProductId    int       `json:"p_id"`
	Units        float32   `json:"units"`
	Profit       float32   `json:"profit"`
	BuyingPrice  float32   `json:"buying_price"`
	SellingPrice float32   `json:"selling_price"`
}

type DashboardData struct {
	ProfitToday float32 `json:"profit"`
	OrdersToday float32 `json:"orders"`
	SalesToday  float32 `json:"sales"`
}

type SalesAndProfitGraph struct {
	SalesData  []float32   `json:"sales_list"`
	ProfitData []float32   `json:"profit_list"`
	Dates      []time.Time `json:"dates"`
}

type OrderDetails struct {
	OrderId     int       `json:"order_id"`
	PhoneNumber string    `json:"phone_number"`
	Profit      float32   `json:"profit"`
	CreatedAt   time.Time `json:"created_at"`
}
