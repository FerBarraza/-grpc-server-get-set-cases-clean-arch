package handler

import (
	"context"
	"fabrikapps/item_products_hamb_code/getProducts/pkg/errors"
	"log"
	"time"

	"google.golang.org/grpc"

	product_model "fabrikapps/item_products_hamb_code/getProducts/pkg/domain/product"
	"fabrikapps/item_products_hamb_code/getProducts/pkg/infrastrucure/delivery/grpc/itemProductspb"
	app_get_service "fabrikapps/item_products_hamb_code/getProducts/pkg/service"
)

type service struct {
	serv app_get_service.Service
}

// NewProductListServerGrpc ...
func NewProductListServerGrpc(gserver *grpc.Server, getserv app_get_service.Service) {

	asdfafd := &service{
		serv: getserv,
	}

	itemProductspb.RegisterProductoStockServer(gserver, asdfafd)
}

//GetProductsList ....
func (s *service) GetProductsList(ctx context.Context, req *itemProductspb.CategoryIdRequest) (*itemProductspb.ItemListResponse, error) {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	reqCatId := req.ReqCategoryId
	var prList []*product_model.Item
	prList, err := s.serv.GetProductListCatService(reqCatId)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	if prList == nil {
		return nil, errors.ErrNoListItems
	}

	var products []*itemProductspb.Item

	for _, element := range prList {
		product := &itemProductspb.Item{
			Category:    element.Category,
			ItemId:      element.ItemId,
			CategoryId:  element.CategoryId,
			NombreItem:  element.NombreItem,
			ImageName:   element.ImageName,
			Descripcion: element.Descripcion,
			ItemPrice:   element.ItemPrice,
		}

		products = append(products, product)
	}

	if products == nil {
		return nil, errors.ErrNoListItemsTemp
	}

	response := &itemProductspb.ItemListResponse{
		ListaItems: products,
	}
	return response, nil
}

func (s *service) SetProduct(ctx context.Context, reqItemToRepl *itemProductspb.Item) (*itemProductspb.ItemStatusResponse, error) {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	reqItem := reqItemToRepl
	//var prItemToRepl *product_model.Item

	prItemToRepl := &product_model.Item{
		Category:    reqItem.Category,
		ItemId:      reqItem.ItemId,
		CategoryId:  reqItem.CategoryId,
		NombreItem:  reqItem.NombreItem,
		ImageName:   reqItem.ImageName,
		Descripcion: reqItem.Descripcion,
		ItemPrice:   reqItem.ItemPrice,
	}

	statusSuccessfulOrNo, err := s.serv.SetProductService(prItemToRepl)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	if statusSuccessfulOrNo == false {
		return nil, errors.ErrNoUpdatedItem
	}

	response := &itemProductspb.ItemStatusResponse{
		StatusSuccessful: statusSuccessfulOrNo,
	}
	return response, nil
}
