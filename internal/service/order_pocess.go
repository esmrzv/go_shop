package service

import (
	"context"
	"errors"

	"github.com/esmrzv/go_shop/internal/cartworker"
	"github.com/esmrzv/go_shop/internal/model"
	"github.com/esmrzv/go_shop/internal/orderworker"
)

func (s *OrderService) Process(
	ctx context.Context,
	cmd orderworker.OrderCommand,
) error {
	// 1. Получаем корзину
	reply := make(chan *model.Cart)
	s.cartWorker.Get(cartworker.GetCartCommand{
		UserID: cmd.UserID,
		Reply:  reply,
	})

	cart := <-reply
	if cart == nil || len(cart.Items) == 0 {
		return errors.New("cart is empty")
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// 2. Создаём order
	var orderID int
	err = tx.QueryRowContext(ctx,
		`INSERT INTO orders (user_id) VALUES ($1) RETURNING id`,
		cmd.UserID,
	).Scan(&orderID)
	if err != nil {
		return err
	}

	// 3. order_items
	for _, item := range cart.Items {
		_, err = tx.ExecContext(ctx,
			`INSERT INTO order_items (order_id, product_id, quantity)
			 VALUES ($1, $2, $3)`,
			orderID, item.ProductID, item.Quantity,
		)
		if err != nil {
			return err
		}
	}

	// 4. commit
	if err := tx.Commit(); err != nil {
		return err
	}

	// 5. очистка корзины (отдельная команда)
	//s.cartWorker.Clear(cmd.UserID)

	return nil
}
