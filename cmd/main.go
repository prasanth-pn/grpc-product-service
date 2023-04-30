package main

import (
	"fmt"
	"log"
	"net"

	"github.com/prasanth-pn/grpc-product-service/pkg/api/handlers"
	"github.com/prasanth-pn/grpc-product-service/pkg/config"
	"github.com/prasanth-pn/grpc-product-service/pkg/db"
	"github.com/prasanth-pn/grpc-product-service/pkg/pb"
	"github.com/prasanth-pn/grpc-product-service/pkg/repository"
	"github.com/prasanth-pn/grpc-product-service/pkg/usecase"
	"google.golang.org/grpc"
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

	if err != nil {
		log.Fatalf("failed to listening:", err)
	}
	sqlDB := db.ConnectDB(c)
	ProductRepository := repository.NewAuthRepository(sqlDB)
	productUsecase := usecase.NewProductUseCase(ProductRepository)
	s := handlers.NeWProductHandler(productUsecase)
	grpcServer := grpc.NewServer()

	pb.RegisterProductServiceServer(grpcServer, s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start the server %s", err)
	}

}
