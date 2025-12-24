package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/x-sushant-x/connective/internal/adapter/config"
	"github.com/x-sushant-x/connective/internal/adapter/server"
	"github.com/x-sushant-x/connective/internal/adapter/storage/postgres"
	redisClient "github.com/x-sushant-x/connective/internal/adapter/storage/redis"
	"github.com/x-sushant-x/connective/internal/connectors"
	dropboxConnector "github.com/x-sushant-x/connective/internal/connectors/dropbox"
	githubConnector "github.com/x-sushant-x/connective/internal/connectors/github"
	"github.com/x-sushant-x/connective/internal/connectors/google/calendar"
	googleDrive "github.com/x-sushant-x/connective/internal/connectors/google/drive"
	googleMail "github.com/x-sushant-x/connective/internal/connectors/google/gmail"
	iSlack "github.com/x-sushant-x/connective/internal/connectors/slack"
	"github.com/x-sushant-x/connective/internal/core/port"
	"github.com/x-sushant-x/connective/internal/core/service"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	if err := godotenv.Load(); err != nil {
		log.Fatal().Err(err).Msg("unable to load environment variables.")
	}

	config.LoadConfiguration()

}

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	ctx := context.Background()

	// Database Connection & Auto Migration

	pg, err := postgres.New(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to connect to database.")
	}

	postgres.AutoMigrateTables(ctx, pg.Pool)

	cache := redisClient.NewRedisClient(ctx)

	connectorRegistry := connectors.NewRegistryHandler(cache)

	// Repository Layer Initialization

	authRepo := postgres.NewAuthRepo(pg)
	projectRepo := postgres.NewProjectRepo(pg)
	providerCredentialsRepo := postgres.NewProviderCredentialsRepo(pg)
	providerRepo := postgres.NewProviderRepo(pg)
	connectedAccountRepo := postgres.NewConnectedAccountRepo(pg)
	userRepo := postgres.NewUserRepo(pg)
	router := server.NewRouter(authRepo, projectRepo, providerCredentialsRepo, providerRepo, connectedAccountRepo, userRepo, connectorRegistry)

	// Seeding Providers
	providerSvc := service.NewProviderSvc(providerRepo)
	providerSvc.SeedOnStartup(ctx)

	httpAddr := fmt.Sprintf(":%s", config.Config.BackendPort)

	srv := &http.Server{
		Addr:    httpAddr,
		Handler: router,
	}

	registerConnectors(ctx, providerRepo, connectorRegistry)

	go func() {
		log.Info().Msgf("HTTP Server started on port: %s", httpAddr)

		err := srv.ListenAndServe()

		if err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("HTTP Server crashed")
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Info().Msg("Shutting down server...")

	ctxShutdown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctxShutdown); err != nil {
		log.Error().Err(err).Msg("Forced shutdown")
	}

	log.Info().Msg("Server exited cleanly")
}

func registerConnectors(ctx context.Context, providerRepo port.ProviderRepo, registry *connectors.Registry) {
	slack := iSlack.New(providerRepo)
	gCalendar := googleCalendar.New(providerRepo)
	github := githubConnector.New(providerRepo)
	dropbox := dropboxConnector.New(providerRepo)
	drive := googleDrive.New(providerRepo)
	gmail := googleMail.New(providerRepo)

	registry.Register(ctx, slack, gCalendar, github, dropbox, drive, gmail)
}
