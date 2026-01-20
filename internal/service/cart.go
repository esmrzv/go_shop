package service

import (
	"context"
	"sync"

	"github.com/esmrzv/go_shop/internal/model"
)

type CartService struct {
	mu sync.Mutex
	carts map[int]*model.Cart
}


func NewCartService() *CartService {
	return &CartService{
		carts: make(map[int]*model.Cart, 0),
	}
}


func (s *CartService) AddItem(ctx context.Context, user_id, product_id int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	cart, ok := s.carts[user_id]
	if !ok {
		cart = &model.Cart{
			UserID: user_id,
		}
		s.carts[user_id] = cart
	}

	for i := range cart.Items {
		if cart.Items[i].ProductID == product_id {
			cart.Items[i].Quantity++
			return
		}
	}

	cart.Items = append(cart.Items, model.CartItem{
		ProductID: product_id,
		Quantity:  1,
	})
}


func (s *CartService) GetCart(ctx context.Context, user_id int) *model.Cart {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.carts[user_id]
}