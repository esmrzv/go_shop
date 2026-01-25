package service
import "github.com/esmrzv/go_shop/internal/cartworker"


type CartService interface {
	AddItem(userID, productID int)
	GetCart(userID int) any
}	

type ChanelCartService struct {
	worker *cartworker.Worker
}


func NewAsyncCartService(worker *cartworker.Worker) *ChanelCartService {
	return &ChanelCartService{worker: worker}
}

func (s *ChanelCartService) AddItem(userID, productID int) {
	s.worker.Add(cartworker.AddItemCommand{
		UserID:    userID,
		ProductID: productID,
	})
}

func (s *ChanelCartService) GetCart(userID int) any {
	replyCh := make(chan any)
	s.worker.Get(cartworker.GetCartCommand{
		UserID: userID,
		Reply:  replyCh,
	})
	return <-replyCh
}		


