// Package controller defines implementation that exposes logics of the app
package controller

import (
	"github.com/rs/zerolog"

	"github.com/odetolakehinde/slack-stickers-be/src/media"
	"github.com/odetolakehinde/slack-stickers-be/src/pkg/environment"
	"github.com/odetolakehinde/slack-stickers-be/src/pkg/helper"
	"github.com/odetolakehinde/slack-stickers-be/src/pkg/middleware"
)

const packageName = "controller"

// Operations enlist all possible operations for this controller across all modules
//go:generate mockgen -source controller.go -destination ./mock/mock_controller.go -package mock Operations
type Operations interface {
	Middleware() *middleware.Middleware
}

// Controller object to hold necessary reference to other dependencies
type Controller struct {
	logger     zerolog.Logger
	env        *environment.Env
	middleware *middleware.Middleware

	// third party services
	mediaService media.Service
}

// New creates a new instance of Controller
func New(z zerolog.Logger, env *environment.Env, m *middleware.Middleware) *Operations {
	l := z.With().Str(helper.LogStrKeyModule, packageName).Logger()

	// init all storage layer under here

	media := media.New(z, env)

	ctrl := &Controller{
		logger:     l,
		env:        env,
		middleware: m,

		mediaService: *media,
	}

	op := Operations(ctrl)
	return &op
}

// Middleware returns the middleware object exposed by this app
func (c *Controller) Middleware() *middleware.Middleware {
	return c.middleware
}
