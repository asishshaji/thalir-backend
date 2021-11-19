package models

import (
	"reflect"
	"time"

	"github.com/asishshaji/thalir-backend/enum"
	"github.com/uptrace/bun"
)

type Product struct {
	bun.BaseModel `bun:"products"`
	Pid           int     `json:"p_id" bun:"p_id,pk"`
	Name          string  `json:"name"`
	BuyPrice      float32 `json:"buy_price"`
	SellPrice     float32 `json:"sell_price"`
	Units         float32 `json:"units"`
	Type          string  `json:"type"`
}

type Order struct {
	bun.BaseModel `bun:"orders"`
	OrderId       int             `json:"order_id" bun:"order_id,pk"`
	CreatedAt     time.Time       `json:"created_at" bun:",nullzero,notnull,default:current_timestamp"`
	OrderHistory  []*OrderHistory `json:"order_history" bun:"rel:has-many,join:order_id=id"`
}

type OrderHistory struct {
	bun.BaseModel `bun:"order_history"`
	Id            int     `json:"id" bun:"id,pk"`
	Profit        float32 `json:"profit"`
	ProductId     int     `json:"p_id"`
	OrderId       int     `json:"order_id"`
	Units         float32 `json:"units"`
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
