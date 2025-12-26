package connectors

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/x-sushant-x/connective/internal/connectors/common"
	"github.com/x-sushant-x/connective/internal/core/port"
)

type ConnectorHandler struct {
	providerRepo            port.ProviderRepo
	connectedAccountRepo    port.ConnectedAccountRepo
	projectRepo             port.ProjectRepo
	providerCredentialsRepo port.ProviderCredentialsRepo
	connectorRegistry       *Registry
}

func NewConnectorHandler(
	providerRepo port.ProviderRepo,
	connectedAccountRepo port.ConnectedAccountRepo,
	projectRepo port.ProjectRepo,
	providerCredentialsRepo port.ProviderCredentialsRepo,
	connectorRegistry *Registry) *ConnectorHandler {

	return &ConnectorHandler{
		providerRepo:            providerRepo,
		connectedAccountRepo:    connectedAccountRepo,
		projectRepo:             projectRepo,
		providerCredentialsRepo: providerCredentialsRepo,
		connectorRegistry:       connectorRegistry,
	}

}

// ExecuteAction godoc
// @Summary Execute Action Request
// @Description Execute Action Request
// @Tags Action
// @Accept  json
// @Produce  json
// @Param  request body common.ActionExecuteRequest true "Execute Action Request"
// @Router /execute/:provider/:action [post]
func (c *ConnectorHandler) ExecuteAction(ctx *gin.Context) {
	var req common.ActionExecuteRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": "Action execution failed.", "error_code": "INVALID_REQUEST"})
		return
	}

	projectID := req.ProjectID
	providerName := ctx.Param("provider")
	actionName := ctx.Param("action")

	project, err := c.projectRepo.GetByID(ctx, projectID)
	if err != nil || project == nil {
		ctx.JSON(400, gin.H{"error": "Action execution failed.", "error_code": "INVALID_PROJECT"})
		return
	}

	if req.ProjectSecret != project.SDKAuthSecret {
		ctx.JSON(400, gin.H{"error": "Action execution failed.", "error_code": "INVALID_PROJECT_SECRET"})
		return
	}

	provider, err := c.providerRepo.GetProviderByName(ctx, providerName)
	if err != nil || provider == nil {
		ctx.JSON(400, gin.H{"error": "Action execution failed.", "error_code": "INVALID_PROVIDER"})
		return
	}

	acc, err := c.connectedAccountRepo.GetByProjectAndProviderAndUserId(ctx, projectID, req.UserID, provider.ID)
	if err != nil || acc == nil {
		ctx.JSON(400, gin.H{"error": "Action execution failed.", "error_code": "NO_CONNECTED_ACCOUNT_FOUND"})
		return
	}

	connector := c.connectorRegistry.GetConnector(providerName)
	if connector == nil {
		ctx.JSON(400, gin.H{"error": "Action execution failed.", "error_code": "CONNECTOR_NOT_REGISTERED"})
		return
	}

	action := connector.GetAction(ctx, actionName)
	if action == nil {
		ctx.JSON(400, gin.H{"error": "Action execution failed.", "error_code": "INVALID_ACTION"})
		return
	}

	userCreds := &common.UserCredentials{
		AccessToken: acc.AccessToken,
		ExpiresAt:   acc.ExpiresAt,
	}

	now := time.Now().UTC()

	// Generating and storing new access token if old is expired.
	if userCreds.ExpiresAt != nil && userCreds.ExpiresAt.UTC().Before(now) {
		providerCredentials, err := c.providerCredentialsRepo.GetByProjectAndProvider(ctx, projectID, provider.ID)
		if err != nil || providerCredentials == nil {
			ctx.JSON(400, gin.H{"error": "Action execution failed.", "error_code": "PROVIDER_CREDENTIALS_NOT_PRESENT"})
			return
		}

		token, err := connector.AuthStrategy().GetNewAccessToken(provider, providerCredentials, acc)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Action execution failed.", "error_code": err.Error()})
			return
		}

		if token == nil {
			ctx.JSON(400, gin.H{"error": "Action execution failed.", "error_code": "UNABLE_TO_REFRESH_TOKEN"})
			return
		}

		userCreds.AccessToken = token.AccessToken

		acc.AccessToken = token.AccessToken

		expireTime := time.Now().UTC().Add(time.Duration(token.ExpiresIn) * time.Second)
		acc.ExpiresAt = &expireTime

		_, err = c.connectedAccountRepo.Update(ctx, acc)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Action execution failed.", "error_code": err.Error()})
			return
		}
	}

	headers := make(map[string]string)
	for k, v := range action.Headers {
		headers[k] = interpolate(v, req.Body, userCreds).(string)
	}

	var bodyBytes []byte
	if req.Body != nil {
		interpolatedBody := interpolate(action.Body, req.Body, userCreds)
		bodyBytes, _ = json.Marshal(interpolatedBody)
	}

	apiCall, err := http.NewRequestWithContext(ctx, action.Method, action.URL, bytes.NewBuffer(bodyBytes))
	if err != nil {
		log.Err(err).Msg("unable to create request for calling action.")
		ctx.JSON(400, gin.H{"error": "Action execution failed.", "error_code": "REQ_CREATION_FAILED"})
		return
	}

	for k, v := range headers {
		apiCall.Header.Set(k, v)
	}

	q := apiCall.URL.Query()
	for k, v := range action.Query {
		q.Set(k, interpolate(v, req.Body, userCreds).(string))
	}
	apiCall.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(apiCall)
	if err != nil {
		log.Err(err).Msg("unable to perform api request for action.")
		ctx.JSON(400, gin.H{"error": "Action execution failed.", "error_code": "REQ_CREATION_FAILED"})
		return
	}

	defer resp.Body.Close()

	var out any
	json.NewDecoder(resp.Body).Decode(&out)

	ctx.JSON(resp.StatusCode, out)
}

func interpolate(value any, input map[string]any, creds *common.UserCredentials) any {
	switch v := value.(type) {
	case string:
		s := v
		s = strings.ReplaceAll(s, "{{access_token}}", creds.AccessToken)

		for k, v := range input {
			s = strings.ReplaceAll(s, "{{"+k+"}}", v.(string))
		}
		return s

	case map[string]any:
		out := make(map[string]any)
		for key, val := range v {
			out[key] = interpolate(val, input, creds)
		}
		return out

	default:
		return value
	}
}
