package repository
import (
	"database/sql"
	"context"
	"errors"
	"github.com/esmrzv/go_shop/internal/model"
)
type UserPostgres struct{
	db *sql.DB
}

func NewUserPostgres(db *sql.DB) *UserPostgres {
	return &UserPostgres{db: db}
}


func (r *UserPostgres) Create(ctx context.Context, email, passwordHash string) error {
	query := `INSERT INTO USERS (email, password_hash) VALUES ($1, $2)`
	_, err := r.db.ExecContext(ctx, query, email, passwordHash)
	return err
}


func (r *UserPostgres) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	query := `SELECT id, email, password_hash, created_at FROM USERS WHERE email = $1`
	row := r.db.QueryRowContext(ctx, query, email)

	var user model.User
	if err := row.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}