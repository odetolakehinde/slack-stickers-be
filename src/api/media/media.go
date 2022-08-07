// Package media houses all media related APIs
package media

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	restModel "github.com/odetolakehinde/slack-stickers-be/src/api/model"
	"github.com/odetolakehinde/slack-stickers-be/src/controller"
	"github.com/odetolakehinde/slack-stickers-be/src/pkg/environment"
)

type authHandler struct {
	logger      zerolog.Logger
	controller  controller.Operations
	environment *environment.Env
}

// New creates a new instance of the auth rest handler
func New(r *gin.RouterGroup, l zerolog.Logger, c controller.Operations, env *environment.Env) {
	media := authHandler{
		logger:      l,
		controller:  c,
		environment: env,
	}

	mediaGroup := r.Group("/media")

	// Endpoints exposed under Media API Handler
	mediaGroup.POST("/upload", media.controller.Middleware().AuthMiddleware(), media.upload())
}

// login handles authentication for users
// @Summary This uploads stickers into the database. Keep in mind that it checks to ensure that it dfors
// @Description Takes the user email and password and returns user and token details
// @Tags Auth
// @Accept json
// @Param uploadRequest body uploadRequest true "Upload Request"
// @Success 201 {object} model.GenericResponse{data=uploadRequest}
// @Failure 400,401,502 {object} model.GenericResponse{error=model.GenericResponse}
// @Router /api/v1/auth/login [post]
func (a authHandler) upload() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req uploadRequest

		// run the validation first
		if err := c.ShouldBindJSON(&req); err != nil {
			a.logger.Error().Msgf("%v", err)
			restModel.ErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		err := restModel.ValidateRequest(req)
		if err != nil {
			a.logger.Error().Msgf("%v", err)
			restModel.ErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		restModel.OkResponse(c, http.StatusCreated, "Sticker(s) uploaded successfully", "response")
	}
}
