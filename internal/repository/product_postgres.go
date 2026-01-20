package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/esmrzv/go_shop/internal/model"
)

type productPostgres struct {
	db *sql.DB
}

func NewProductPostgres(db *sql.DB) ProductRepository {
	return &productPostgres{db: db}
}

func (p *productPostgres) Create(ctx context.Context, name string, price float64, category_id *int) error {
	var dbName string
	err := p.db.QueryRow("SELECT current_database()").Scan(&dbName)
	if err != nil {
		log.Println("db name error:", err)
	} else {
		log.Println("CONNECTED DB:", dbName)
	}
	query := `
		INSERT INTO products (name, price, category_id)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	var id int
	if err := p.db.QueryRowContext(ctx, query, name, price, category_id).Scan(&id); err != nil {
		return err
	}

	log.Println("product inserted:", id)
	return nil
}


func (p *productPostgres) List(ctx context.Context) ([]model.Product, error) {
	var dbName string
	_ = p.db.QueryRow("SELECT current_database()").Scan(&dbName)
	log.Println("CONNECTED DB:", dbName)

	query := `SELECT id, name, price, category_id, created_at FROM products`

	rows, err := p.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]model.Product, 0)

	for rows.Next() {
		var product model.Product
		if err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.Category_id,
			&product.CreatedAt,
			
		); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}