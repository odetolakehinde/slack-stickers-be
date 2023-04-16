// Package controller defines implementation that exposes logics of the app
package controller

import (
	"context"
	"github.com/rs/zerolog"

	"github.com/odetolakehinde/slack-stickers-be/src/media"
	"github.com/odetolakehinde/slack-stickers-be/src/model"
	"github.com/odetolakehinde/slack-stickers-be/src/pkg/environment"
	"github.com/odetolakehinde/slack-stickers-be/src/pkg/helper"
	"github.com/odetolakehinde/slack-stickers-be/src/pkg/middleware"
	"github.com/odetolakehinde/slack-stickers-be/src/slack"
	"github.com/odetolakehinde/slack-stickers-be/src/store"
)

const packageName = "controller"

// Operations enlist all possible operations for this controller across all modules
//go:generate mockgen -source controller.go -destination ./mock/mock_controller.go -package mock Operations
type Operations interface {
	Middleware() *middleware.Middleware

	SendSticker(ctx context.Context, channelID, imageURL string) error
	ShowSearchModal(ctx context.Context, triggerID, channelID string) error
	SearchByTag(ctx context.Context, triggerID, tag, countToReturn, channelID string, externalViewID *string) error
	SaveAuthDetails(ctx context.Context, authDetails model.SlackAuthDetails) error
}

// Controller object to hold necessary reference to other dependencies
type Controller struct {
	logger     zerolog.Logger
	env        *environment.Env
	middleware *middleware.Middleware

	// third party services
	cloudinary   *media.Cloudinary
	slackService slack.Provider
	store        store.Store
}

// New creates a new instance of Controller
func New(z zerolog.Logger, env *environment.Env, m *middleware.Middleware, store store.Store) *Operations {
	l := z.With().Str(helper.LogStrKeyModule, packageName).Logger()

	// init all storage layer under here
	_ = store.Connect()

	// init all third party packages
	cloudinary := media.NewCloudinary(z, env)
	s := slack.New(z, env)

	ctrl := &Controller{
		logger:     l,
		env:        env,
		middleware: m,

		cloudinary:   cloudinary,
		slackService: *s,
		store:        store,
	}

	op := Operations(ctrl)
	return &op
}

// Middleware returns the middleware object exposed by this app
func (c *Controller) Middleware() *middleware.Middleware {
	return c.middleware
}
