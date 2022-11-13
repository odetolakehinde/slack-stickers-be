// Package api defines API related operations
package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	// This import is needed for swagger to work
	_ "github.com/odetolakehinde/slack-stickers-be/src/docs"

	"github.com/odetolakehinde/slack-stickers-be/src/api/media"
	slackApi "github.com/odetolakehinde/slack-stickers-be/src/api/slack"
	"github.com/odetolakehinde/slack-stickers-be/src/controller"
	"github.com/odetolakehinde/slack-stickers-be/src/pkg/environment"
	"github.com/odetolakehinde/slack-stickers-be/src/pkg/helper"
)

// Handler object
type Handler struct {
	application controller.Operations
	logger      *zerolog.Logger
	env         *environment.Env
	api         *gin.RouterGroup
}

// New creates a new instance of Handler
func New(z zerolog.Logger, ev *environment.Env, engine *gin.Engine, a controller.Operations) *Handler {
	log := z.With().Str(helper.LogStrPartnerLevel, "handler").Logger()
	//engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	apiGroup := engine.Group("/api")
	return &Handler{
		application: a,
		logger:      &log,
		env:         ev,
		api:         apiGroup,
	}
}

// Build sets up the API handlers
func (h *Handler) Build() {
	v1 := h.api.Group("/v1")
	media.New(v1, *h.logger, h.application, h.env)
	slackApi.New(v1, *h.logger, h.application, h.env)
}
