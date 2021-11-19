package repository

import (
	"context"
	"log"

	"github.com/asishshaji/thalir-backend/models"
	"github.com/uptrace/bun"
)

type ProductRepo struct {
	db *bun.DB
}

func NewProductRepo(db *bun.DB) PRInterface {
	return ProductRepo{db: db}
}

func (pR ProductRepo) CreateProduct(p interface{}) error {

	m := models.InterfaceToModel(p)

	_, err := pR.db.NewInsert().Model(&m).Exec(context.Background())
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (pR ProductRepo) GetAllProducts() (interface{}, error) {
	ms := models.ArrayOfEmptyProducts()
	err := pR.db.NewSelect().Model(&ms).Limit(20).Scan(context.TODO())
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return ms, nil
}
