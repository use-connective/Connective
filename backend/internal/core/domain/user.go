package domain

import "time"

type UserState string

const (
	UserStateProjectPending              UserState = "PROJECT_PENDING"
	UserStateActive                      UserState = "ACTIVE"
	UserStateOnboardingCompletionPending UserState = "ONBOARDING_COMPLETION_PENDING"
)

type User struct {
	ID                    int        `json:"id"`
	Name                  string     `json:"name"`
	Email                 string     `json:"email"`
	Password              string     `json:"password,omitempty"`
	CreatedAt             *time.Time `json:"created_at"`
	UpdatedAt             *time.Time `json:"updated_at"`
	IsOnboardingCompleted bool       `json:"is_onboarding_completed"`
	State                 UserState  `json:"state"`
}
