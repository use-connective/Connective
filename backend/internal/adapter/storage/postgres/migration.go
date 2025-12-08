package postgres

import (
	"context"
	"os"
	"path/filepath"
	"runtime"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

func AutoMigrateTables(ctx context.Context, db *pgxpool.Pool) {
	_, currentFile, _, _ := runtime.Caller(0)

	baseDir := filepath.Dir(currentFile)
	migrationsDir := filepath.Join(baseDir, "migrations")

	files, err := filepath.Glob(filepath.Join(migrationsDir, "*.sql"))
	if err != nil {
		log.Fatal().Err(err).Msg("failed to read migrations folder")
	}

	for _, filePath := range files {
		queryBytes, err := os.ReadFile(filePath)
		if err != nil {
			log.Fatal().Err(err).Str("file", filePath).Msg("failed to read migration")
		}

		_, err = db.Exec(ctx, string(queryBytes))
		if err != nil {
			log.Fatal().Err(err).Str("file", filePath).Msg("failed to execute migration")
		}
	}
}
