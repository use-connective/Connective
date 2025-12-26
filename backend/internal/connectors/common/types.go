package common

import (
	"context"
	"time"
)

type Connector interface {
	Name() string
	AuthStrategy() AuthStrategy
	GetAction(ctx context.Context, actionName string) *ConnectorAction
}

type TokenResponse struct {
	Ok           bool   `json:"ok"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	Error        string `json:"error"`
}

type ActionExecuteRequest struct {
	ProjectSecret string         `json:"project_secret"`
	UserID        string         `json:"user_id"`
	Body          map[string]any `json:"request_body"`
	ProjectID     string         `json:"project_id"`
}

type SlackChannelList struct {
	Ok       bool `json:"ok"`
	Channels []struct {
		ID                 string `json:"id"`
		Name               string `json:"name"`
		IsChannel          bool   `json:"is_channel"`
		IsGroup            bool   `json:"is_group"`
		IsIm               bool   `json:"is_im"`
		Created            int    `json:"created"`
		Creator            string `json:"creator"`
		IsArchived         bool   `json:"is_archived"`
		IsGeneral          bool   `json:"is_general"`
		Unlinked           int    `json:"unlinked"`
		NameNormalized     string `json:"name_normalized"`
		IsShared           bool   `json:"is_shared"`
		IsExtShared        bool   `json:"is_ext_shared"`
		IsOrgShared        bool   `json:"is_org_shared"`
		PendingShared      []any  `json:"pending_shared"`
		IsPendingExtShared bool   `json:"is_pending_ext_shared"`
		IsMember           bool   `json:"is_member"`
		IsPrivate          bool   `json:"is_private"`
		IsMpim             bool   `json:"is_mpim"`
		Updated            int64  `json:"updated"`
		Topic              struct {
			Value   string `json:"value"`
			Creator string `json:"creator"`
			LastSet int    `json:"last_set"`
		} `json:"topic"`
		Purpose struct {
			Value   string `json:"value"`
			Creator string `json:"creator"`
			LastSet int    `json:"last_set"`
		} `json:"purpose"`
		PreviousNames []any `json:"previous_names"`
		NumMembers    int   `json:"num_members"`
	} `json:"channels"`
	ResponseMetadata struct {
		NextCursor string `json:"next_cursor"`
	} `json:"response_metadata"`
}

type UserCredentials struct {
	AccessToken  string
	RefreshToken string
	ExpiresAt    *time.Time
}

type ConnectorAction struct {
	CacheKey string            `json:"cacheKey"`
	Key      string            `json:"key"`
	Name     string            `json:"name"`
	Method   string            `json:"method"`
	URL      string            `json:"url"`
	Query    map[string]string `json:"query"`
	Headers  map[string]string `json:"headers"`
	Body     map[string]any    `json:"body"`
}
