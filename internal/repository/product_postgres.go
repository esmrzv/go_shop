package repository

import (
	"context"
	"database/sql"

	"github.com/esmrzv/go_shop/internal/model"
)

type ProductPostgres struct{
	db *sql.DB
}

func NewProductPostgres(db *sql.DB) *ProductPostgres{
	return &ProductPostgres{db: db}
}

func (p *ProductPostgres) Create(ctx context.Context, name string, price float64) error {
	query := `INSERT INTO PRODUCTS (name, price) VALUES ($1, $2)`
	_, err :=p.db.ExecContext(ctx, query, name, price)
	return err
}


func (p *ProductPostgres) List(ctx context.Context) ([]model.Product, error){
	query := `SELECT * FROM products`
	rows, err := p.db.QueryContext(ctx, query)
	if err != nil{
		return nil, err
	}
	rows.Close()

	var products []model.Product

	for rows.Next() {
		var product model.Product
		rows.Scan(&product.ID, product.Name, product.Price, product.CreatedAt)
		products = append(products, product)
	}
	return products, nil
}