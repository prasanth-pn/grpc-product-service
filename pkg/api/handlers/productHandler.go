package handlers

import (
	"context"
	"fmt"

	"github.com/prasanth-pn/grpc-product-service/pkg/pb"
	services "github.com/prasanth-pn/grpc-product-service/pkg/usecase/interfaces"
)

type ProductHandler struct {
	productservice services.ProductUseCaseInterface
}

func NeWProductHandler(prodservice services.ProductUseCaseInterface) *ProductHandler {
	return &ProductHandler{productservice: prodservice}
}

func (cr *ProductHandler) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	fmt.Println("handler is initiated")

	return &pb.CreateProductResponse{
		Name: "prasanth",
		Id: 3434,
		Error: "ready",
	},nil
}
