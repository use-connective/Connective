// TODO - Add check for revoked token.

package githubConnector

import (
	"github.com/x-sushant-x/connective/internal/connectors/common"
	"github.com/x-sushant-x/connective/internal/core/domain"
)

type Strategy struct {
	common.OAuth2Strategy
}

func NewGithubStrategy(provider *domain.Provider) common.AuthStrategy {
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

func (s Strategy) ParseTokenResponse(body []byte) (*common.TokenResponse, error) {
	sBody := string(body)

	token := sBody[13:]

	return &common.TokenResponse{
		Ok:           true,
		AccessToken:  token,
		RefreshToken: "",
	}, nil
}
