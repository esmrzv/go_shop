package model

type CartItem struct {
	ProductID int
	Quantity  int
}


type Cart struct {
	UserID int
	Items []CartItem
}