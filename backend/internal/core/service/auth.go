package service

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/x-sushant-x/connective/internal/core/domain"
	"github.com/x-sushant-x/connective/internal/core/dto"
	cErrors "github.com/x-sushant-x/connective/internal/core/errors"
	"github.com/x-sushant-x/connective/internal/core/port"
	"github.com/x-sushant-x/connective/internal/core/util"
)

type AuthService struct {
	repo     port.AuthRepo
	userRepo port.UserRepo
}

func NewAuthService(repo port.AuthRepo, userRepo port.UserRepo) *AuthService {
	return &AuthService{
		repo,
		userRepo,
	}
}

func (as *AuthService) CreateUser(ctx context.Context, userReq *domain.User) (*dto.SignupResponse, error) {
	existingUser, err := as.repo.GetUser(ctx, userReq.Email)
	if err != nil {
		log.Err(err).Str("email", userReq.Email).Msg("unable to get user for signup")
		return nil, cErrors.ErrUnableToCreateAccount
	}

	if existingUser != nil {
		return nil, cErrors.ErrUserAlreadyExists
	}

	rawPass := userReq.Password

	hashPass, err := util.HashPassword(rawPass)
	if err != nil {
		log.Err(err).Str("email", userReq.Email).Msg("unable to hash password")
		return nil, cErrors.ErrUnableToCreateAccount
	}

	userReq.Password = hashPass
	newUser, err := as.repo.CreateUser(ctx, userReq)
	if err != nil {
		log.Err(err).Str("email", userReq.Email).Msg("unable to create user")
		return nil, cErrors.ErrUnableToCreateAccount
	}

	token, err := util.GenerateJWTToken(newUser.ID, userReq.Email)
	if err != nil {
		log.Err(err).Str("email", userReq.Email).Msg("unable to create jwt token")
		return nil, cErrors.ErrUnableToCreateAccount
	}

	return &dto.SignupResponse{
		User:  newUser,
		Token: token,
	}, nil
}

func (as *AuthService) LoginUser(ctx context.Context, email, password string) (*dto.LoginResponse, error) {
	user, err := as.repo.GetUser(ctx, email)
	if err != nil {
		log.Err(err).Str("email", email).Msg("unable to get user for login")
		return nil, cErrors.ErrUnableToLogin
	}

	if user == nil {
		return nil, cErrors.ErrUserDoesNotExists
	}

	if isValid := util.CheckPasswordHash(password, user.Password); !isValid {
		return nil, cErrors.ErrWrongPassword
	}

	token, err := util.GenerateJWTToken(user.ID, user.Email)
	if err != nil {
		log.Err(err).Str("email", user.Email).Msg("unable to create jwt token")
		return nil, cErrors.ErrUnableToLogin
	}

	return &dto.LoginResponse{
		User:  util.ToUserResponse(user),
		Token: token,
	}, nil
}

func (as *AuthService) MarkOnboardingCompleted(ctx context.Context, userID int) error {
	user, err := as.userRepo.GetByID(ctx, userID)

	if err != nil {
		log.Err(err).Int("userID", userID).Msg("unable to find user")
		return cErrors.ErrUnableToCompleteOnboarding
	}

	user.IsOnboardingCompleted = true
	user.State = domain.UserStateActive

	_, err = as.userRepo.Update(ctx, user)
	if err != nil {
		log.Err(err).Int("userID", userID).Msg("unable to update user for onboarding completion")
		return cErrors.ErrUnableToCompleteOnboarding
	}

	return nil
}
