package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/x-sushant-x/connective/internal/core/port"
	"github.com/x-sushant-x/connective/internal/core/util"
)

type Middleware struct {
	authRepo port.AuthRepo
}

func NewMiddlewareHandler(authRepo port.AuthRepo) *Middleware {
	return &Middleware{
		authRepo,
	}
}

func (m *Middleware) CheckAuth(ctx *gin.Context) {
	// authHeader := ctx.GetHeader("Authorization")
	// if authHeader == "" {
	// 	m.unauthorized(ctx)
	// 	return
	// }

	// parts := strings.SplitN(authHeader, " ", 2)
	// if len(parts) != 2 || parts[0] != "Bearer" {
	// 	m.unauthorized(ctx)
	// 	return
	// }

	// tokenString := parts[1]

	authToken, err := ctx.Cookie("user_token")
	if err != nil || authToken == "" {
		m.unauthorized(ctx)
		return
	}

	token, err := util.ValidateJWTToken(authToken)
	if err != nil || !token.Valid {
		m.unauthorized(ctx)
		return
	}

	claims, ok := util.GetJWTClaims(token)
	if !ok {
		m.unauthorized(ctx)
		return
	}

	exp, ok := claims["exp"].(float64)
	if !ok || float64(time.Now().Unix()) > exp {
		m.unauthorized(ctx)
		return
	}

	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		m.unauthorized(ctx)
		return
	}
	userID := int(userIDFloat)

	user, err := m.authRepo.GetUserById(ctx, userID)
	if err != nil || user == nil {
		m.unauthorized(ctx)
		return
	}

	ctx.Set("currentUser", user)
	ctx.Next()
}

func (m *Middleware) unauthorized(ctx *gin.Context) {
	Unauthorized(ctx)
	ctx.AbortWithStatus(http.StatusUnauthorized)
}
