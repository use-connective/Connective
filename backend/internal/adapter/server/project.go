package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/x-sushant-x/connective/internal/core/domain"
	"github.com/x-sushant-x/connective/internal/core/dto"
	"github.com/x-sushant-x/connective/internal/core/port"
)

type ProjectHandler struct {
	svc port.ProjectService
}

func NewProjectHandler(svc port.ProjectService) *ProjectHandler {
	return &ProjectHandler{
		svc,
	}
}

// HandleCreateProject godoc
// @Security BearerAuth
// @Summary Create a new project.
// @Description Create a new project.
// @Tags Project
// @Accept  json
// @Produce  json
// @Param  request body dto.CreateProjectRequest true "Create Project Body"
// @Success 200 {object} APIResponse{data=domain.Project}
// @Router /api/v1/project/create [post]
func (ph *ProjectHandler) HandleCreateProject(ctx *gin.Context) {
	var req *dto.CreateProjectRequest

	if err := ctx.BindJSON(&req); err != nil {
		BadRequest(ctx, "Bad Request")
		return
	}

	currentUser, ok := ctx.Get("currentUser")
	if !ok {
		Unauthorized(ctx)
		return
	}

	user := currentUser.(*domain.User)

	newProject, err := ph.svc.CreateProject(ctx, user, req)
	if err != nil {
		BadRequest(ctx, err.Error())
		return
	}

	Success(ctx, newProject, "Project Created", http.StatusOK)
}

// HandleGetAllProjects godoc
// @Summary Get all projects of a user.
// @Description Get all projects of a user.
// @Tags Project
// @Accept  json
// @Produce  json
// @Success 200 {object} APIResponse{data=[]domain.Project}
// @Router /api/v1/project/get-all [get]
func (ph *ProjectHandler) HandleGetAllProjects(ctx *gin.Context) {
	currentUser, ok := ctx.Get("currentUser")
	if !ok {
		Unauthorized(ctx)
		return
	}

	user := currentUser.(*domain.User)

	resp, err := ph.svc.GetByOwner(ctx, user.ID)
	if err != nil {
		BadRequest(ctx, err.Error())
		return
	}

	Success(ctx, resp, "Projects List Fetched", http.StatusOK)
}
