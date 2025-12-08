package domain

import (
	"encoding/json"
	"time"
)

type ConnectedAccount struct {
	ID             string          `json:"id"`
	ProjectID      string          `json:"project_id"`
	ProviderID     int             `json:"provider_id"`
	UserId         string          `json:"user_id"`
	ExternalUserID string          `json:"external_user_id"`
	ExternalTeamID string          `json:"external_team_id"`
	AccountEmail   string          `json:"account_email"`
	AccessToken    string          `json:"access_token"`
	RefreshToken   string          `json:"refresh_token"`
	ExpiresAt      *time.Time      `json:"expires_at"`
	TokenType      string          `json:"token_type"`
	Scope          string          `json:"scope"`
	RawResponse    json.RawMessage `json:"raw_response"`
	ConnectedAt    time.Time       `json:"connected_at"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
}
