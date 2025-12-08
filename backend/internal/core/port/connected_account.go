package port

import (
	"context"

	"github.com/x-sushant-x/connective/internal/core/dao"
	"github.com/x-sushant-x/connective/internal/core/domain"
)

type ConnectedAccountRepo interface {
	Create(ctx context.Context, ca *domain.ConnectedAccount) (*domain.ConnectedAccount, error)
	Update(ctx context.Context, ca *domain.ConnectedAccount) (*domain.ConnectedAccount, error)
	Delete(ctx context.Context, id string) error
	GetByID(ctx context.Context, id string) (*domain.ConnectedAccount, error)
	GetByProjectAndProvider(ctx context.Context, projectID string, providerID int) (*domain.ConnectedAccount, error)
	GetByProjectAndProviderAndUserId(ctx context.Context, projectID, userID string, providerID int) (*domain.ConnectedAccount, error)
	GetConnectedUsers(ctx context.Context, projectId, search string) ([]dao.ConnectedUsers, error)
}
