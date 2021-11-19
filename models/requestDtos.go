package models

import (
	"reflect"

	"github.com/asishshaji/thalir-backend/enum"
	"github.com/uptrace/bun"
)

type productCreationRequest struct {
	bun.BaseModel `bun:"products"`
	Pid           int     `json:"p_id" bun:"p_id,pk"`
	Name          string  `json:"name"`
	BuyPrice      float32 `json:"buy_price"`
	SellPrice     float32 `json:"sell_price"`
	Units         float32 `json:"units"`
	Type          string  `json:"type"`
}

func ArrayOfEmptyProducts() []productCreationRequest {
	return []productCreationRequest{}
}

func NewEmptyProductCreationRequest() productCreationRequest {
	return productCreationRequest{
		Name:      "",
		BuyPrice:  2,
		SellPrice: 2,
		Units:     2,
		Type:      enum.ProductType.String(1),
	}
}

func InterfaceToModel(v interface{}) productCreationRequest {
	pVal := reflect.ValueOf(v)

	return productCreationRequest{
		Pid:       int(pVal.FieldByName("Pid").Int()),
		Name:      pVal.FieldByName("Name").String(),
		BuyPrice:  float32(pVal.FieldByName("BuyPrice").Float()),
		SellPrice: float32(pVal.FieldByName("SellPrice").Float()),
		Units:     float32(pVal.FieldByName("Units").Float()),
		Type:      pVal.FieldByName("Type").String(),
	}
}
