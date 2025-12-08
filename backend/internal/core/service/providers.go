package service

import (
	"context"
	"encoding/json"
	"os"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/x-sushant-x/connective/internal/core/domain"
	cErrors "github.com/x-sushant-x/connective/internal/core/errors"
	"github.com/x-sushant-x/connective/internal/core/port"
)

type ProviderService struct {
	repo port.ProviderRepo
}

func NewProviderSvc(repo port.ProviderRepo) *ProviderService {
	return &ProviderService{
		repo,
	}
}

func (ps *ProviderService) GetByID(ctx context.Context, id int) (*domain.Provider, error) {
	provider, err := ps.repo.GetByID(ctx, id)
	if err != nil {
		log.Err(err).Int("id", id).Msg("unable to get provider")
		return nil, cErrors.ErrUnableToGetProvider
	}

	return provider, nil
}

func (ps *ProviderService) SeedOnStartup(ctx context.Context) {
	env := os.Getenv("ENV")
	if env == "" {
		log.Fatal().Msg("ENV variable does not exists in .env file")
	}

	var fileName string

	if strings.EqualFold(fileName, "Prod") {
		fileName = "providers.json"
	} else {
		fileName = "providers.local.json"
	}

	providers := make([]domain.Provider, 0)

	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to read providers file")
	}

	if err := json.Unmarshal(data, &providers); err != nil {
		log.Fatal().Err(err).Msg("unable to unmarshal providers")

	}

	for _, p := range providers {

		provider, err := ps.repo.GetProviderByName(ctx, p.Name)
		if err != nil {
			log.Fatal().Err(err).Str("name", p.Name).Msg("unable to seed provider")
		}

		if provider == nil {
			_, err := ps.repo.Create(ctx, &p)
			if err != nil {
				log.Fatal().Err(err).Str("name", p.Name).Msg("unable to save provider")
			}
			log.Info().Str("name", p.Name).Msg("Provider Seeded")
		}
	}
}

func (ps *ProviderService) GetProviderByName(ctx context.Context, name string) (*domain.Provider, error) {
	provider, err := ps.repo.GetProviderByName(ctx, name)

	if err != nil {
		log.Err(err).Str("name", name).Msg("unable to get provider")
		return nil, cErrors.ErrUnableToGetProvider
	}

	return provider, nil
}
