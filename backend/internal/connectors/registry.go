package connectors

import (
	"context"
	"encoding/json"
	"io"
	"os"
	"path"

	"github.com/rs/zerolog/log"
	"github.com/x-sushant-x/connective/internal/connectors/common"
	"github.com/x-sushant-x/connective/internal/core/port"
)

type Registry struct {
	connectors map[string]common.Connector
	cache      port.Cache
}

func NewRegistryHandler(cache port.Cache) *Registry {
	return &Registry{
		connectors: make(map[string]common.Connector),
		cache:      cache,
	}
}

func (r *Registry) Register(ctx context.Context, connectors ...common.Connector) {
	for _, connector := range connectors {
		name := connector.Name()

		r.connectors[name] = connector
		r.saveActionsMetaInCache(ctx, name)
		log.Info().Str("name", name).Msg("Connector Registered")
	}
}

func (r *Registry) GetConnector(name string) common.Connector {
	return r.connectors[name]
}

func parseActions(connectorName string) []common.ConnectorAction {
	actionsFilePath := path.Join(".", "config", "actions", connectorName+".json")

	actionsFile, err := os.Open(actionsFilePath)
	if err != nil {
		log.Fatal().Err(err).
			Str("connector", connectorName).
			Msg("unable to open actions file.")
		os.Exit(1)
	}

	actionsBytes, err := io.ReadAll(actionsFile)
	if err != nil {
		log.Fatal().Err(err).
			Str("connector", connectorName).
			Msg("unable read actions file.")
		os.Exit(1)
	}

	var actions []common.ConnectorAction

	err = json.Unmarshal(actionsBytes, &actions)
	if err != nil {
		log.Fatal().Err(err).Str("connector", connectorName).Msg("invalid json actions file.")
		os.Exit(1)
	}

	return actions
}

func (r *Registry) saveActionsMetaInCache(ctx context.Context, connectorName string) {
	actions := parseActions(connectorName)

	for _, action := range actions {
		err := r.cache.SetJson(ctx, action.CacheKey, action, 0)
		if err != nil {
			log.Fatal().Err(err).
				Str("actionKey", action.CacheKey).
				Msg("unable to save connector action to cache.")
		}

		log.Info().Str("connector", connectorName).Str("action", action.CacheKey).Msg("Action Registered")
	}
}
