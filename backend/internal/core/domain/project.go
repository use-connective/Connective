package domain

import "time"

type Project struct {
	ID            string     `json:"id"`
	Name          string     `json:"name"`
	Owner         int        `json:"owner"`
	SDKAuthSecret string     `json:"sdk_auth_secret"`
	CreatedAt     *time.Time `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
}
