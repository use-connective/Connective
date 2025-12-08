package port

import (
	"context"

	"github.com/x-sushant-x/connective/internal/core/dao"
	"github.com/x-sushant-x/connective/internal/core/domain"
	"github.com/x-sushant-x/connective/internal/core/dto"
)

type IntegrationService interface {
	SaveProviderCreds(ctx context.Context, req *dto.CreateUpdateProviderCreds) (*domain.ProviderCredentials, error)
	GetProviderCreds(ctx context.Context, projectID string, providerID int) (*domain.ProviderCredentials, error)
	GetProviderListDisplayable(ctx context.Context, search, category string) ([]dao.ProviderListDisplayable, error)
	GetCategories(ctx context.Context) []string
	GetConnectedUsers(ctx context.Context, projectId, userId string) ([]dao.ConnectedUsers, error)
}

type IntegrationRepo interface {
	SaveProviderCreds(ctx context.Context, creds *domain.ProviderCredentials) (*domain.ProviderCredentials, error)
	GetByProviderCredsTypeAndProjectId(ctx context.Context, integrationType, projectId string) (*domain.ProviderCredentials, error)
}
