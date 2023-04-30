package usecase 


import( 
	services "github.com/prasanth-pn/grpc-product-service/pkg/usecase/interfaces"
	 "github.com/prasanth-pn/grpc-product-service/pkg/repository/interfaces"

)


type ProductUseCase struct{
Repo interfaces.ProductRepository
}
func NewProductUseCase(repo interfaces.ProductRepository)services.ProductUseCaseInterface{
	return &ProductUseCase{
		Repo :repo,
	}
}