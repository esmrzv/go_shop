package service

import (
	"context"

	"github.com/esmrzv/go_shop/internal/model"
	"github.com/esmrzv/go_shop/internal/repository"
)

type CategoryService interface {
	Create(ctx context.Context, name string) (int, error)
	List(ctx context.Context) ([]model.Category, error)
}	

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(r repository.CategoryRepository) CategoryService{
	return &categoryService{repo: r}
}


func (s *categoryService) Create(ctx context.Context, name string) (int, error){
	return s.repo.Create(ctx, name)
}


func (s *categoryService) List(ctx context.Context) ([]model.Category, error){
	return s.repo.List(ctx)
}
