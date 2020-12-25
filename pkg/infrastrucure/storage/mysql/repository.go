package mysql

import (
	"database/sql"
	"fabrikapps/item_products_hamb_code/getProducts/pkg/domain/product"
	"log"
)

const (
	mysqlGetProductListCat = "SELECT CategoriasMenuNombre, TipoMenuItem_id, TipoItem_categoriasMenuId, TipoItem_NombreItem, TipoItem_ImageName, TipoItem_Descripcion, TipoItem_Price FROM cerveceriahamburgueseriaSistema.TipoItem INNER JOIN cerveceriahamburgueseriaSistema.CategoriasMenu ON TipoItem_categoriasMenuId=CategoriasMenu_id WHERE TipoItem_categoriasMenuId=?;"
	mysqlSetProductQuery   = "UPDATE cerveceriahamburgueseriaSistema.TipoItem SET TipoItem_NombreItem=?, TipoItem_Descripcion=?, TipoItem_Price=? WHERE TipoMenuItem_id=?;"
)

// Storage keeps data in db
type mysqlProductsRepository struct {
	db *sql.DB
}

func NewMysqlProductsRepository(db *sql.DB) product.Repository {
	return &mysqlProductsRepository{db}
}

// GetProductListCat get the given Product list to the repository
func (reposincampos *mysqlProductsRepository) GetProductListCat(reqCatId uint32) ([]*product.Item, error) {

	res, err := reposincampos.db.Query(mysqlGetProductListCat, reqCatId)
	if err != nil {
		log.Printf("Error %s when creating DB\n", err)
		return nil, err
	}

	var products []*product.Item

	for res.Next() {
		product := &product.Item{}

		_ = res.Scan(&product.Category, &product.ItemId, &product.CategoryId, &product.NombreItem, &product.ImageName,
			&product.Descripcion, &product.ItemPrice)

		products = append(products, product)
	}

	return products, nil
}

func (repoR *mysqlProductsRepository) SetProductRepo(reqItemToRepl *product.Item) (bool, error) {
	_, err := repoR.db.Exec(mysqlSetProductQuery, reqItemToRepl.NombreItem, reqItemToRepl.Descripcion, reqItemToRepl.ItemPrice, reqItemToRepl.ItemId)
	if err != nil {
		log.Printf("Error %s when creating DB\n", err)
		return false, err
	}
	return true, nil
}
