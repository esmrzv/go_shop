package service

import (
	"context"

	"github.com/esmrzv/go_shop/internal/model"
	"github.com/esmrzv/go_shop/internal/repository"
)

type ProductService struct {
	repo repository.ProductRepository
}

func NewProductService(r repository.ProductRepository) *ProductService{
	return &ProductService{
		repo: r,
	}
}

func (s *ProductService) Create(ctx context.Context, name string, price float64) error{
	return s.repo.Create(ctx, name, price)
}


func (s *ProductService) List(ctx context.Context) ([]model.Product, error){
	return s.repo.List(ctx)
}