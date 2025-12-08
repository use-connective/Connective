package dao

type ProviderListDisplayable struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	ImageURL    string `json:"image_url"`
}
