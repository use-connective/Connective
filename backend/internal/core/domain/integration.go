package domain

import "time"

type ProviderCredentials struct {
	ID           int       `json:"id"`
	ProjectID    string    `json:"project_id"`
	ProviderID   int       `json:"provider_id"`
	ClientID     string    `json:"client_id"`
	ClientSecret string    `json:"client_secret"`
	Scopes       []string  `json:"scopes"`
	IsActive     bool      `json:"is_active"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
