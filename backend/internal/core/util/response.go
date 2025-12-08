package util

import "github.com/x-sushant-x/connective/internal/core/domain"

// ToUserResponse takes a user and returns a copy excluding password.
func ToUserResponse(u *domain.User) *domain.User {
	return &domain.User{
		ID:                    u.ID,
		Name:                  u.Name,
		Email:                 u.Email,
		CreatedAt:             u.CreatedAt,
		UpdatedAt:             u.UpdatedAt,
		IsOnboardingCompleted: u.IsOnboardingCompleted,
		State:                 u.State,
	}
}
