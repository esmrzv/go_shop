package model
import "time"

type Order struct {
	ID    int        `json:"id"`
	UserID int        `json:"user_id"`
	Total float64        `json:"total"`
	CreatedAt time.Time    `json:"created_at"`
}


type OrderItem struct {
	OrderID   int       `json:"order_id"`
	ProductID int       `json:"product_id"`
	Price     float64   `json:"price"`
	Quantity  int       `json:"quantity"`
	
}