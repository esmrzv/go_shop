package service

import (
	"context"

	"github.com/esmrzv/go_shop/internal/model"
	"github.com/esmrzv/go_shop/internal/repository"
)

type ProductService interface {
	Create(ctx context.Context, name string, price float64, category_id *int) error
	List(ctx context.Context) ([]model.Product, error)
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(r repository.ProductRepository) ProductService {
	return &productService{repo: r}
}

func (s *productService) Create(ctx context.Context, name string, price float64, category_id *int) error {
	return s.repo.Create(ctx, name, price, category_id)
}

func (s *productService) List(ctx context.Context) ([]model.Product, error) {
	return s.repo.List(ctx)
}