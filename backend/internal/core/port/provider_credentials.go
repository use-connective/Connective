package port

import (
	"context"

	"github.com/x-sushant-x/connective/internal/core/domain"
)

type ProviderCredentialsRepo interface {
	Create(ctx context.Context, creds *domain.ProviderCredentials) (*domain.ProviderCredentials, error)
	Update(ctx context.Context, creds *domain.ProviderCredentials) (*domain.ProviderCredentials, error)
	Delete(ctx context.Context, id int) error
	GetByProjectAndProvider(ctx context.Context, projectID string, providerID int) (*domain.ProviderCredentials, error)
}
