package connectors

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/x-sushant-x/connective/internal/connectors/common"
	"github.com/x-sushant-x/connective/internal/core/domain"
	"github.com/x-sushant-x/connective/internal/core/port"
)

type OAuthHandler struct {
	integrationSvc          port.IntegrationService
	providerRepo            port.ProviderRepo
	connectedAccountRepo    port.ConnectedAccountRepo
	projectRepo             port.ProjectRepo
	providerCredentialsRepo port.ProviderCredentialsRepo
	connectorRegistry       *Registry
}

// oauthLogContext holds common context for OAuth logging
type oauthLogContext struct {
	provider   string
	projectID  string
	userID     string
	providerID int
}

func NewOAuthHandler(
	integrationSvc port.IntegrationService,
	providerRepo port.ProviderRepo,
	connectedAccountRepo port.ConnectedAccountRepo,
	projectRepo port.ProjectRepo,
	providerCredentialsRepo port.ProviderCredentialsRepo,
	connectorRegistry *Registry,
) *OAuthHandler {
	return &OAuthHandler{
		integrationSvc,
		providerRepo,
		connectedAccountRepo,
		projectRepo,
		providerCredentialsRepo,
		connectorRegistry,
	}
}

// HandleConnect godoc
// @Summary Initiate OAuth connection
// @Description Starts the OAuth flow for a specific provider by redirecting the user to the provider's authorization URL.
// @Tags OAuth
// @Accept json
// @Produce json
//
// @Param provider   query string true "Provider name (e.g., slack, GitHub)"
// @Param projectID  query string true "Project ID associated with the authorization"
// @Param userID     query string true "User ID initiating the connection"
//
// @Success 302 {string} string "Redirects to provider OAuth URL"
// @Failure 404 {object} map[string]string "Unknown or unsupported provider"
// @Failure 500 {object} map[string]string "Missing provider credentials"
//
// @Router /oauth/connect [get]
func (h *OAuthHandler) HandleConnect(ctx *gin.Context) {
	providerName := ctx.Query("provider")
	projectID := ctx.Query("projectID")
	projectSecret := ctx.Query("projectSecret")
	userId := ctx.Query("userID")

	logCtx := oauthLogContext{provider: providerName, projectID: projectID, userID: "PENDING_RESOLUTION"}
	log.Info().Str("provider", providerName).Str("projectID", projectID).Str("userID", "PENDING_RESOLUTION").Msg("OAuth connect initiated")

	if !h.validateConnectParams(ctx, providerName, projectID, projectSecret, userId) {
		return
	}

	project, err := h.projectRepo.GetByID(ctx, projectID)
	if err != nil || project == nil {
		logError("Failed to retrieve project", logCtx, err)
		ctx.JSON(400, gin.H{"error": "Unable to connect with " + providerName, "error_code": "INVALID_PROJECT"})
		return
	}

	if projectSecret != project.SDKAuthSecret {
		ctx.JSON(400, gin.H{"error": "Unable to connect with " + providerName, "error_code": "INVALID_PROJECT_SECRET"})
		return
	}

	connector := h.connectorRegistry.GetConnector(providerName)
	if connector == nil {
		logError("Unknown provider connector", logCtx)
		ctx.JSON(400, gin.H{"error": "Unable to connect with " + providerName, "error_code": "UNKNOWN_PROVIDER"})
		return
	}

	provider, err := h.providerRepo.GetProviderByName(ctx, providerName)
	if err != nil {
		logError("Failed to retrieve provider", logCtx, err)
		ctx.JSON(400, gin.H{"error": "Unable to connect with " + providerName, "error_code": "UNKNOWN_PROVIDER"})
		return
	}

	if provider == nil {
		logError("Provider not found in database", logCtx)
		ctx.JSON(400, gin.H{"error": "Unable to connect with " + providerName, "error_code": "UNKNOWN_PROVIDER"})
		return
	}

	auth := connector.AuthStrategy()
	logCtx.providerID = provider.ID

	creds, err := h.integrationSvc.GetProviderCreds(ctx, projectID, provider.ID)
	if err != nil {
		logError("Failed to retrieve integration credentials", logCtx, err)
		ctx.JSON(400, gin.H{"error": "Unable to connect with " + providerName, "error_code": "MISSING_CREDENTIALS"})
		return
	}

	if creds == nil {
		logError("No provider credentials found", logCtx)
		ctx.JSON(400, gin.H{"error": "Unable to connect with " + providerName, "error_code": "MISSING_CREDENTIALS"})
		return
	}

	u := auth.BuildAuthURL(creds, projectID, userId)

	log.Info().Str("provider", providerName).Str("projectID", projectID).Str("userID", userId).Str("authURL", auth.AuthURL()).Msg("Redirecting to OAuth provider")
	ctx.Redirect(http.StatusTemporaryRedirect, u)
}

