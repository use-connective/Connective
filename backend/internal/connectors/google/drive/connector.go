package googleDrive

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/x-sushant-x/connective/internal/connectors/common"
	"github.com/x-sushant-x/connective/internal/core/port"
)

type Connector struct {
	providerRepo port.ProviderRepo
	cache        port.Cache
}

func New(providerRepo port.ProviderRepo, cache port.Cache) *Connector {
	return &Connector{
		providerRepo,
		cache,
	}
}

func (s *Connector) Name() string {
	return "google-drive"
}

func (s *Connector) AuthStrategy() common.AuthStrategy {
	ctx := context.Background()

	provider, err := s.providerRepo.GetProviderByName(ctx, "google-drive")
	if err != nil || provider == nil {
		// TODO - See where it impacts and handle accordingly.
		log.Err(err).Msg("Unable to get auth strategy for google drive")
		return nil
	}

	return NewDriveStrategy(provider)
}

func (s *Connector) GetAction(ctx context.Context, actionName string) *common.ConnectorAction {
	actionCacheKey := fmt.Sprintf("%s_%s", s.Name(), actionName)

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
