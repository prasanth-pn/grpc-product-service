package usecase 


import( 
	services "github.com/prasanth-pn/grpc-product-service/pkg/usecase/interfaces"
)


type ProductUseCase struct{
Repo services.ProductRepository
}
func NewProductUseCase(repo services.ProductRepository)*ProductUseCase{
	return &ProductUseCase{
		Repo :repo,
	}
}