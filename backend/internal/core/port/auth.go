package port

import (
	"context"

	"github.com/x-sushant-x/connective/internal/core/domain"
	"github.com/x-sushant-x/connective/internal/core/dto"
)

type AuthService interface {
	CreateUser(ctx context.Context, user *domain.User) (*dto.SignupResponse, error)
	LoginUser(ctx context.Context, email, password string) (*dto.LoginResponse, error)
	MarkOnboardingCompleted(ctx context.Context, userID int) error
}

type AuthRepo interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	GetUser(ctx context.Context, email string) (*domain.User, error)
	GetUserById(ctx context.Context, userID int) (*domain.User, error)
}
