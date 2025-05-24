// Package controller defines implementation that exposes logics of the app
package controller

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
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
//
//go:generate mockgen -source controller.go -destination ./mock/mock_controller.go -package mock Operations
type Operations interface {
	Middleware() *middleware.Middleware

	// TODO: would prob delete this
	SendMessage(ctx context.Context, channelID, imageURL, teamID string) error

	ShowSearchModal(ctx context.Context, triggerID, channelID, teamID string) error
	ShowSearchResultModal(ctx context.Context, triggerID, channelID, teamID string, sticker model.StickerBlockMetadata, externalViewID *string) error

	GetStickerSearchResult(ctx context.Context, teamID, userID, channelID, tag string, threadTS, mentionTS *string) error
	CancelSticker(ctx context.Context, teamID, channelID, responseURL string) error
	SendSticker(ctx context.Context, teamID, userID, channelID, responseURL string, sticker model.StickerBlockMetadata) error
	ShuffleSticker(ctx context.Context, teamID, userID, channelID, responseURL string, sticker model.StickerBlockMetadata) error

	SaveAuthDetails(ctx context.Context, authDetails model.SlackAuthDetails) error
	RemoveAuthDetails(ctx context.Context, teamID string) error
}

// Controller object to hold necessary reference to other dependencies
type Controller struct {
	logger     zerolog.Logger
	env        *environment.Env
	middleware *middleware.Middleware

	// third party services
	cloudinary *media.Cloudinary
	tenor      *media.Tenor
	store      store.Store
}

// New creates a new instance of Controller
func New(z zerolog.Logger, env *environment.Env, m *middleware.Middleware, store store.Store) *Operations {
	l := z.With().Str(helper.LogStrKeyModule, packageName).Logger()

	// init all storage layer under here
	_ = store.Connect()

	restyClient := resty.New()

	// init all third party packages
	cloudinary := media.NewCloudinary(z, env)
	tenor := media.NewTenor(z, env, *restyClient)

	ctrl := &Controller{
		logger:     l,
		env:        env,
		middleware: m,

		cloudinary: cloudinary,
		tenor:      tenor,

		store: store,
	}

	op := Operations(ctrl)
	return &op
}

// Middleware returns the middleware object exposed by this app
func (c *Controller) Middleware() *middleware.Middleware {
	return c.middleware
}

func (c *Controller) getSlackService(ctx context.Context, teamID string) (slack.Provider, error) {
	log := c.logger.With().Str(helper.LogStrKeyMethod, "getSlackService").Logger()
	log.Info().Str("team_id", teamID).Msg("about to get the slack service now...")

	var keyValue string

	// get the token
	key := fmt.Sprintf("%s:%s", model.RedisSlackAuthPrefix, teamID)
	err := c.store.GetValue(ctx, key, &keyValue)
	if err != nil {
		log.Err(err).Msgf("redis.GetValue[%s] failed", teamID)
		return slack.Provider{}, err
	}

	var authDetails model.SlackAuthDetails
	err = json.Unmarshal([]byte(keyValue), &authDetails)
	if err != nil {
		log.Err(err).Msg("json.Unmarshal failed")
		return slack.Provider{}, err
	}

	accessToken := authDetails.AccessToken

	s := slack.New(c.logger, c.env, accessToken)
	return *s, nil
}
