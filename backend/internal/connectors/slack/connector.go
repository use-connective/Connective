package iSlack

import (
	"context"
	"encoding/json"

	"github.com/rs/zerolog/log"
	"github.com/x-sushant-x/connective/internal/connectors/common"
	"github.com/x-sushant-x/connective/internal/core/port"
)

type SlackConnector struct {
	providerRepo port.ProviderRepo
	cache        port.Cache
}

func New(providerRepo port.ProviderRepo, cache port.Cache) *SlackConnector {
	return &SlackConnector{
		providerRepo,
		cache,
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

func (s *SlackConnector) GetAction(ctx context.Context, actionName string) *common.ConnectorAction {
	actionCacheKey := s.Name() + "_" + actionName

	actionStr, err := s.cache.GetJson(ctx, actionCacheKey)
	if err != nil {
		log.Err(err).Str("actionName", actionName).Msg("Unable to get action.")
		return nil
	}

	if actionStr == "" {
		log.Err(err).Str("actionName", actionName).Msg("Action string empty.")
		return nil
	}

	var action common.ConnectorAction

	err = json.Unmarshal([]byte(actionStr), &action)
	if err != nil {
		log.Err(err).Str("actionName", actionName).Msg("Unable unmarshal action.")
		return nil
	}

	return &action
}
