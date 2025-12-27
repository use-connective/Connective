/*
	Author: Sushant <sushant.dhiman9812@gmail.com>

	Purpose:
	* Entry point for Connective backend system.
	* Connects to databases.
	* Setup API endpoints.
	* Register connectors and actions into system and provide functionality for graceful shutdown.

	IMPORTANT to know:

	* You must know how oAuth works before going through this source code.

	* This project is built using Hexagonal Architecture so having its knowledge is also recommended.

	* Throughout this project (backend and dashboard) source code words connector, provider, integration means the same thing.
      It is 3rd party product that we are connecting to such as Drive, Notion, Slack, GitHub etc.

	* Actions - These are the actions which can be performed on 3rd party product
	  such as sending a message in slack workspace or getting user files from dropbox.
*/

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
	googleCalendar "github.com/x-sushant-x/connective/internal/connectors/google/calendar"
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

	// Load config from either config.json or config.local.json
	// based on environment and initialize into Config global variable
	// present in config.go file.
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

	// Creates a new redis connection and return implementation of
	// port.Cache interface
	cache := redisClient.NewRedisClient(ctx)

	// Connector Registry is a registry that holds a list of connectors. These connectors are registered
	// on startup of the system. It can be used to register and retrieve connectors.
	// It can be used to get action that can be performed by this connector. (Used in action_handler.go)
	connectorRegistry := connectors.NewRegistryHandler(cache)

	// Repository Layer Initialization - All these implements repo interface from port folder.
	authRepo := postgres.NewAuthRepo(pg)
	projectRepo := postgres.NewProjectRepo(pg)
	providerCredentialsRepo := postgres.NewProviderCredentialsRepo(pg)
	providerRepo := postgres.NewProviderRepo(pg)
	connectedAccountRepo := postgres.NewConnectedAccountRepo(pg)
	userRepo := postgres.NewUserRepo(pg)

	router := server.NewRouter(authRepo, projectRepo, providerCredentialsRepo, providerRepo, connectedAccountRepo, userRepo, connectorRegistry)

	// Providers (connectors) are stored in providers.local.json or providers.json file (as per environment).
	// Seeding means we are storing these in database (if not already stored)
	providerSvc := service.NewProviderSvc(providerRepo)
	providerSvc.SeedOnStartup(ctx)

	httpAddr := fmt.Sprintf(":%s", config.Config.BackendPort)

	srv := &http.Server{
		Addr:    httpAddr,
		Handler: router,
	}

	registerConnectors(ctx, providerRepo, connectorRegistry, cache)

	go func() {
		log.Info().Msgf("HTTP Server started on port: %s", httpAddr)

		if err := srv.ListenAndServe(); err != nil {
			log.Fatal().Err(err).Msg("HTTP Server crashed")

		}
	}()

	// Graceful shutdown.
	// We must close opened connections before exiting application.
	quit := make(chan os.Signal, 1)

	// signal.Notify will listen for SIGTERM signal (CTRL+C) and send that to quit channel.
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	// All code below this line will only be executed if CTRL+C is pressed on server.
	log.Info().Msg("Shutting down server...")

	ctxShutdown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctxShutdown); err != nil {
		log.Error().Err(err).Msg("Forced shutdown")
	}

	log.Info().Msg("Server exited cleanly")
}

func registerConnectors(ctx context.Context, providerRepo port.ProviderRepo, registry *connectors.Registry, cache port.Cache) {
	slack := iSlack.New(providerRepo, cache)
	gCalendar := googleCalendar.New(providerRepo, cache)
	github := githubConnector.New(providerRepo, cache)
	dropbox := dropboxConnector.New(providerRepo, cache)
	drive := googleDrive.New(providerRepo, cache)
	gmail := googleMail.New(providerRepo, cache)

	registry.Register(ctx, slack, gCalendar, github, dropbox, drive, gmail)
}
