package dao

import "time"

type ConnectedUsers struct {
	UserID              string     `json:"user_id"`
	IntegrationsEnabled string     `json:"integrations_enabled"`
	DateCreated         *time.Time `json:"date_created"`
	DisplayableDate     string     `json:"displayable_date"`
}

func BuildConnectedUser(userID, integrationsEnabled string, dateCreated *time.Time) ConnectedUsers {
	dateOnly := dateCreated.Format("2006-01-02")

	return ConnectedUsers{
		UserID:              userID,
		IntegrationsEnabled: integrationsEnabled,
		DisplayableDate:     dateOnly,
	}
}
