package errors

import (
	"errors"
	//"fabrikapps/item_products_hamb_code/asdf/pkg/infrastrucure/storage/mysql/repository.go"
)

var (
	// ErrSQLStatement returned if the query was incorrect
	ErrSQLStatement = errors.New("Could not prepare statement")
	// ErrSQLConnect ...
	ErrSQLConnect = errors.New("Could not connect to the database")
	// ErrSQLInsert returned if the row could not insert into db
	ErrSQLInsert = errors.New("Could not read row")
	// ErrNoListItems ...
	ErrNoListItems = errors.New("Could not read the id requested or no exists")
	// ErrNoListItemsTemp ...
	ErrNoListItemsTemp = errors.New("There not items in temporal products list")
	// ErrNoUpdatedItem ...
	ErrNoUpdatedItem = errors.New("The item is not updated")
)
