package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/x-sushant-x/connective/internal/core/domain"
	"github.com/x-sushant-x/connective/internal/core/port"
)

type UserHandler struct {
	svc port.UserService
}

func NewUserHandler(svc port.UserService) *UserHandler {
	return &UserHandler{
		svc,
	}
}

// HandleGetUser godoc
// @Security BearerAuth
// @Summary Get logged-in user details.
// @Description Get logged-in user details.
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200 {object} APIResponse{data=domain.User}
// @Router /api/v1/user [get]
func (uh *UserHandler) HandleGetUser(ctx *gin.Context) {
	currentUser, ok := ctx.Get("currentUser")
	if !ok {
		Unauthorized(ctx)
		return
	}

	user := currentUser.(*domain.User)

	if user == nil {
		Unauthorized(ctx)
		return
	}

	Success(ctx, user, "User Details", http.StatusOK)
}
