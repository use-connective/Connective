package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/x-sushant-x/connective/internal/core/domain"
)

type UserRepo struct {
	db *DB
}

func NewUserRepo(db *DB) *UserRepo {
	return &UserRepo{db}
}

func (r *UserRepo) Create(ctx context.Context, u *domain.User) (*domain.User, error) {
	query := `
		INSERT INTO users (name, email, password, is_onboarding_completed)
		VALUES ($1, $2, $3, $4)
		RETURNING *
	`

	var out domain.User

	err := r.db.QueryRow(ctx, query,
		u.Name,
		u.Email,
		u.Password,
		u.IsOnboardingCompleted,
	).Scan(
		&out.ID,
		&out.Name,
		&out.Email,
		&out.Password,
		&out.CreatedAt,
		&out.UpdatedAt,
		&out.IsOnboardingCompleted,
		&out.State,
	)

	if err != nil {
		return nil, err
	}

	return &out, nil
}

func (r *UserRepo) Update(ctx context.Context, u *domain.User) (*domain.User, error) {
	query := `
		UPDATE users
		SET
			name = $1,
			email = $2,
			is_onboarding_completed = $3,
			updated_at = NOW(),
			state = $4
		WHERE id = $5
		RETURNING *
	`

	var out domain.User

	err := r.db.QueryRow(ctx, query,
		u.Name,
		u.Email,
		u.IsOnboardingCompleted,
		u.State,
		u.ID,
	).Scan(
		&out.ID,
		&out.Name,
		&out.Email,
		&out.Password,
		&out.CreatedAt,
		&out.UpdatedAt,
		&out.IsOnboardingCompleted,
		&out.State,
	)

	if err != nil {
		return nil, err
	}

	return &out, nil
}

func (r *UserRepo) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM users WHERE id=$1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}

func (r *UserRepo) GetByID(ctx context.Context, id int) (*domain.User, error) {
	query := `
	SELECT
		id,
		name,
		email,
		password,
		created_at,
		updated_at,
		is_onboarding_completed,
		state
	FROM users WHERE id=$1
`

	var u domain.User

	err := r.db.QueryRow(ctx, query, id).
		Scan(
			&u.ID,
			&u.Name,
			&u.Email,
			&u.Password,
			&u.CreatedAt,
			&u.UpdatedAt,
			&u.IsOnboardingCompleted,
			&u.State,
		)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &u, nil
}

func (r *UserRepo) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `
		SELECT *
		FROM users WHERE email=$1
	`

	var u domain.User

	err := r.db.QueryRow(ctx, query, email).
		Scan(
			&u.ID,
			&u.Name,
			&u.Email,
			&u.Password,
			&u.CreatedAt,
			&u.UpdatedAt,
			&u.IsOnboardingCompleted,
			&u.State,
		)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &u, nil
}
