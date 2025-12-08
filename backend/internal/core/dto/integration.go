package dto

type CreateUpdateProviderCreds struct {
	ProjectID    string   `json:"project_id"`
	ProviderID   int      `json:"provider_id"`
	ClientID     string   `json:"client_id"`
	ClientSecret string   `json:"client_secret"`
	Scopes       []string `json:"scopes"`
}
