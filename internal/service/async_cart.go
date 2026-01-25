package service

import (
	"context"
	"errors"

	"github.com/esmrzv/go_shop/internal/cartworker"
	"github.com/esmrzv/go_shop/internal/model"
)


type CartService interface {
	AddItem(ctx context.Context, userID, productID int) error
	GetCart(ctx context.Context, userID int) (*model.Cart, error)
}	

type ChanelCartService struct {
	worker *cartworker.Worker
}


func NewAsyncCartService(worker *cartworker.Worker) *ChanelCartService {
	return &ChanelCartService{worker: worker}
}

func (s *ChanelCartService) AddItem(ctx context.Context, userID, productID int) error {
	s.worker.Add(cartworker.AddItemCommand{
		UserID:    userID,
		ProductID: productID,
	})
	return nil
}

func (s *ChanelCartService) GetCart(ctx context.Context, userID int) (*model.Cart, error) {
	reply := make(chan *model.Cart)
	s.worker.Get(cartworker.GetCartCommand{
		UserID: userID,
		Reply:  reply,
	})
	cart := <-reply
	if cart == nil {
		return nil, errors.New("cart not found")
	}
	return cart, nil

}
