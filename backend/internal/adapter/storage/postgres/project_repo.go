package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/x-sushant-x/connective/internal/core/domain"
)

type ProjectRepo struct {
	db *DB
}

func NewProjectRepo(db *DB) *ProjectRepo {
	return &ProjectRepo{db}
}

func (r *ProjectRepo) Create(ctx context.Context, p *domain.Project) (*domain.Project, error) {
	query := `
		INSERT INTO projects (id, name, owner, sdk_auth_secret)
		VALUES ($1, $2, $3, $4)
		RETURNING *
	`

	var out domain.Project

	err := r.db.QueryRow(ctx, query, p.ID, p.Name, p.Owner, p.SDKAuthSecret).
		Scan(&out.ID, &out.Name, &out.Owner, &out.SDKAuthSecret, &out.CreatedAt, &out.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

func (r *ProjectRepo) GetByID(ctx context.Context, id string) (*domain.Project, error) {
	query := `
		SELECT *
		FROM projects WHERE id=$1
	`

	var p domain.Project

	err := r.db.QueryRow(ctx, query, id).
		Scan(&p.ID, &p.Name, &p.Owner, &p.SDKAuthSecret, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &p, nil
}

func (r *ProjectRepo) GetByOwner(ctx context.Context, userId int) ([]domain.Project, error) {
	query := `
		SELECT *
		FROM projects WHERE owner=$1
	`

	rows, err := r.db.Query(ctx, query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []domain.Project

	for rows.Next() {
		var p domain.Project
		err := rows.Scan(&p.ID, &p.Name, &p.Owner, &p.SDKAuthSecret, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return nil, err
		}
		list = append(list, p)
	}

	return list, nil
}

func (r *ProjectRepo) Update(ctx context.Context, p *domain.Project) (*domain.Project, error) {
	query := `
		UPDATE projects
		SET name=$1, updated_at=NOW()
		WHERE id=$2
		RETURNING *
	`

	var out domain.Project

	err := r.db.QueryRow(ctx, query, p.Name, p.ID).
		Scan(&out.ID, &out.Name, &out.Owner, &out.SDKAuthSecret, &out.CreatedAt, &out.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

func (r *ProjectRepo) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM projects WHERE id=$1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}

func (r *ProjectRepo) GetByOwnerAndName(ctx context.Context, ownerID int, name string) (*domain.Project, error) {
	query := `
		SELECT *
		FROM projects
		WHERE owner = $1 AND name = $2
		LIMIT 1
	`

	var project domain.Project

	err := r.db.
		QueryRow(ctx, query, ownerID, name).
		Scan(&project.ID, &project.Name, &project.Owner, &project.SDKAuthSecret, &project.CreatedAt, &project.UpdatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &project, nil
}
