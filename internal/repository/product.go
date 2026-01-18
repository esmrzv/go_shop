package repository

import (
	"context"

	"github.com/esmrzv/go_shop/internal/model"
)

type Product interface {
	Crate(ctx context.Context, name string, price float64) error
	List(ctx context.Context)([]model.Product, error)
}