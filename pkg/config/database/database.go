package database

import (
	"database/sql"
	"log"

	//	"sync"

	//log "github.com/sirupsen/logrus"
	"fabrikapps/item_products_hamb_code/getProducts/pkg/errors"

	_ "github.com/go-sql-driver/mysql"
)

var (
	//	once               sync.Once
	databaseConnection *sql.DB
)

func GetConnection() (*sql.DB, error) {

	//once.Do(func() {
	var err error
	connectionString := "root:yourpass@tcp(localhost:3306)/databaseName"
	databaseConnection, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatalf("Can not connect to database: %v", err)
		err = errors.ErrSQLConnect
		return nil, err

	}
	err = databaseConnection.Ping()
	if err != nil {
		log.Fatalf("Database error: %v", err)
		err = errors.ErrSQLConnect
		return nil, err
	}
	//})

	return databaseConnection, err

}
