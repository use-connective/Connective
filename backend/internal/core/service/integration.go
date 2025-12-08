package service

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/x-sushant-x/connective/internal/core/dao"
	"github.com/x-sushant-x/connective/internal/core/domain"
	"github.com/x-sushant-x/connective/internal/core/dto"
	cErrors "github.com/x-sushant-x/connective/internal/core/errors"
	"github.com/x-sushant-x/connective/internal/core/port"
)

type IntegrationService struct {
	providerCredsRepo     port.ProviderCredentialsRepo
	providerRepo          port.ProviderRepo
	connectedAccountsRepo port.ConnectedAccountRepo
}

func NewIntegrationService(providerCredsRepo port.ProviderCredentialsRepo, providerRepo port.ProviderRepo, connectedAccountsRepo port.ConnectedAccountRepo) *IntegrationService {
	return &IntegrationService{
		providerCredsRepo,
		providerRepo,
		connectedAccountsRepo,
	}
}

func (is *IntegrationService) SaveProviderCreds(ctx context.Context, req *dto.CreateUpdateProviderCreds) (*domain.ProviderCredentials, error) {
	existing, err := is.providerCredsRepo.GetByProjectAndProvider(ctx, req.ProjectID, req.ProviderID)
	if err != nil {
		log.Error().Err(err).
			Str("project_id", req.ProjectID).
			Int("provider_id", req.ProviderID).
			Msg("failed to fetch provider credentials")

		return nil, cErrors.ErrUnableToSaveCredentials
	}

	if existing != nil {
		now := time.Now()

		existing.ClientID = req.ClientID
		existing.ClientSecret = req.ClientSecret
		existing.Scopes = req.Scopes
		existing.UpdatedAt = now

		updated, err := is.providerCredsRepo.Update(ctx, existing)
		if err != nil {
			log.Error().Err(err).
				Str("project_id", req.ProjectID).
				Int("provider_id", req.ProviderID).
				Msg("failed to update provider credentials")

			return nil, cErrors.ErrUnableToSaveCredentials
		}

		return updated, nil
	}

	creds := domain.ProviderCredentials{
		ProjectID:    req.ProjectID,
		ProviderID:   req.ProviderID,
		ClientID:     req.ClientID,
		ClientSecret: req.ClientSecret,
		Scopes:       req.Scopes,
		IsActive:     true,
	}

	created, err := is.providerCredsRepo.Create(ctx, &creds)
	if err != nil {
		log.Error().Err(err).
			Str("project_id", req.ProjectID).
			Int("provider_id", req.ProviderID).
			Msg("failed to create provider credentials")

		return nil, cErrors.ErrUnableToSaveCredentials
	}

	return created, nil
}

func (is *IntegrationService) GetProviderCreds(ctx context.Context, projectID string, providerID int) (*domain.ProviderCredentials, error) {
	creds, err := is.providerCredsRepo.GetByProjectAndProvider(ctx, projectID, providerID)
	if err != nil {
		log.Err(err).Msg("unable to get providers credentials.")
		return nil, cErrors.ErrUnableToGetProviderCreds
	}

	return creds, nil
}

func (is *IntegrationService) GetProviderListDisplayable(ctx context.Context, search, category string) ([]dao.ProviderListDisplayable, error) {
	providers, err := is.providerRepo.GetProviderListDisplayable(ctx, search, category)

	if err != nil {
		log.Err(err).Msg("unable to get providers list with category")
		return nil, cErrors.ErrUnableToGetProvidersList
	}

	return providers, nil
}

func (is *IntegrationService) GetCategories(ctx context.Context) []string {
	return []string{
		"All",
		"CRM",
		"Sales",
		"Marketing",
		"Automation",
		"Analytics",
		"Advertising",
		"E-Commerce",
		"Communication",
		"Calendar",
		"Office",
		"Storage",
		"Project Management",
		"Support & Ticketing",
		"Payments",
		"HR",
		"Business Intelligence",
		"Social Media",
	}
}

func (is *IntegrationService) GetConnectedUsers(ctx context.Context, projectId, userID string) ([]dao.ConnectedUsers, error) {
	connectedAccounts, err := is.connectedAccountsRepo.GetConnectedUsers(ctx, projectId, userID)
	if err != nil {
		log.Err(err).Msg("unable to get connected accounts")
		return nil, cErrors.ErrUnableToGetConnectedAccount
	}

	return connectedAccounts, nil
}
