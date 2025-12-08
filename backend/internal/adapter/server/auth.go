package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/x-sushant-x/connective/internal/core/domain"
	"github.com/x-sushant-x/connective/internal/core/dto"
	"github.com/x-sushant-x/connective/internal/core/port"
)

type AuthHandler struct {
	svc port.AuthService
}

func NewAuthHandler(svc port.AuthService) *AuthHandler {
	return &AuthHandler{
		svc,
	}
}

// HandleCreateUser godoc
// @Summary Create a new user.
// @Description Create a new user.
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param  request body domain.User true "Signup Body"
// @Success 200 {object} APIResponse{data=dto.SignupResponse}
// @Router /api/v1/user/create [post]
func (ah *AuthHandler) HandleCreateUser(ctx *gin.Context) {
	var req domain.User

	if err := ctx.BindJSON(&req); err != nil {
		Error(ctx, err.Error(), 400)
		return
	}

	newUser, err := ah.svc.CreateUser(ctx, &req)
	if err != nil {
		Error(ctx, err.Error(), 400)
		return
	}

	ctx.SetCookie("user_token", newUser.Token, 3600, "/", "localhost", false, true)

	Success(ctx, newUser, "User Created", http.StatusOK)
}

// HandleLoginUser godoc
// @Summary Login User
// @Description Login User and return JWT token.
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param  request body dto.LoginRequest true "Login Body"
// @Success 200 {object} APIResponse{data=dto.LoginResponse}
// @Router /api/v1/user/login [post]
func (ah *AuthHandler) HandleLoginUser(ctx *gin.Context) {
	var req *dto.LoginRequest

	if err := ctx.BindJSON(&req); err != nil {
		Error(ctx, "Bad Request", 400)
		return
	}

	resp, err := ah.svc.LoginUser(ctx, req.Email, req.Password)
	if err != nil {
		Error(ctx, err.Error(), 400)
		return
	}

	ctx.SetCookie("user_token", resp.Token, 36000, "/", "localhost", false, true)

	Success(ctx, resp, "Login Successful", http.StatusOK)
}

// HandleCompleteOnboarding godoc
// @Summary Complete User Onboarding
// @Description Complete User Onboarding
// @Tags Auth
// @Accept  json
// @Produce  json
// @Success 200 {object} APIResponse{data=any}
// @Router /api/v1/user/complete-onboarding [post]
func (ah *AuthHandler) HandleCompleteOnboarding(ctx *gin.Context) {
	currentUser, ok := ctx.Get("currentUser")
	if !ok {
		Unauthorized(ctx)
		return
	}

	user := currentUser.(*domain.User)

	err := ah.svc.MarkOnboardingCompleted(ctx, user.ID)
	if err != nil {
		Error(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	Success(ctx, nil, "Onboarding Completed Successfully", http.StatusOK)
}
