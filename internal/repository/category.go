package repository

import (
	"context"
	"database/sql"

	"github.com/esmrzv/go_shop/internal/model"
)

type CategoryRepository interface {
    Create(ctx context.Context, name string) (int, error)
    List(ctx context.Context) ([]model.Category, error)
}

type categoryPostgres struct {
    db *sql.DB
}

func NewCategoryRepo(db *sql.DB) CategoryRepository {
    return &categoryPostgres{db: db}
}


func (r *categoryPostgres) Create(ctx context.Context, name string) (int, error) {
	var id int
	query := `INSERT INTO categories (name) VALUES ($1) RETURNING id`
	err := r.db.QueryRowContext(ctx, query, name).Scan(&id)
	if err != nil{
		return 0, err
	}
	return id, nil
}

func (r *categoryPostgres) List(ctx context.Context) ([]model.Category, error){
	query := `SELECT id, name from categories`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil{
		return nil, err
	}
	defer rows.Close()

	categories := make([]model.Category, 0)
	for rows.Next(){
		var category model.Category
		if err := rows.Scan(&category.ID, &category.Name); err != nil{
			return nil, err
		}
		categories = append(categories, category)

	}
	return categories, nil
}