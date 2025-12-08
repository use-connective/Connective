package port

import (
	"context"

	"github.com/x-sushant-x/connective/internal/core/dao"
	"github.com/x-sushant-x/connective/internal/core/domain"
)

type ProviderService interface {
	SeedOnStartup(ctx context.Context)
	GetByID(ctx context.Context, id int) (*domain.Provider, error)
	GetProviderByName(ctx context.Context, name string) (*domain.Provider, error)
}

type ProviderRepo interface {
	Create(ctx context.Context, provider *domain.Provider) (*domain.Provider, error)
	Update(ctx context.Context, provider *domain.Provider) (*domain.Provider, error)
	Delete(ctx context.Context, id int) error
	GetByID(ctx context.Context, id int) (*domain.Provider, error)
	GetAll(ctx context.Context) ([]*domain.Provider, error)
	GetProviderByName(ctx context.Context, name string) (*domain.Provider, error)
	GetProviderListDisplayable(ctx context.Context, search, category string) ([]dao.ProviderListDisplayable, error)
}
