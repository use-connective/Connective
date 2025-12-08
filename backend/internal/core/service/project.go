package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/x-sushant-x/connective/internal/core/domain"
	"github.com/x-sushant-x/connective/internal/core/dto"
	cErrors "github.com/x-sushant-x/connective/internal/core/errors"
	"github.com/x-sushant-x/connective/internal/core/port"
)

type ProjectService struct {
	repo    port.ProjectRepo
	userSvc port.UserService
}

func NewProjectService(repo port.ProjectRepo, userSvc port.UserService) *ProjectService {
	return &ProjectService{
		repo,
		userSvc,
	}
}

func (ps *ProjectService) CreateProject(ctx context.Context, user *domain.User, req *dto.CreateProjectRequest) (*domain.Project, error) {

	existing, err := ps.repo.GetByOwnerAndName(ctx, user.ID, req.Name)

	if err != nil {
		log.Err(err).
			Str("project name", req.Name).
			Int("owner", user.ID).
			Msg("unable to get existing project while creating new.")
		return nil, cErrors.ErrUnableToCreateProject
	}

	if existing != nil {
		return nil, cErrors.ErrProjectExists
	}

	newProject := domain.Project{
		ID:            uuid.New().String(),
		Name:          req.Name,
		Owner:         user.ID,
		SDKAuthSecret: uuid.NewString(),
	}

	createdProject, err := ps.repo.Create(ctx, &newProject)
	if err != nil {
		log.Err(err).
			Str("project name", req.Name).
			Int("owner", user.ID).
			Msg("unable to create new project")
		return nil, cErrors.ErrUnableToCreateProject
	}

	user.State = domain.UserStateOnboardingCompletionPending

	_, err = ps.userSvc.Update(ctx, user)
	if err != nil {
		return nil, cErrors.ErrUnableToCreateProject
	}

	return createdProject, nil
}

func (ps *ProjectService) GetByOwner(ctx context.Context, ownerID int) ([]domain.Project, error) {
	projects, err := ps.repo.GetByOwner(ctx, ownerID)
	if err != nil {
		log.Err(err).
			Int("owner", ownerID).
			Msg("unable to get all project")
		return nil, err
	}

	return projects, nil
}
