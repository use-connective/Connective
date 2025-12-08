package iSlack

import (
	"github.com/x-sushant-x/connective/internal/connectors/common"
	"github.com/x-sushant-x/connective/internal/core/domain"
)

type Strategy struct {
	common.AuthStrategy
}

func NewSlackStrategy(provider *domain.Provider) common.AuthStrategy {
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
