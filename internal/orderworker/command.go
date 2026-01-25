package orderworker

type OrderCommand struct {
	UserID int
	Reply  chan error
}	