func (h *OAuthHandler) HandleCallback(ctx *gin.Context) {
	providerName := ctx.Param("provider")
	code := ctx.Query("code")
	state := ctx.Query("state")

	log.Info().
		Str("provider", providerName).
		Str("code", maskCode(code)).
		Str("state", maskState(state)).
		Msg("OAuth callback received")

	if !h.validateCallbackParams(ctx, providerName, code, state) {
		return
	}

	parts := strings.Split(state, ":")
	if len(parts) != 2 {
		log.Error().Str("provider", providerName).Str("state", maskState(state)).Int("partsCount", len(parts)).Msg("Invalid state format")
		ctx.JSON(400, gin.H{"error": "Unable to connect with " + providerName, "error_code": "INVALID_STATE"})
		return
	}

	projectID, userId := parts[0], parts[1]
	logCtx := oauthLogContext{provider: providerName, projectID: projectID, userID: userId}

	connector := h.connectorRegistry.GetConnector(providerName)
	if connector == nil {
		logError("Unknown provider connector", logCtx)
		ctx.JSON(400, gin.H{"error": "Unable to connect with " + providerName, "error_code": "UNKNOWN_PROVIDER"})
		return
	}

	provider, err := h.providerRepo.GetProviderByName(ctx, providerName)
	if err != nil {
		logError("Failed to retrieve provider", logCtx, err)
		ctx.JSON(400, gin.H{"error": "Unable to connect with " + providerName, "error_code": "INTERNAL_SERVER_ERROR"})
		return
	}

	if provider == nil {
		logError("Provider not found in database", logCtx)
		ctx.JSON(400, gin.H{"error": "Unable to connect with " + providerName, "error_code": "INTERNAL_SERVER_ERROR"})
		return
	}

	logCtx.providerID = provider.ID

	creds, err := h.integrationSvc.GetProviderCreds(ctx, projectID, provider.ID)
	if err != nil {
		logError("Failed to retrieve integration credentials", logCtx, err)
		ctx.JSON(400, gin.H{"error": "Unable to connect with " + providerName, "error_code": "MISSING_CREDENTIALS"})
		return
	}

	if creds == nil {
		logError("No provider credentials found", logCtx)
		ctx.JSON(400, gin.H{"error": "Unable to connect with " + providerName, "error_code": "MISSING_CREDENTIALS"})
		return
	}

	token, rawResponse, err := h.exchangeCode(connector, creds, code)
	if err != nil {
		logError("Failed to exchange authorization code", logCtx, err)
		ctx.JSON(400, gin.H{"error": "Unable to connect with " + providerName, "error_code": "OAUTH_CLIENT_ERROR"})
		return
	}

	log.Info().Str("provider", providerName).Str("projectID", projectID).Str("userID", userId).Bool("hasRefreshToken", token.RefreshToken != "").Msg("Token exchange successful")

	existingAcc, err := h.connectedAccountRepo.GetByProjectAndProviderAndUserId(ctx, projectID, userId, provider.ID)

	if existingAcc != nil {
		existingAcc.UpdatedAt = time.Now()
		existingAcc.RawResponse = rawResponse
		existingAcc.AccessToken = token.AccessToken
		existingAcc.Scope = strings.Join(creds.Scopes, ",")

		if token.ExpiresIn > 0 {
			expireTime := time.Now().UTC().Add(time.Duration(token.ExpiresIn) * time.Second)
			existingAcc.ExpiresAt = &expireTime
		}

		_, err := h.connectedAccountRepo.Update(ctx, existingAcc)
		if err != nil {
			logError("Failed to update connected account", logCtx, err)
			ctx.JSON(400, gin.H{"error": "Unable to connect with " + providerName, "error_code": "INTERNAL_SERVER_ERROR"})
			return
		}
	} else {
		connectedAccount := &domain.ConnectedAccount{
			ID:           uuid.NewString(),
			ProjectID:    projectID,
			ProviderID:   provider.ID,
			UserId:       userId,
			AccessToken:  token.AccessToken,
			RefreshToken: token.RefreshToken,
			TokenType:    "oAuth2",
			Scope:        strings.Join(creds.Scopes, ","),
			ConnectedAt:  time.Now(),
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			RawResponse:  rawResponse,
		}

		if token.ExpiresIn > 0 {
			expireTime := time.Now().UTC().Add(time.Duration(token.ExpiresIn) * time.Second)
			connectedAccount.ExpiresAt = &expireTime
		}

		_, err = h.connectedAccountRepo.Create(ctx, connectedAccount)
		if err != nil {
			logError("Failed to create connected account", logCtx, err)
			ctx.JSON(400, gin.H{"error": "Unable to connect with " + providerName, "error_code": "INTERNAL_SERVER_ERROR"})
			return
		}
	}

	log.Info().Str("provider", providerName).Str("projectID", projectID).Str("userID", userId).Int("providerID", provider.ID).Msg("Connected account created successfully")
	ctx.Data(200, "text/html; charset=utf-8", []byte(`
		<html>
  			<body>
    			<script>
      				if (window.opener) {
        				window.opener.postMessage({ success: true, provider: "`+providerName+`" }, "*");
      				}

      				// Close the popup automatically
      				window.close();
    			</script>
  			</body>
		</html>
		`))
}

