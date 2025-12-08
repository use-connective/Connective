package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/x-sushant-x/connective/internal/core/domain"
	"github.com/x-sushant-x/connective/internal/core/util"
)

type AuthRepo struct {
	db *DB
}

func NewAuthRepo(db *DB) *AuthRepo {
	return &AuthRepo{
		db: db,
	}
}

func (ar *AuthRepo) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING *`

	var newUser domain.User

	err := ar.db.QueryRow(ctx, query, user.Name, user.Email, user.Password).Scan(&newUser.ID, &newUser.Name, &newUser.Email, &newUser.Password, &newUser.CreatedAt, &newUser.UpdatedAt, &newUser.IsOnboardingCompleted, &newUser.State)
	if err != nil {
		return nil, err
	}

	return util.ToUserResponse(&newUser), nil
}

// GetUser IMPORTANT - This function will also return hashed user password in response.
// Use it carefully. NEVER expose it's response to frontend client.
func (ar *AuthRepo) GetUser(ctx context.Context, email string) (*domain.User, error) {
	query := `SELECT * FROM users WHERE email = $1`

	var user domain.User

	err := ar.db.QueryRow(ctx, query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.IsOnboardingCompleted, &user.State)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (ar *AuthRepo) GetUserById(ctx context.Context, userID int) (*domain.User, error) {
	query := `SELECT * FROM users WHERE id = $1`

	var user domain.User

	err := ar.db.QueryRow(ctx, query, userID).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.IsOnboardingCompleted, &user.State)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return util.ToUserResponse(&user), nil
}
