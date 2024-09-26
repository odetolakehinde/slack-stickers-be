// Package controller defines implementation that exposes logics of the app
package controller

import (
	"context"
	"encoding/json"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

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
//
//go:generate mockgen -source controller.go -destination ./mock/mock_controller.go -package mock Operations
type Operations interface {
	Middleware() *middleware.Middleware

	SendSticker(ctx context.Context, channelID, imageURL, teamID, threadTs string) error
	ShowSearchModal(ctx context.Context, triggerID, channelID, teamID string) error
	SearchByTag(ctx context.Context, triggerID, tag, countToReturn, channelID, teamID string, externalViewID *string) error
	SaveAuthDetails(ctx context.Context, authDetails model.SlackAuthDetails) error
}

// Controller object to hold necessary reference to other dependencies
type Controller struct {
	logger     zerolog.Logger
	env        *environment.Env
	middleware *middleware.Middleware

	// third party services
	cloudinary *media.Cloudinary
	store      store.Store
}

// New creates a new instance of Controller
func New(z zerolog.Logger, env *environment.Env, m *middleware.Middleware, store store.Store) *Operations {
	l := z.With().Str(helper.LogStrKeyModule, packageName).Logger()

	// init all storage layer under here
	_ = store.Connect()

	// init all third party packages
	cloudinary := media.NewCloudinary(z, env)

	ctrl := &Controller{
		logger:     l,
		env:        env,
		middleware: m,

		cloudinary: cloudinary,
		store:      store,
	}

	op := Operations(ctrl)
	return &op
}

// Middleware returns the middleware object exposed by this app
func (c *Controller) Middleware() *middleware.Middleware {
	return c.middleware
}

func (c *Controller) getSlackService(ctx context.Context, teamID string) slack.Provider {
	c.logger.Info().Str("team_id", teamID).Msg("about to get the slack service now...")

	var (
		keyValue string
	)

	// get the token
	err := c.store.GetValue(context.Background(), teamID, &keyValue)
	if err != nil {
		log.Err(err).Msgf("redis.GetValue[%s] failed", teamID)
		//return err
	}

	var authDetails model.SlackAuthDetails
	err = json.Unmarshal([]byte(keyValue), &authDetails)
	if err != nil {
		c.logger.Err(err).Msg("json.Unmarshal failed")
		//return err
	}

	s := slack.New(c.logger, c.env, authDetails.AccessToken)
	return *s
}
