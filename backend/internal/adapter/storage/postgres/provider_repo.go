package postgres

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/x-sushant-x/connective/internal/core/dao"
	"github.com/x-sushant-x/connective/internal/core/domain"
)

type ProviderRepo struct {
	db *DB
}

func NewProviderRepo(db *DB) *ProviderRepo {
	return &ProviderRepo{db}
}

func (r *ProviderRepo) Create(ctx context.Context, p *domain.Provider) (*domain.Provider, error) {
	query := `
		INSERT INTO providers (
			name, display_name, auth_type, image_url, category,
			description, primary_color,
			auth_url, token_url, refresh_token_url, redirect_url, default_scopes
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		RETURNING
			*
	`

	var out domain.Provider

	err := r.db.QueryRow(ctx, query,
		p.Name, p.DisplayName, p.AuthType, p.ImageURL, p.Category,
		p.Description, p.PrimaryColor,
		p.AuthURL, p.TokenURL, p.RefreshTokenURL, p.RedirectURL, p.DefaultScopes,
	).Scan(
		&out.ID, &out.Name, &out.DisplayName, &out.AuthType, &out.ImageURL,
		&out.Category, &out.Description, &out.PrimaryColor,
		&out.AuthURL, &out.TokenURL, &out.RefreshTokenURL, &out.RedirectURL,
		&out.DefaultScopes, &out.CreatedAt, &out.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

func (r *ProviderRepo) GetByID(ctx context.Context, id int) (*domain.Provider, error) {
	query := `
		SELECT *
		FROM providers
		WHERE id = $1
	`

	var p domain.Provider

	err := r.db.QueryRow(ctx, query, id).Scan(
		&p.ID, &p.Name, &p.DisplayName, &p.AuthType, &p.ImageURL,
		&p.Category, &p.Description, &p.PrimaryColor,
		&p.AuthURL, &p.TokenURL, &p.RefreshTokenURL, &p.RedirectURL,
		&p.DefaultScopes, &p.CreatedAt, &p.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &p, nil
}

func (r *ProviderRepo) GetAll(ctx context.Context) ([]*domain.Provider, error) {
	query := `
		SELECT *
		FROM providers
		ORDER BY id
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var providers []*domain.Provider

	for rows.Next() {
		var p domain.Provider
		err := rows.Scan(
			&p.ID, &p.Name, &p.DisplayName, &p.AuthType, &p.ImageURL,
			&p.Category, &p.Description, &p.PrimaryColor,
			&p.AuthURL, &p.TokenURL, &p.RefreshTokenURL, &p.RedirectURL,
			&p.DefaultScopes, &p.CreatedAt, &p.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		providers = append(providers, &p)
	}

	return providers, nil
}

func (r *ProviderRepo) Update(ctx context.Context, p *domain.Provider) (*domain.Provider, error) {
	query := `
		UPDATE providers
		SET
			name=$1,
			display_name=$2,
			auth_type=$3,
			image_url=$4,
			category=$5,
			description=$6,
			primary_color=$7,
			auth_url=$8,
			token_url=$9,
			refresh_token_url=$10,
			redirect_url=$11,
			default_scopes=$12,
			updated_at = NOW()
		WHERE id=$13
		RETURNING
			*
	`

	var out domain.Provider

	err := r.db.QueryRow(ctx, query,
		p.Name, p.DisplayName, p.AuthType, p.ImageURL, p.Category,
		p.Description, p.PrimaryColor,
		p.AuthURL, p.TokenURL, p.RefreshTokenURL, p.RedirectURL, p.DefaultScopes,
		p.ID,
	).Scan(
		&out.ID, &out.Name, &out.DisplayName, &out.AuthType, &out.ImageURL,
		&out.Category, &out.Description, &out.PrimaryColor,
		&out.AuthURL, &out.TokenURL, &out.RefreshTokenURL, &out.RedirectURL,
		&out.DefaultScopes, &out.CreatedAt, &out.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

func (r *ProviderRepo) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM providers WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}

func (r *ProviderRepo) GetProviderByName(ctx context.Context, name string) (*domain.Provider, error) {
	query := `
		SELECT *
		FROM providers
		WHERE name = $1
	`

	var p domain.Provider

	err := r.db.QueryRow(ctx, query, name).Scan(
		&p.ID, &p.Name, &p.DisplayName, &p.AuthType, &p.ImageURL,
		&p.Category, &p.Description, &p.PrimaryColor,
		&p.AuthURL, &p.TokenURL, &p.RefreshTokenURL, &p.RedirectURL,
		&p.DefaultScopes, &p.CreatedAt, &p.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &p, nil
}

func (r *ProviderRepo) GetProviderListDisplayable(ctx context.Context, search, category string) ([]dao.ProviderListDisplayable, error) {
	resp := make([]dao.ProviderListDisplayable, 0)

	baseQuery := `SELECT 
        id, name, display_name, image_url
        FROM providers`

	var conditions []string
	var args []interface{}
	idx := 1

	if search != "" {
		conditions = append(conditions, fmt.Sprintf("name ILIKE $%d", idx))
		args = append(args, "%"+search+"%")
		idx++
	}

	if category != "" {
		conditions = append(conditions, fmt.Sprintf("category ILIKE $%d", idx))
		args = append(args, category)
		idx++
	}

	if len(conditions) > 0 {
		baseQuery += " WHERE " + strings.Join(conditions, " AND ")
	}

	rows, err := r.db.Query(ctx, baseQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p dao.ProviderListDisplayable
		if err := rows.Scan(&p.ID, &p.Name, &p.DisplayName, &p.ImageURL); err != nil {
			return nil, err
		}
		resp = append(resp, p)
	}

	return resp, nil
}