/*
Returns:
- *TokenResponse
- []byte - raw response from the provider
- error
*/
func (h *OAuthHandler) exchangeCode(connector common.Connector, credentials *domain.ProviderCredentials, code string) (*common.TokenResponse, []byte, error) {
	auth := connector.AuthStrategy()
	tokenURL := auth.TokenURL()
	clientID := credentials.ClientID

	log.Info().Str("tokenURL", tokenURL).Str("clientID", maskClientID(clientID)).Msg("Exchanging code for token")

	data := url.Values{}
	auth.PrepareTokenRequest(credentials, code, data)

	resp, err := http.PostForm(tokenURL, data)
	if err != nil {
		log.Error().Err(err).Str("tokenURL", tokenURL).Msg("HTTP request failed")
		return nil, nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error().Err(err).Str("tokenURL", tokenURL).Int("statusCode", resp.StatusCode).Msg("Failed to read response body")
		return nil, nil, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Error().Err(err).Str("tokenURL", tokenURL).Str("response", string(body)).Msg("Non 200 status code.")
		return nil, nil, errors.New("unable to exchange token")
	}

	token, err := auth.ParseTokenResponse(body)
	if err != nil {
		log.Error().Err(err).Str("body", string(body)).Int("statusCode", resp.StatusCode).Msg("Failed to parse token")
		return nil, nil, err
	}

	log.Info().Str("tokenURL", tokenURL).Bool("hasRefreshToken", token.RefreshToken != "").Msg("Token exchange successful")
	return token, body, nil
}

func (h *OAuthHandler) getNewAccessToken(ctx context.Context, connector common.Connector, provider *domain.Provider, projectID string, connAcc *domain.ConnectedAccount) (*common.TokenResponse, error) {
	auth := connector.AuthStrategy()

	creds, err := h.providerCredentialsRepo.GetByProjectAndProvider(ctx, projectID, provider.ID)
	if err != nil || creds == nil {
		return nil, err
	}

	data := url.Values{}
	data.Set("client_id", creds.ClientID)
	data.Set("client_secret", creds.ClientSecret)
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", connAcc.RefreshToken)

	resp, err := http.PostForm(auth.RefreshTokenURL(), data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var token common.TokenResponse
	if err := json.Unmarshal(bodyBytes, &token); err != nil {
		return nil, err
	}

	return &token, nil
}

// Helper function for consistent error logging
func logError(msg string, logCtx oauthLogContext, err ...error) {
	event := log.Error()
	if len(err) > 0 && err[0] != nil {
		event = event.Err(err[0])
	}
	if logCtx.provider != "" {
		event = event.Str("provider", logCtx.provider)
	}
	if logCtx.projectID != "" {
		event = event.Str("projectID", logCtx.projectID)
	}
	if logCtx.userID != "" {
		event = event.Str("userID", logCtx.userID)
	}
	if logCtx.providerID != 0 {
		event = event.Int("providerID", logCtx.providerID)
	}
	event.Msg(msg)
}

// Validation helpers
func (h *OAuthHandler) validateConnectParams(ctx *gin.Context, provider, projectID, projectSecret, userID string) bool {
	if provider == "" {
		log.Warn().Msg("Missing provider parameter")
		ctx.JSON(400, gin.H{"error": "provider parameter is required"})
		return false
	}

	if projectID == "" {
		log.Warn().Msg("Missing projectID parameter")
		ctx.JSON(400, gin.H{"error": "projectID parameter is required"})
		return false
	}

	if projectSecret == "" {
		log.Warn().Msg("Missing projectSecret parameter")
		ctx.JSON(400, gin.H{"error": "projectSecret parameter is required"})
		return false
	}

	if userID == "" {
		log.Warn().Msg("Missing userId parameter")
		ctx.JSON(400, gin.H{"error": "userId parameter is required"})
		return false
	}
	return true
}

func (h *OAuthHandler) validateCallbackParams(ctx *gin.Context, provider, code, state string) bool {
	if provider == "" {
		log.Warn().Msg("Missing provider parameter")
		ctx.JSON(400, gin.H{"error": "provider parameter is required"})
		return false
	}
	if code == "" {
		log.Warn().Str("provider", provider).Msg("Missing authorization code")
		ctx.JSON(400, gin.H{"error": "authorization code is required"})
		return false
	}
	if state == "" {
		log.Warn().Str("provider", provider).Msg("Missing state parameter")
		ctx.JSON(400, gin.H{"error": "state parameter is required"})
		return false
	}
	return true
}
