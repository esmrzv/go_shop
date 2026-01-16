package service

import (
	"context"
	"errors"
	"time"

	"github.com/esmrzv/go_shop/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"github.com/esmrzv/go_shop/internal/auth"
)
type AuthService struct {
	repo repository.UserRepository
	jwtSectet string
}

func NewAuthService(repo repository.UserRepository, secret string) *AuthService {
	return &AuthService{repo: repo, jwtSectet: secret}
}

func (s *AuthService) Register(ctx context.Context, email, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return s.repo.Create(ctx, email, string(hash))
}

func (s *AuthService) Login(ctx context.Context, email, password string) (string, error){
	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil{
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(password),
		
	); err != nil{
		return "",  errors.New("invalid credentions")
	}

	return auth.GenerateJWT(user.ID, s.jwtSectet, 24*time.Hour)
}