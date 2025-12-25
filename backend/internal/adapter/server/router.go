package server

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/x-sushant-x/connective/docs"
	"github.com/x-sushant-x/connective/internal/adapter/config"
	"github.com/x-sushant-x/connective/internal/connectors"
	"github.com/x-sushant-x/connective/internal/core/port"
	"github.com/x-sushant-x/connective/internal/core/service"
)

type Router struct {
	*gin.Engine
}

func NewRouter(
	authRepo port.AuthRepo,
	projectRepo port.ProjectRepo,
	providerCredentialsRepo port.ProviderCredentialsRepo,
	providerRepo port.ProviderRepo,
	connectedAccountRepo port.ConnectedAccountRepo,
	userRepo port.UserRepo,
	connectorsRegistry *connectors.Registry,
) *Router {

	// if config.Config.Environment == "prod" {
	gin.SetMode(gin.ReleaseMode)
	// }

	middlewareHandler := NewMiddlewareHandler(authRepo)

	router := gin.New()

	corsConfig := cors.Config{
		AllowOrigins:     config.Config.CORSAllowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	router.Use(cors.New(corsConfig))

	router.Use(gin.Recovery())

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]string{
			"status": "UP",
		})
	})

	// Shared Services Definitions

	integrationSvc := service.NewIntegrationService(providerCredentialsRepo, providerRepo, connectedAccountRepo)
	userSvc := service.NewUserService(userRepo)
	providerSvc := service.NewProviderSvc(providerRepo)

	oAuthHandler := connectors.NewOAuthHandler(integrationSvc, providerRepo, connectedAccountRepo, projectRepo, providerCredentialsRepo, connectorsRegistry)

	setupSwaggerRoute(router)
	setupAuthRoute(router, authRepo, userRepo, middlewareHandler)
	setupProjectRoutes(router, middlewareHandler, projectRepo, userSvc)
	setupIntegrationsRoute(router, middlewareHandler, integrationSvc, providerSvc)
	setupConnectorRoutes(router, oAuthHandler)
	setupUserRoutes(router, userRepo, middlewareHandler)
	setupActionHandler(router, providerRepo, connectedAccountRepo, projectRepo, providerCredentialsRepo, connectorsRegistry)

	return &Router{
		router,
	}
}

func setupSwaggerRoute(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func setupAuthRoute(router *gin.Engine, repo port.AuthRepo, userRepo port.UserRepo, middleware *Middleware) {

	svc := service.NewAuthService(repo, userRepo)
	handler := NewAuthHandler(svc)

	authGroup := router.Group("/api/v1/user")

	authGroup.POST("/create", handler.HandleCreateUser)
	authGroup.POST("/login", handler.HandleLoginUser)

	authGroup.POST("/complete-onboarding", middleware.CheckAuth, handler.HandleCompleteOnboarding)
}

func setupProjectRoutes(router *gin.Engine, middleware *Middleware, repo port.ProjectRepo, userSvc port.UserService) {
	svc := service.NewProjectService(repo, userSvc)
	handler := NewProjectHandler(svc)

	projectGroup := router.Group("/api/v1/project")

	projectGroup.POST("/create", middleware.CheckAuth, handler.HandleCreateProject)
	projectGroup.GET("/get-all", middleware.CheckAuth, handler.HandleGetAllProjects)

}

func setupIntegrationsRoute(router *gin.Engine, middleware *Middleware, svc port.IntegrationService, providerSvc port.ProviderService) {
	h := NewIntegrationHandler(svc, providerSvc)
	integrationGroup := router.Group("/api/v1/integration")

	integrationGroup.POST("/creds/save", middleware.CheckAuth, h.SaveProviderCreds)
	integrationGroup.GET("/providers", middleware.CheckAuth, h.GetAllProvidersList)
	integrationGroup.GET("/provider", middleware.CheckAuth, h.GetProviderByID)
	integrationGroup.GET("/provider/creds", middleware.CheckAuth, h.GetProviderCreds)
	integrationGroup.GET("/provider/get-provider", middleware.CheckAuth, h.GetProviderByName)
	integrationGroup.GET("/provider/categories", middleware.CheckAuth, h.GetAllCategories)
	integrationGroup.GET("/connected-accounts", middleware.CheckAuth, h.GetConnectedAccounts)

}

func setupConnectorRoutes(
	router *gin.Engine,
	h *connectors.OAuthHandler) {

	router.GET("/oauth/connect", h.HandleConnect)
	router.GET("/oauth/callback/:provider", h.HandleCallback)
}

func setupUserRoutes(router *gin.Engine, userSvc port.UserService, middleware *Middleware) {
	h := NewUserHandler(userSvc)

	userGroup := router.Group("/api/v1/user")

	userGroup.GET("/", middleware.CheckAuth, h.HandleGetUser)
}

func setupActionHandler(
	router *gin.Engine,
	providerRepo port.ProviderRepo,
	connectedAccountRepo port.ConnectedAccountRepo,
	projectRepo port.ProjectRepo,
	providerCredentialsRepo port.ProviderCredentialsRepo,
	connectorRegistry *connectors.Registry) {

	h := connectors.NewConnectorHandler(providerRepo, connectedAccountRepo, projectRepo, providerCredentialsRepo, connectorRegistry)

	router.POST("/execute/:provider/:action", h.ExecuteAction)
}
