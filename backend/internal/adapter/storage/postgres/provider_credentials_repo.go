package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/x-sushant-x/connective/internal/core/domain"
)

type ProviderCredentialsRepo struct {
	db *DB
}

func NewProviderCredentialsRepo(db *DB) *ProviderCredentialsRepo {
	return &ProviderCredentialsRepo{db}
}

func (r *ProviderCredentialsRepo) Create(ctx context.Context, c *domain.ProviderCredentials) (*domain.ProviderCredentials, error) {
	query := `
		INSERT INTO provider_credentials
		(project_id, provider_id, client_id, client_secret, scopes, is_active)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, project_id, provider_id, client_id, client_secret,
		          scopes, is_active, created_at, updated_at
	`

	var out domain.ProviderCredentials

	err := r.db.QueryRow(ctx, query,
		c.ProjectID, c.ProviderID, c.ClientID, c.ClientSecret,
		c.Scopes, c.IsActive,
	).Scan(
		&out.ID, &out.ProjectID, &out.ProviderID, &out.ClientID, &out.ClientSecret,
		&out.Scopes, &out.IsActive, &out.CreatedAt, &out.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &out, nil
}

func (r *ProviderCredentialsRepo) GetByProjectAndProvider(ctx context.Context, projectID string, providerID int) (*domain.ProviderCredentials, error) {
	query := `
		SELECT id, project_id, provider_id, client_id, client_secret,scopes, is_active, created_at, updated_at
		FROM provider_credentials
		WHERE project_id=$1 AND provider_id=$2
	`

	var out domain.ProviderCredentials

	err := r.db.QueryRow(ctx, query, projectID, providerID).Scan(
		&out.ID, &out.ProjectID, &out.ProviderID, &out.ClientID, &out.ClientSecret,
		&out.Scopes, &out.IsActive, &out.CreatedAt, &out.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &out, nil
}

func (r *ProviderCredentialsRepo) Update(ctx context.Context, c *domain.ProviderCredentials) (*domain.ProviderCredentials, error) {
	query := `
		UPDATE provider_credentials
		SET client_id=$1, client_secret=$2,
		    scopes=$3, is_active=$4, updated_at=NOW()
		WHERE id=$5
		RETURNING id, project_id, provider_id, client_id, client_secret,
		        scopes, is_active, created_at, updated_at
	`

	var out domain.ProviderCredentials

	err := r.db.QueryRow(ctx, query,
		c.ClientID, c.ClientSecret,
		c.Scopes, c.IsActive, c.ID,
	).Scan(
		&out.ID, &out.ProjectID, &out.ProviderID, &out.ClientID, &out.ClientSecret,
		&out.Scopes, &out.IsActive, &out.CreatedAt, &out.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &out, nil
}

func (r *ProviderCredentialsRepo) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM provider_credentials WHERE id=$1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}
