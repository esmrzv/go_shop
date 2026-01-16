package repository

import (
	"context"

	"github.com/esmrzv/go_shop/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, email, passwordHash string) error
	GetByEmail(ctx context.Context, email string) (*model.User, error)

}