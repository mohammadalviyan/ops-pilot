package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrUserNotFound = errors.New("user not found")

type Repository interface {
	FindByEmail(ctx context.Context, email string) (User, error)
	FindByID(ctx context.Context, id string) (User, error)
}

type PostgresRepository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) FindByEmail(ctx context.Context, email string) (User, error) {
	return r.findOne(ctx, "users.email = $1", email)
}

func (r *PostgresRepository) FindByID(ctx context.Context, id string) (User, error) {
	return r.findOne(ctx, "users.id = $1", id)
}

func (r *PostgresRepository) findOne(ctx context.Context, condition string, value string) (User, error) {
	query := fmt.Sprintf(`
		SELECT users.id::text, users.name, users.email, users.password_hash, users.status, roles.code
		FROM users
		JOIN roles ON roles.id = users.role_id
		WHERE %s AND users.deleted_at IS NULL
		LIMIT 1
	`, condition)

	var user User
	err := r.db.QueryRow(ctx, query, value).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.Status,
		&user.RoleCode,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return User{}, ErrUserNotFound
	}
	if err != nil {
		return User{}, fmt.Errorf("find user: %w", err)
	}

	return user, nil
}
