package repository

import (
	"context"
	"errors"
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

func (pR ProductRepo) CreateProduct(p interface{}) (interface{}, error) {

	m := models.InterfaceToModel(p)

	_, err := pR.db.NewInsert().Model(&m).Exec(context.Background())
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return m, nil
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

func (pR ProductRepo) UpdateProduct(p interface{}) error {
	m := models.InterfaceToModel(p)
	_, err := pR.db.NewUpdate().Model(&m).Column("name", "units", "buy_price", "sell_price", "type").Where("p_id = ?", m.Pid).Exec(context.Background())
	if err != nil {
		log.Fatalln(err)
		return err
	}
	return nil
}

func (pR ProductRepo) DeleteProduct(pid int) error {
	res, err := pR.db.NewDelete().Table("products").Where("p_id = ?", pid).Exec(context.Background())
	if err != nil {
		return err
	}

	rowId, err := res.RowsAffected()
	if rowId == 0 {
		return errors.New("no resource found")
	}
	if err != nil {
		return err
	}

	return nil
}

func (pR ProductRepo) GetProduct(pid int) (interface{}, error) {
	mp := models.NewEmptyProductCreationRequest()
	err := pR.db.NewSelect().Model(&mp).Where("p_id = ?", pid).Scan(context.Background())
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return mp, nil

}
