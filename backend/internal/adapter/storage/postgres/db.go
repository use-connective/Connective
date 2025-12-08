package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
	"github.com/x-sushant-x/connective/internal/adapter/config"
)

type DB struct {
	*pgxpool.Pool
}

func New(ctx context.Context) (*DB, error) {
	conn, err := pgxpool.New(ctx, buildDsn())
	if err != nil {
		return nil, err
	}

	err = conn.Ping(ctx)
	if err != nil {
		return nil, err
	}

	log.Info().Msg("Connected to PostgresDB.")
	return &DB{
		conn,
	}, nil
}

func buildDsn() string {
	c := config.Config

	return fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
		c.Database,
		c.DatabaseUser,
		c.DatabasePassword,
		c.DatabaseHost,
		c.DatabasePort,
		c.DatabaseName,
	)
}
