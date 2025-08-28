package catalog

import (
	"context"
	"fmt"

	"github.com/segmentio/ksuid"
)

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

func (s *catalogService) PostProduct(ctx context.Context, name, description string, price float64) (*Product, error) {
	p := &Product{
		ID:          ksuid.New().String(),
		Name:        name,
		Description: description,
		Price:       fmt.Sprintf("%v", price),
	}

	if err := s.repository.PutProduct(ctx, *p); err != nil {
		return nil, err
	}

	return p, nil
}

func (s *catalogService) GetProduct(ctx context.Context, id string) (*Product, error) {
	return s.repository.GetProductByID(ctx, id)
}

func (s *catalogService) GetProducts(ctx context.Context, skip, take uint64) ([]Product, error) {
	if take > 100 || (skip == 0 && take == 0) {
		take = 100
	}

	return s.repository.ListProducts(ctx, skip, take)
}

func (s *catalogService) GetProductsByIDs(ctx context.Context, ids []string) ([]Product, error) {
	return s.repository.ListProductsWithIDs(ctx, ids)
}

func (s *catalogService) SearchProducts(ctx context.Context, query string, skip, take uint64) ([]Product, error) {
	if take > 100 || (skip == 0 && take == 0) {
		take = 100
	}

	return s.repository.SearchProducts(ctx, query, skip, take)
}
