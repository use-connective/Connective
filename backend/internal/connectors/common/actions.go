package common

import "context"

type ActionHandler func(ctx context.Context, creds *UserCredentials, payload map[string]any) (any, error)
