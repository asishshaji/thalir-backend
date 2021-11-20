package models

import (
	"reflect"
	"time"

	"github.com/asishshaji/thalir-backend/enum"
	"github.com/uptrace/bun"
)

type Product struct {
	bun.BaseModel `bun:"products"`
	Pid           int       `json:"p_id" bun:"p_id,pk"`
	Name          string    `json:"name"`
	Type          string    `json:"type"`
	BuyPrice      float32   `json:"buy_price"`
	SellPrice     float32   `json:"sell_price"`
	Units         float32   `json:"units"`
	CreatedAt     time.Time `json:"created_at" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time `json:"updated_at" bun:",nullzero,notnull,default:current_timestamp"`
}

type Order struct {
	bun.BaseModel `bun:"orders"`
	OrderId       int            `json:"order_id" bun:"order_id,pk"`
	CreatedAt     time.Time      `json:"created_at" bun:",nullzero,notnull,default:current_timestamp"`
	PhoneNumber   string         `json:"phone_number"`
	OrderDetails  []*OrderDetail `json:"order_details" bun:"rel:has-many,join:order_id=id"`
}

type OrderDetail struct {
	bun.BaseModel `bun:"order_details"`
	Id            int     `json:"id" bun:"id,pk"`
	OrderId       int     `json:"order_id"`
	ProductId     int     `json:"p_id"`
	Units         float32 `json:"units"`
	Profit        float32 `json:"profit"`
	Product       Product `json:"product" bun:"rel:has-one,join:id=p_id"`
}

func ArrayOfEmptyProducts() []Product {
	return []Product{}
}

func NewEmptyproduct() Product {
	return Product{
		Name:      "",
		BuyPrice:  2,
		SellPrice: 2,
		Units:     2,
		Type:      enum.ProductType.String(1),
	}
}

func InterfaceToProduct(v interface{}) Product {
	pVal := reflect.ValueOf(v)

	return Product{
		Pid:       int(pVal.FieldByName("Pid").Int()),
		Name:      pVal.FieldByName("Name").String(),
		BuyPrice:  float32(pVal.FieldByName("BuyPrice").Float()),
		SellPrice: float32(pVal.FieldByName("SellPrice").Float()),
		Units:     float32(pVal.FieldByName("Units").Float()),
		Type:      pVal.FieldByName("Type").String(),
	}
}
