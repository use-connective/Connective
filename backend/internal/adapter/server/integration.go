package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/x-sushant-x/connective/internal/core/dto"
	"github.com/x-sushant-x/connective/internal/core/port"
)

type IntegrationHandler struct {
	svc         port.IntegrationService
	providerSvc port.ProviderService
}

func NewIntegrationHandler(svc port.IntegrationService, providerSvc port.ProviderService) *IntegrationHandler {
	return &IntegrationHandler{
		svc,
		providerSvc,
	}
}

// SaveProviderCreds
// @Summary Create or update integration credential.
// @Security BearerAuth
// @Description Create or update integration credential.
// @Tags Integration
// @Accept  json
// @Produce  json
// @Param  request body dto.CreateUpdateProviderCreds true "Create Project Body"
// @Success 200 {object} APIResponse{data=any}
// @Router /api/v1/integration/creds/save [post]
func (ih *IntegrationHandler) SaveProviderCreds(ctx *gin.Context) {
	var req dto.CreateUpdateProviderCreds

	if err := ctx.BindJSON(&req); err != nil {
		BadRequest(ctx, "Bad Request")
		return
	}

	resp, err := ih.svc.SaveProviderCreds(ctx, &req)
	if err != nil {
		BadRequest(ctx, err.Error())
		return
	}

	Success(ctx, resp, "Credentials Saved", http.StatusOK)
}

// GetAllProvidersList
// @Summary Get all providers list.
// @Security BearerAuth
// @Description Get all providers list.
// @Tags Integration
// @Accept  json
// @Produce  json
// @Param  search query string false "Search"
// @Param  category query string false "Category"
// @Success 200 {object} APIResponse{data=dao.ProviderListDisplayable}
// @Router /api/v1/integration/providers [get]
func (ih *IntegrationHandler) GetAllProvidersList(ctx *gin.Context) {
	search := ctx.Query("search")
	category := ctx.Query("category")

	if category == "All" {
		category = ""
	}

	resp, err := ih.svc.GetProviderListDisplayable(ctx, search, category)
	if err != nil {
		BadRequest(ctx, err.Error())
		return
	}

	Success(ctx, resp, "Providers Retrieved", http.StatusOK)
}

// GetProviderByID
// @Summary Get provider by id.
// @Security BearerAuth
// @Description Get provider by id.
// @Tags Integration
// @Accept  json
// @Produce  json
// @Param  id query string true "Provider Id"
// @Success 200 {object} APIResponse{data=domain.Provider}
// @Router /api/v1/integration/provider [get]
func (ih *IntegrationHandler) GetProviderByID(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		BadRequest(ctx, "Bad Request")
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		BadRequest(ctx, "Invalid provider id.")
		return
	}

	resp, err := ih.providerSvc.GetByID(ctx, idInt)
	if err != nil {
		BadRequest(ctx, err.Error())
		return
	}

	Success(ctx, resp, "Provider Retrieved", http.StatusOK)
}

// GetProviderCreds
// @Summary Get provider credentials.
// @Description Get provider credentials.
// @Tags Integration
// @Accept  json
// @Produce  json
// @Param  providerID query string true "Provider ID"
// @Param  projectID query string true "Project ID"
// @Success 200 {object} APIResponse{data=domain.Provider}
// @Router /api/v1/integration/provider/creds [get]
func (ih *IntegrationHandler) GetProviderCreds(ctx *gin.Context) {
	projectID := ctx.Query("projectID")
	providerID := ctx.Query("providerID")

	if projectID == "" || providerID == "" {
		BadRequest(ctx, "providerID and projectID must be provided")
		return
	}

	providerIDInt, err := strconv.Atoi(providerID)
	if err != nil {
		BadRequest(ctx, "providerID must be an integer")
		return
	}

	creds, err := ih.svc.GetProviderCreds(ctx, projectID, providerIDInt)
	if err != nil {
		BadRequest(ctx, err.Error())
		return
	}

	Success(ctx, creds, "Credentials Reterieved", http.StatusOK)
}

// GetProviderByName
// @Summary Get provider by name.
// @Description Get provider by name.
// @Tags Integration
// @Accept  json
// @Produce  json
// @Param  provider query string true "Provider"
// @Success 200 {object} APIResponse{data=domain.Provider}
// @Router /api/v1/integration/provider/get-provider [get]
func (ih *IntegrationHandler) GetProviderByName(ctx *gin.Context) {
	provider := ctx.Query("provider")

	if provider == "" {
		BadRequest(ctx, "provider and projectID must be provided")
		return
	}

	p, err := ih.providerSvc.GetProviderByName(ctx, provider)
	if err != nil {
		BadRequest(ctx, err.Error())
		return
	}

	if p == nil {
		BadRequest(ctx, "Invalid Integration Name")
		return
	}

	Success(ctx, p, "Provider Fetched", http.StatusOK)
}

// GetAllCategories
// @Summary Get all categories.
// @Description Get all categories.
// @Tags Integration
// @Accept  json
// @Produce  json
// @Router /api/v1/integration/provider/categories [get]
func (ih *IntegrationHandler) GetAllCategories(ctx *gin.Context) {
	Success(ctx, ih.svc.GetCategories(ctx), "Categories Fetched", http.StatusOK)
}

// GetConnectedAccounts
// @Summary Get connected accounts.
// @Description Get connected accounts.
// @Tags Integration
// @Accept  json
// @Produce  json
// @Param  projectId query string true "Project I'd"
// @Param  userId query string true "User I'd"
// @Router /api/v1/integration/connected-accounts [get]
func (ih *IntegrationHandler) GetConnectedAccounts(ctx *gin.Context) {
	projectID := ctx.Query("projectId")
	userID := ctx.Query("userId")

	connectedAccounts, err := ih.svc.GetConnectedUsers(ctx, projectID, userID)
	if err != nil {
		BadRequest(ctx, err.Error())
		return
	}

	Success(ctx, connectedAccounts, "Connected Accounts Fetched", http.StatusOK)
}
