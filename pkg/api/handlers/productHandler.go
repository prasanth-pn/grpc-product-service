package handlers

import (
	services "github.com/prasanth-pn/grpc-product-service/pkg/usecase/interfaces"
)

type ProductHandler struct {
	productservice services.ProductUseCaseInterface
}

func NeWProductHandler(prodservice services.ProductUseCaseInterface) *ProductHandler {
	return &ProductHandler{productservice: prodservice}
}
