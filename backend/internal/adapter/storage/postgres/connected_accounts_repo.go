package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/x-sushant-x/connective/internal/core/dao"
	"github.com/x-sushant-x/connective/internal/core/domain"
)

type ConnectedAccountRepo struct {
	db *DB
}

func NewConnectedAccountRepo(db *DB) *ConnectedAccountRepo {
	return &ConnectedAccountRepo{db}
}

func (r *ConnectedAccountRepo) Create(ctx context.Context, ca *domain.ConnectedAccount) (*domain.ConnectedAccount, error) {
	query := `
		INSERT INTO connected_accounts
		(id, project_id, provider_id, external_user_id, external_team_id,
		 account_email, access_token, refresh_token, expires_at, token_type,
		 scope, raw_response, connected_at, user_id)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13, $14)
		RETURNING id, project_id, provider_id, external_user_id, external_team_id,
		          account_email, access_token, refresh_token, expires_at, token_type,
		          scope, raw_response, connected_at, created_at, updated_at, user_id
	`

	var out domain.ConnectedAccount

	err := r.db.QueryRow(ctx, query,
		ca.ID, ca.ProjectID, ca.ProviderID, ca.ExternalUserID, ca.ExternalTeamID,
		ca.AccountEmail, ca.AccessToken, ca.RefreshToken, ca.ExpiresAt, ca.TokenType,
		ca.Scope, ca.RawResponse, ca.ConnectedAt, ca.UserId,
	).Scan(
		&out.ID, &out.ProjectID, &out.ProviderID, &out.ExternalUserID, &out.ExternalTeamID,
		&out.AccountEmail, &out.AccessToken, &out.RefreshToken, &out.ExpiresAt, &out.TokenType,
		&out.Scope, &out.RawResponse, &out.ConnectedAt, &out.CreatedAt, &out.UpdatedAt, &out.UserId,
	)

	if err != nil {
		return nil, err
	}

	return &out, nil
}

func (r *ConnectedAccountRepo) GetByID(ctx context.Context, id string) (*domain.ConnectedAccount, error) {
	query := `
		SELECT *
		FROM connected_accounts
		WHERE id=$1
	`

	var out domain.ConnectedAccount

	err := r.db.QueryRow(ctx, query, id).Scan(
		&out.ID, &out.ProjectID, &out.ProviderID, &out.UserId, &out.ExternalUserID, &out.ExternalTeamID,
		&out.AccountEmail, &out.AccessToken, &out.RefreshToken, &out.ExpiresAt, &out.TokenType,
		&out.Scope, &out.RawResponse, &out.ConnectedAt, &out.CreatedAt, &out.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &out, nil
}

func (r *ConnectedAccountRepo) GetByProjectAndProvider(ctx context.Context, projectID string, providerID int) (*domain.ConnectedAccount, error) {
	query := `
		SELECT *
		FROM connected_accounts
		WHERE project_id=$1 AND provider_id=$2
	`

	var out domain.ConnectedAccount

	err := r.db.QueryRow(ctx, query, projectID, providerID).Scan(
		&out.ID, &out.ProjectID, &out.ProviderID, &out.UserId, &out.ExternalUserID, &out.ExternalTeamID,
		&out.AccountEmail, &out.AccessToken, &out.RefreshToken, &out.ExpiresAt, &out.TokenType,
		&out.Scope, &out.RawResponse, &out.ConnectedAt, &out.CreatedAt, &out.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &out, nil
}

func (r *ConnectedAccountRepo) Update(ctx context.Context, ca *domain.ConnectedAccount) (*domain.ConnectedAccount, error) {
	query := `
		UPDATE connected_accounts
		SET external_user_id=$1, external_team_id=$2, account_email=$3,
		    access_token=$4, refresh_token=$5, expires_at=$6, token_type=$7,
		    scope=$8, raw_response=$9, updated_at=NOW(), user_id=$10
		WHERE id=$11
		RETURNING id, project_id, provider_id, external_user_id, external_team_id,
		          account_email, access_token, refresh_token, expires_at, token_type,
		          scope, raw_response, connected_at, created_at, updated_at, user_id
	`

	var out domain.ConnectedAccount

	err := r.db.QueryRow(ctx, query,
		ca.ExternalUserID, ca.ExternalTeamID, ca.AccountEmail,
		ca.AccessToken, ca.RefreshToken, ca.ExpiresAt, ca.TokenType,
		ca.Scope, ca.RawResponse, ca.UserId, ca.ID,
	).Scan(
		&out.ID, &out.ProjectID, &out.ProviderID, &out.ExternalUserID, &out.ExternalTeamID,
		&out.AccountEmail, &out.AccessToken, &out.RefreshToken, &out.ExpiresAt, &out.TokenType,
		&out.Scope, &out.RawResponse, &out.ConnectedAt, &out.CreatedAt, &out.UpdatedAt, &out.UserId,
	)

	if err != nil {
		return nil, err
	}

	return &out, nil
}

func (r *ConnectedAccountRepo) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM connected_accounts WHERE id=$1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}

func (r *ConnectedAccountRepo) GetByProjectAndProviderAndUserId(ctx context.Context, projectID, userID string, providerID int) (*domain.ConnectedAccount, error) {
	query := `
		SELECT *
		FROM connected_accounts
		WHERE project_id=$1 AND provider_id=$2 AND user_id=$3
	`

	var out domain.ConnectedAccount

	err := r.db.QueryRow(ctx, query, projectID, providerID, userID).Scan(
		&out.ID, &out.ProjectID, &out.ProviderID, &out.UserId, &out.ExternalUserID, &out.ExternalTeamID,
		&out.AccountEmail, &out.AccessToken, &out.RefreshToken, &out.ExpiresAt, &out.TokenType,
		&out.Scope, &out.RawResponse, &out.ConnectedAt, &out.CreatedAt, &out.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &out, nil
}

func (r *ConnectedAccountRepo) GetConnectedUsers(ctx context.Context, projectId, search string) ([]dao.ConnectedUsers, error) {
	var resp []dao.ConnectedUsers

	query := `
		SELECT 
			ca.user_id,
			STRING_AGG(p.display_name, ', ') AS integrations_enabled,
			MIN(ca.created_at) AS date_created
		FROM connected_accounts ca
		JOIN providers p ON ca.provider_id = p.id
		WHERE ca.project_id = $1
			AND ($2 = '' OR ca.user_id ILIKE '%' || $2 || '%')
		GROUP BY ca.user_id;
	`

	rows, err := r.db.Query(ctx, query, projectId, search)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return resp, nil
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var r dao.ConnectedUsers

		if err := rows.Scan(&r.UserID, &r.IntegrationsEnabled, &r.DateCreated); err != nil {
			return nil, err
		}

		resp = append(resp, dao.BuildConnectedUser(r.UserID, r.IntegrationsEnabled, r.DateCreated))
	}

	return resp, nil
}
