package cartworker

type AddItemCommand struct {
	UserID int
	ProductID int
}

type GetCartCommand struct {
	UserID int
	Reply chan any
}

