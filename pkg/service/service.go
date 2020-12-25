package service

import (
	"fabrikapps/item_products_hamb_code/getProducts/pkg/domain/product"
	"log"
)

type Service interface {
	GetProductListCatService(reqCatId uint32) ([]*product.Item, error)
	SetProductService(reqItemToRepl *product.Item) (bool, error)
}

type service struct {
	getRepo product.Repository
}

func NewService(p product.Repository) Service {
	return &service{p}
}

func (s *service) GetProductListCatService(reqCatId uint32) ([]*product.Item, error) {
	products, err := s.getRepo.GetProductListCat(reqCatId)
	if err != nil {
		log.Printf("Error %s when calling the getProductListCat\n", err)
		return nil, nil
	}
	return products, nil
}
func (s *service) SetProductService(reqItemToRepl *product.Item) (bool, error) {
	valueSuccessful, err := s.getRepo.SetProductRepo(reqItemToRepl)
	if err != nil {
		log.Printf("Error %s when calling the setProduct\n", err)
		return valueSuccessful, err
	}
	return valueSuccessful, nil
}
