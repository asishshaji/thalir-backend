package repository

import (
	"log"

	"github.com/asishshaji/thalir-backend/models"
	"gorm.io/gorm"
)

type ProductRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) PRInterface {

	return ProductRepo{db: db}
}

func (pR ProductRepo) CreateProduct(p models.Product) (interface{}, error) {

	err := pR.db.Create(&p).Error

	if err != nil {
		log.Println(err)
		return nil, err
	}

	pBack := p.ToBackup()
	err = pR.db.Create(&pBack).Error

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return p, nil
}

func (pR ProductRepo) GetAllProducts() (interface{}, error) {
	ms := []models.Product{}
	err := pR.db.Model(&models.Product{}).Find(&ms).Error
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return ms, nil
}

func (pR ProductRepo) UpdateProduct(p models.Product) error {
	err := pR.db.Debug().Model(&models.Product{}).Where("id = ?", p.ID).Updates(map[string]interface{}{
		"id":         p.ID,
		"name":       p.Name,
		"type":       p.Type,
		"buy_price":  p.BuyPrice,
		"sell_price": p.SellPrice,
		"units":      p.Units,
	}).Error
	if err != nil {
		log.Fatalln(err)
		return err
	}

	// pBack := p.ToBackup()
	// err = pR.db.Create(&pBack).Error

	// if err != nil {
	// 	log.Println(err)
	// 	return err
	// }
	return nil
}

func (pR ProductRepo) DeleteProduct(pid int) error {
	err := pR.db.Model(&models.Product{}).Where("id = ?", pid).Take(&models.Product{}).Delete(&models.Product{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (pR ProductRepo) GetProduct(pid int) (models.Product, error) {
	mp := models.Product{}
	err := pR.db.Model(&models.Product{}).Where("id = ?", pid).Take(&mp).Error
	if err != nil {
		log.Println(err)
		return models.Product{}, err
	}
	return mp, nil

}
