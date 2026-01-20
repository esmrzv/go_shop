package cartworker

import "github.com/esmrzv/go_shop/internal/model"

type Worker struct {
	addCh chan AddItemCommand
	getCh chan GetCartCommand
}


func NewWorker() *Worker {
	return &Worker{
		addCh: make(chan AddItemCommand),
		getCh: make(chan GetCartCommand),
	}
}

func (w *Worker) Run() {
	carts := make(map[int]*model.Cart)

	for {
		select {
		case cmd := <-w.addCh:
			cart, ok := carts[cmd.UserID]
			if !ok {
				cart = &model.Cart{UserID: cmd.UserID}
				carts[cmd.UserID] = cart
			}

			found := false
			for i := range cart.Items {
				if cart.Items[i].ProductID == cmd.ProductID {
					cart.Items[i].Quantity++
					found = true
					break
				}
			}

			if !found {
				cart.Items = append(cart.Items, model.CartItem{
					ProductID: cmd.ProductID,
					Quantity:  1,
				})
			}

		case cmd := <-w.getCh:
			cmd.Reply <- carts[cmd.UserID]
		}
	}
}


func (w *Worker) Add(cmd AddItemCommand) {
	w.addCh <- cmd
}

func (w *Worker) Get(cmd GetCartCommand) {
	w.getCh <- cmd
}