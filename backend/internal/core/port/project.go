package port

import (
	"context"

	"github.com/x-sushant-x/connective/internal/core/domain"
	"github.com/x-sushant-x/connective/internal/core/dto"
)

type ProjectService interface {
	CreateProject(ctx context.Context, user *domain.User, req *dto.CreateProjectRequest) (*domain.Project, error)
	GetByOwner(ctx context.Context, ownerID int) ([]domain.Project, error)
}

type ProjectRepo interface {
	Create(ctx context.Context, project *domain.Project) (*domain.Project, error)
	Update(ctx context.Context, project *domain.Project) (*domain.Project, error)
	Delete(ctx context.Context, id string) error
	GetByID(ctx context.Context, id string) (*domain.Project, error)
	GetByOwner(ctx context.Context, ownerID int) ([]domain.Project, error)
	GetByOwnerAndName(ctx context.Context, ownerID int, name string) (*domain.Project, error)
}
