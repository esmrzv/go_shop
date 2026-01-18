package repository

import (
	"context"

	"github.com/esmrzv/go_shop/internal/model"
)

type ProductRepository interface {
	Create(ctx context.Context, name string, price float64) error
	List(ctx context.Context)([]model.Product, error)
}