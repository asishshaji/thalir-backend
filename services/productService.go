package services

import (
	"log"

	repository "github.com/asishshaji/thalir-backend/repositories"
)

type ProductService struct {
	pRepo repository.PRInterface
}

func NewProductService(pRepo repository.PRInterface) PSInterface {
	return ProductService{pRepo: pRepo}
}

func (pS ProductService) CreateProduct(product interface{}) error {

	err := pS.pRepo.CreateProduct(product)
	if err != nil {
		return err
	}
	return nil
}

func (pS ProductService) GetAllProducts() (interface{}, error) {
	d, err := pS.pRepo.GetAllProducts()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return d, nil
}
