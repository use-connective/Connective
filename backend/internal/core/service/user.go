package service

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/x-sushant-x/connective/internal/core/domain"
	cErrors "github.com/x-sushant-x/connective/internal/core/errors"
	"github.com/x-sushant-x/connective/internal/core/port"
)

type UserService struct {
	repo port.UserRepo
}

func NewUserService(repo port.UserRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}

// Create creates a new user.
func (us *UserService) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	// Check if user exists
	exists, err := us.repo.GetByEmail(ctx, user.Email)
	if err != nil {
		log.Error().Err(err).Str("email", user.Email).Msg("UserService: unable to check existing user")
		return nil, cErrors.ErrUnableToCreateAccount // or define a user-specific error
	}

	if exists != nil {
		return nil, cErrors.ErrUserAlreadyExists
	}

	// Save new user
	created, err := us.repo.Create(ctx, user)
	if err != nil {
		log.Error().Err(err).Str("email", user.Email).Msg("UserService: unable to create user")
		return nil, cErrors.ErrUnableToCreateAccount
	}

	return created, nil
}

// Update updates an existing user.
func (us *UserService) Update(ctx context.Context, user *domain.User) (*domain.User, error) {
	// Check if user exists
	existing, err := us.repo.GetByID(ctx, user.ID)
	if err != nil {
		log.Error().Err(err).Int("id", user.ID).Msg("UserService: unable to get user for update")
		return nil, cErrors.ErrUnableToUpdateUser
	}

	if existing == nil {
		return nil, cErrors.ErrUserDoesNotExists
	}

	updated, err := us.repo.Update(ctx, user)
	if err != nil {
		log.Error().Err(err).Int("id", user.ID).Msg("UserService: unable to update user")
		return nil, cErrors.ErrUnableToUpdateUser
	}

	return updated, nil
}

// Delete removes a user.
func (us *UserService) Delete(ctx context.Context, id int) error {
	// Check if exists first
	existing, err := us.repo.GetByID(ctx, id)
	if err != nil {
		log.Error().Err(err).Int("id", id).Msg("UserService: unable to get user for delete")
		return cErrors.ErrUnableToDeleteUser
	}

	if existing == nil {
		return cErrors.ErrUserDoesNotExists
	}

	// Delete user
	err = us.repo.Delete(ctx, id)
	if err != nil {
		log.Error().Err(err).Int("id", id).Msg("UserService: unable to delete user")
		return cErrors.ErrUnableToDeleteUser
	}

	return nil
}

// GetByID returns a user by ID.
func (us *UserService) GetByID(ctx context.Context, id int) (*domain.User, error) {
	user, err := us.repo.GetByID(ctx, id)
	if err != nil {
		log.Error().Err(err).Int("id", id).Msg("UserService: unable to get user by ID")
		return nil, cErrors.ErrUnableToGetUser
	}

	if user == nil {
		return nil, cErrors.ErrUserDoesNotExists
	}

	return user, nil
}

// GetByEmail returns a user by email.
func (us *UserService) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	user, err := us.repo.GetByEmail(ctx, email)
	if err != nil {
		log.Error().Err(err).Str("email", email).Msg("UserService: unable to get user by email")
		return nil, cErrors.ErrUnableToGetUser
	}

	if user == nil {
		return nil, cErrors.ErrUserDoesNotExists
	}

	return user, nil
}
