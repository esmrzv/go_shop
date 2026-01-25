package cartworker

import "github.com/esmrzv/go_shop/internal/model"

type AddItemCommand struct {
	UserID    int
	ProductID int
}

type GetCartCommand struct {
	UserID int
	Reply  chan *model.Cart
}
