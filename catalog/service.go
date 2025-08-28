package catalog

import "context"

type Service interface {
	PostProduct(ctx context.Context, name, description string, price float64) (*Product, error)
	GetProduct(ctx context.Context, id string) (*Product, error)
	GetProducts(ctx context.Context, skip, take uint64) ([]Product, error)
	GetProductsByIDs(ctx context.Context, ids []string) ([]Product, error)
	SearchProducts(ctx context.Context, query string, skip, take uint64) ([]Product, error)
}

type Product struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       string `json:"price"`
}

type catalogService struct {
	repository Repository
}

func NewCatalogService(r Repository) Service {
	return &catalogService{repository: r}
}

func (s *catalogService) PostProduct(ctx context.Context, name, description string, price float64) (*Product, error){

}

func (s *catalogService) GetProduct(ctx context.Context, id string) (*Product, error){

}

func (s *catalogService) GetProducts(ctx context.Context, skip, take uint64) ([]Product, error) {

}

func (s *catalogService) GetProductsByIDs(ctx context.Context, ids []string) ([]Product, error) {

}

func (s *catalogService) SearchProducts(ctx context.Context, query string, skip, take uint64) ([]Product, error){
	
}
