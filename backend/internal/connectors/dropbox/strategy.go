package dropboxConnector

import (
	"net/url"
	"strings"

	"github.com/x-sushant-x/connective/internal/connectors/common"
	"github.com/x-sushant-x/connective/internal/core/domain"
)

type Strategy struct {
	common.AuthStrategy
}

func NewDropboxStrategy(provider *domain.Provider) common.AuthStrategy {
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

func (s Strategy) BuildAuthURL(providerCredentials *domain.ProviderCredentials, projectID, userID string) string {
	state := projectID + ":" + userID

	u := s.AuthURL() + "?" +
		"client_id=" + providerCredentials.ClientID +
		"&scope=" + strings.Join(providerCredentials.Scopes, " ") +
		"&redirect_uri=" + s.RedirectURL() +
		"&state=" + state +
		"&response_type=code" +
		"&token_access_type=offline"

	return u
}

func (s Strategy) PrepareTokenRequest(providerCredentials *domain.ProviderCredentials, code string, reqData url.Values) {
	reqData.Set("client_id", providerCredentials.ClientID)
	reqData.Set("client_secret", providerCredentials.ClientSecret)
	reqData.Set("code", code)
	reqData.Set("redirect_uri", s.RedirectURL())
	reqData.Set("grant_type", "authorization_code")
}
