package product

// Repository provides access to role repository
type Repository interface {
	GetProductListCat(reqCatId uint32) ([]*Item, error)
	SetProductRepo(reqItemToRepl *Item) (bool, error)
}
