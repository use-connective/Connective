package googleDrive

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/x-sushant-x/connective/internal/connectors/common"
	"github.com/x-sushant-x/connective/internal/core/port"
)

type Connector struct {
	providerRepo port.ProviderRepo
}

func New(providerRepo port.ProviderRepo) *Connector {
	return &Connector{
		providerRepo,
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

func (s *Connector) Actions() map[string]common.ActionHandler {
	return map[string]common.ActionHandler{}
}
