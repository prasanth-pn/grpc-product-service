package main

import (
	"fmt"
	"log"
	"net"

	"github.com/prasanth-pn/grpc-product-service/pkg/config"
	"github.com/prasanth-pn/grpc-product-service/pkg/db"
	"github.com/prasanth-pn/grpc-product-service/pkg/repository"
	"github.com/prasanth-pn/grpc-product-service/pkg/usecase"
)

func main() {
	c, err := config.LoadConfig()
	fmt.Println(c, "port")
	if err != nil {
		log.Fatalln("failed at config", err)
	}
	_, err = db.ConnectGorm(c)
	if err != nil {
		log.Fatalf("error connecting gorm : %s", err)
	}
	lis, err := net.Listen("tcp", c.DBPort)
	if err!=nil{
		log.Fatalf("failed to listening:",err)
	}
	sqlDB:=db.ConnectDB(c)
	ProductRepository:=repository.NewAuthRepository(sqlDB)
	productUsecase:=usecase.NewProductUseCase(ProductRepository)
	handler.

}
