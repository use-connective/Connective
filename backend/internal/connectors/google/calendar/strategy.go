package googleCalendar

import (
	"net/url"
	"strings"

	"github.com/x-sushant-x/connective/internal/connectors/common"
	"github.com/x-sushant-x/connective/internal/core/domain"
)

type Strategy struct {
	common.OAuth2Strategy
}

func NewGoogleCalendarStrategy(provider *domain.Provider) common.AuthStrategy {
	return Strategy{
		common.OAuth2Strategy{
			AuthURLVal:         provider.AuthURL,
			TokenURLVal:        provider.TokenURL,
			ScopesVal:          provider.DefaultScopes,
			RedirectURLVal:     provider.RedirectURL,
			RefreshTokenURLVal: provider.RefreshTokenURL,
		},
	}
}

func (g Strategy) BuildAuthURL(providerCredentials *domain.ProviderCredentials, projectID, userID string) string {
	state := projectID + ":" + userID

	params := url.Values{}
	params.Set("client_id", providerCredentials.ClientID)
	params.Set("redirect_uri", g.RedirectURLVal)
	params.Set("response_type", "code")
	params.Set("scope", strings.Join(providerCredentials.Scopes, " "))
	params.Set("access_type", "offline")
	params.Set("prompt", "consent")
	params.Set("state", state)

	return g.AuthURLVal + "?" + params.Encode()
}

func (g Strategy) PrepareTokenRequest(providerCredentials *domain.ProviderCredentials, code string, reqData url.Values) {
	reqData.Set("client_id", providerCredentials.ClientID)
	reqData.Set("client_secret", providerCredentials.ClientSecret)
	reqData.Set("redirect_uri", g.RedirectURLVal)
	reqData.Set("grant_type", "authorization_code")
	reqData.Set("code", code)
}
