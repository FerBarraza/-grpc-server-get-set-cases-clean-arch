package product

type Item struct {
	Category    string  `db:"CategoriasMenuNombre"`
	ItemId      uint32  `db:"TipoMenuItem_id"`
	CategoryId  uint32  `db:"TipoItem_categoriasMenuId"`
	NombreItem  string  `db:"TipoItem_NombreItem"`
	ImageName   string  `db:"TipoItem_ImageName"`
	Descripcion string  `db:"TipoItem_Descripcion"`
	ItemPrice   float32 `db:"TipoItem_Price"`
}

type ProductList struct {
	ListOfItems []*Item
}
