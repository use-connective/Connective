package port

import (
	"context"
	"time"
)

type Cache interface {
	Set(ctx context.Context, key string, val any, expiry time.Duration) error
	Get(ctx context.Context, key string) (string, error)

	SetJson(ctx context.Context, key string, val any, expiry time.Duration) error
	GetJson(ctx context.Context, key string) (string, error)
}
