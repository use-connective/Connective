package connectors

import (
	"github.com/gin-gonic/gin"
	"github.com/x-sushant-x/connective/internal/core/port"
)

type ConnectorHandler struct {
	providerRepo            port.ProviderRepo
	connectedAccountRepo    port.ConnectedAccountRepo
	projectRepo             port.ProjectRepo
	providerCredentialsRepo port.ProviderCredentialsRepo
}

func NewConnectorHandler(providerRepo port.ProviderRepo, connectedAccountRepo port.ConnectedAccountRepo, projectRepo port.ProjectRepo, providerCredentialsRepo port.ProviderCredentialsRepo) *ConnectorHandler {
	return &ConnectorHandler{
		providerRepo:            providerRepo,
		connectedAccountRepo:    connectedAccountRepo,
		projectRepo:             projectRepo,
		providerCredentialsRepo: providerCredentialsRepo,
	}
}

func (c *ConnectorHandler) ExecuteAction(ctx *gin.Context) {
	// ToDo - This needs to be rewritten according to new cache architecture.

	//var req common.ActionExecuteRequest
	//
	//if err := ctx.ShouldBindJSON(&req); err != nil {
	//	ctx.JSON(400, gin.H{"error": "Action execution failed.", "error_code": "INVALID_REQUEST"})
	//	return
	//}
	//
	//projectID := req.ProjectID
	//providerName := ctx.Param("provider")
	//actionName := ctx.Param("action")
	//
	//project, err := c.projectRepo.GetByID(ctx, projectID)
	//if err != nil || project == nil {
	//	ctx.JSON(400, gin.H{"error": "Action execution failed.", "error_code": "INVALID_PROJECT"})
	//	return
	//}
	//
	//if req.ProjectSecret != project.SDKAuthSecret {
	//	ctx.JSON(400, gin.H{"error": "Action execution failed.", "error_code": "INVALID_PROJECT_SECRET"})
	//	return
	//}
	//
	//provider, err := c.providerRepo.GetProviderByName(ctx, providerName)
	//if err != nil || provider == nil {
	//	ctx.JSON(400, gin.H{"error": "Action execution failed.", "error_code": "INVALID_PROVIDER"})
	//	return
	//}
	//
	//acc, err := c.connectedAccountRepo.GetByProjectAndProviderAndUserId(ctx, projectID, req.UserID, provider.ID)
	//if err != nil || acc == nil {
	//	ctx.JSON(400, gin.H{"error": "Action execution failed.", "error_code": "NO_CONNECTED_ACCOUNT_FOUND"})
	//	return
	//}
	//
	//connector, ok := Registry[providerName]
	//if !ok {
	//	ctx.JSON(400, gin.H{"error": "Action execution failed.", "error_code": "CONNECTOR_NOT_REGISTERED"})
	//	return
	//}
	//
	//action, ok := connector.Actions()[actionName]
	//if !ok {
	//	ctx.JSON(400, gin.H{"error": "Action execution failed.", "error_code": "INVALID_ACTION"})
	//	return
	//}
	//
	//userCreds := &common.UserCredentials{
	//	AccessToken:  acc.AccessToken,
	//	RefreshToken: acc.RefreshToken,
	//	ExpiresAt:    acc.ExpiresAt,
	//}
	//
	//now := time.Now().UTC()
	//
	//// Generating and storing new access token if old is expired.
	//if userCreds.ExpiresAt.UTC().Before(now) {
	//	providerCredentials, err := c.providerCredentialsRepo.GetByProjectAndProvider(ctx, projectID, provider.ID)
	//	if err != nil || providerCredentials == nil {
	//		ctx.JSON(400, gin.H{"error": "Action execution failed.", "error_code": "PROVIDER_CREDENTIALS_NOT_PRESENT"})
	//		return
	//	}
	//
	//	token, err := connector.AuthStrategy().GetNewAccessToken(provider, providerCredentials, acc)
	//	if err != nil {
	//		ctx.JSON(400, gin.H{"error": "Action execution failed.", "error_code": err.Error()})
	//		return
	//	}
	//
	//	if token == nil {
	//		ctx.JSON(400, gin.H{"error": "Action execution failed.", "error_code": "UNABLE_TO_REFRESH_TOKEN"})
	//		return
	//	}
	//
	//	userCreds.AccessToken = token.AccessToken
	//
	//	acc.AccessToken = token.AccessToken
	//
	//	expireTime := time.Now().UTC().Add(time.Duration(token.ExpiresIn) * time.Second)
	//	acc.ExpiresAt = &expireTime
	//
	//	_, err = c.connectedAccountRepo.Update(ctx, acc)
	//	if err != nil {
	//		ctx.JSON(400, gin.H{"error": "Action execution failed.", "error_code": err.Error()})
	//		return
	//	}
	//}
	//
	//resp, err := action(ctx, userCreds, req.RequestBody)
	//if err != nil {
	//	ctx.JSON(400, gin.H{"error": "Action execution failed.", "error_code": "ACTION_FAILED"})
	//	return
	//}
	//
	//ctx.JSON(200, gin.H{
	//	"success": true,
	//	"data":    resp,
	//})
}
