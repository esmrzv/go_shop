package service
import "github.com/esmrzv/go_shop/internal/cartworker"

type AsyncCartService struct {
	worker *cartworker.Worker
}


func NewAsyncCartService(worker *cartworker.Worker) *AsyncCartService {
	return &AsyncCartService{worker: worker}
}

func (s *AsyncCartService) AddItem(userID, productID int) {
	s.worker.Add(cartworker.AddItemCommand{
		UserID:    userID,
		ProductID: productID,
	})
}

func (s *AsyncCartService) GetCart(userID int) any {
	replyCh := make(chan any)
	s.worker.Get(cartworker.GetCartCommand{
		UserID: userID,
		Reply:  replyCh,
	})
	return <-replyCh
}		


