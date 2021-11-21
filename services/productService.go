package services

import (
	"log"

	"github.com/asishshaji/thalir-backend/models"
	repository "github.com/asishshaji/thalir-backend/repositories"
)

type ProductService struct {
	pRepo repository.PRInterface
}

func NewProductService(pRepo repository.PRInterface) PSInterface {
	return ProductService{pRepo: pRepo}
}

func (pS ProductService) CreateProduct(product models.Product) (interface{}, error) {

	p, err := pS.pRepo.CreateProduct(product)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (pS ProductService) GetAllProducts() (interface{}, error) {
	d, err := pS.pRepo.GetAllProducts()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return d, nil
}

func (pS ProductService) UpdateProduct(product models.Product) error {
	err := pS.pRepo.UpdateProduct(product)
	if err != nil {
		return err
	}

	return nil
}

func (pS ProductService) DeleteProduct(pid int) error {
	err := pS.pRepo.DeleteProduct(pid)
	if err != nil {
		return err
	}

	return nil
}

func (pS ProductService) GetProduct(pid int) (models.Product, error) {
	return pS.pRepo.GetProduct(pid)

}
