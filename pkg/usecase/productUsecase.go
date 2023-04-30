package usecase 


type ProductUseCase struct{
Repo interfaces.ProductRepository
}
func NewProductUseCase(repo interfaces.ProductRepository)service.ProductUseCase{
	return &ProProductUseCase{
		Repo :repo,
	}
}