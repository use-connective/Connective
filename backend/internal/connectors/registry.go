package connectors

import (
	"github.com/rs/zerolog/log"
	"github.com/x-sushant-x/connective/internal/connectors/common"
)

var Registry = map[string]common.Connector{}

func RegisterConnector(connectors ...common.Connector) {
	for _, connector := range connectors {
		Registry[connector.Name()] = connector
		log.Info().Str("name", connector.Name()).Msg("Connector Registered")
	}
}

func GetConnector(name string) common.Connector {
	return Registry[name]
}
