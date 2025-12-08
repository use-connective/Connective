package iSlack

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/x-sushant-x/connective/internal/connectors/common"
	"github.com/x-sushant-x/connective/internal/core/port"
)

type SlackConnector struct {
	providerRepo port.ProviderRepo
}

func New(providerRepo port.ProviderRepo) *SlackConnector {
	return &SlackConnector{
		providerRepo,
	}
}

func (s *SlackConnector) Name() string {
	return "slack"
}

// AuthStrategy TODO - Handle cases where app is deleted by user of access is revoked by user.
func (s *SlackConnector) AuthStrategy() common.AuthStrategy {
	ctx := context.Background()

	provider, err := s.providerRepo.GetProviderByName(ctx, "slack")
	if err != nil || provider == nil {
		// TODO - See where it impacts and handle accordingly.
		log.Err(err).Msg("Unable to get auth strategy for slack")
		return nil
	}

	return NewSlackStrategy(provider)
}

func (s *SlackConnector) Actions() map[string]common.ActionHandler {
	return map[string]common.ActionHandler{
		"chat.postMessage":   s.SendMessage,
		"conversations.list": s.ListChannels,
	}
}
