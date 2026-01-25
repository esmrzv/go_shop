package service

import (
	"database/sql"
	"context"

	"github.com/esmrzv/go_shop/internal/cartworker"
	"github.com/esmrzv/go_shop/internal/orderworker"
)

type OrderService struct {
	db *sql.DB
	cartWorker *cartworker.Worker
	worker *orderworker.Worker
}

func NewOrderService(db *sql.DB, cartWorker *cartworker.Worker, w *orderworker.Worker) *OrderService {
	return &OrderService{
		db: db,
		cartWorker: cartWorker,
		worker: w,
	}
}

func (s *OrderService) Create (ctx context.Context, userID int) error {
	reply := make(chan error)

	s.worker.Enqueue(orderworker.OrderCommand{
		UserID: userID,
		Reply:  reply,
	})
	return <-reply
}