package domain

import "time"

type Provider struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	DisplayName     string    `json:"display_name"`
	AuthType        string    `json:"auth_type"`
	ImageURL        string    `json:"image_url"`
	Category        string    `json:"category"`
	Description     string    `json:"description"`
	PrimaryColor    string    `json:"primary_color"`
	AuthURL         string    `json:"auth_url"`
	TokenURL        string    `json:"token_url"`
	RefreshTokenURL string    `json:"refresh_token_url"`
	RedirectURL     string    `json:"redirect_url"`
	DefaultScopes   []string  `json:"default_scopes"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
