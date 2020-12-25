package main

import (
	"fabrikapps/item_products_hamb_code/getProducts/pkg/config/database"
	"fabrikapps/item_products_hamb_code/getProducts/pkg/infrastrucure/delivery/handler"
	"fabrikapps/item_products_hamb_code/getProducts/pkg/service"

	"fabrikapps/item_products_hamb_code/getProducts/pkg/infrastrucure/storage/mysql"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	log.Println("iniciando grpc server")

	dbConn, err := database.GetConnection()
	if err != nil {
		fmt.Println(err)
	}
	defer dbConn.Close()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	rp := mysql.NewMysqlProductsRepository(dbConn)
	ss := service.NewService(rp)

	grpc_srv := grpc.NewServer()
	handler.NewProductListServerGrpc(grpc_srv, ss)

	log.Println("Listen on :50051")
	if err := grpc_srv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
