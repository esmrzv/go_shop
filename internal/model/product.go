package model

import "time"

type Product struct {
	ID        int   
	Name      string
	Price     float64 
	CreatedAt time.Time 
}