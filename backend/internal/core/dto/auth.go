package dto

import "github.com/x-sushant-x/connective/internal/core/domain"

type LoginResponse struct {
	User  *domain.User `json:"user"`
	Token string       `json:"token"`
}

type SignupResponse struct {
	User  *domain.User `json:"user"`
	Token string       `json:"token"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
