package common

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/x-sushant-x/connective/internal/core/domain"
)

type AuthStrategy interface {
	AuthURL() string
	TokenURL() string
	Scopes() []string
	RedirectURL() string
	RefreshTokenURL() string

	BuildAuthURL(providerCredentials *domain.ProviderCredentials, projectID, userID string) string
	PrepareTokenRequest(providerCredentials *domain.ProviderCredentials, code string, reqData url.Values)
	ParseTokenResponse(body []byte) (*TokenResponse, error)
	GetNewAccessToken(provider *domain.Provider, providerCredentials *domain.ProviderCredentials, connAcc *domain.ConnectedAccount) (*TokenResponse, error)
}

type OAuth2Strategy struct {
	AuthURLVal         string
	TokenURLVal        string
	ScopesVal          []string
	RedirectURLVal     string
	RefreshTokenURLVal string
}

func (o OAuth2Strategy) AuthURL() string {
	return o.AuthURLVal
}

func (o OAuth2Strategy) TokenURL() string {
	return o.TokenURLVal
}

func (o OAuth2Strategy) Scopes() []string {
	return o.ScopesVal
}

func (o OAuth2Strategy) RedirectURL() string {
	return o.RedirectURLVal
}

func (o OAuth2Strategy) RefreshTokenURL() string {
	return o.RefreshTokenURLVal
}

func (o OAuth2Strategy) BuildAuthURL(providerCredentials *domain.ProviderCredentials, projectID, userID string) string {
	state := projectID + ":" + userID

	u := o.AuthURLVal + "?" +
		"client_id=" + providerCredentials.ClientID +
		"&scope=" + strings.Join(providerCredentials.Scopes, ",") +
		"&redirect_uri=" + o.RedirectURLVal +
		"&state=" + state

	return u
}

func (o OAuth2Strategy) PrepareTokenRequest(providerCredentials *domain.ProviderCredentials, code string, reqData url.Values) {
	reqData.Set("client_id", providerCredentials.ClientID)
	reqData.Set("client_secret", providerCredentials.ClientSecret)
	reqData.Set("code", code)
	reqData.Set("redirect_uri", o.RedirectURLVal)
}

func (o OAuth2Strategy) ParseTokenResponse(body []byte) (*TokenResponse, error) {
	var token TokenResponse
	if err := json.Unmarshal(body, &token); err != nil {
		return nil, err
	}
	return &token, nil
}

func (o OAuth2Strategy) GetNewAccessToken(provider *domain.Provider, providerCredentials *domain.ProviderCredentials, connAcc *domain.ConnectedAccount) (*TokenResponse, error) {

	data := url.Values{}
	data.Set("client_id", providerCredentials.ClientID)
	data.Set("client_secret", providerCredentials.ClientSecret)
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", connAcc.RefreshToken)

	resp, err := http.PostForm(provider.RefreshTokenURL, data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var token TokenResponse
	if err := json.Unmarshal(bodyBytes, &token); err != nil {
		return nil, err
	}

	return &token, nil
}
