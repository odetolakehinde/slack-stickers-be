// Package media defines the media related operations and activities
package media

import (
	"context"

	"github.com/rs/zerolog"

	"github.com/odetolakehinde/slack-stickers-be/src/pkg/environment"
	"github.com/odetolakehinde/slack-stickers-be/src/pkg/helper"
)

const name = "media"

// Service enlist all possible operations for media in the platform
//go:generate mockgen -source media.go -destination ./mock_media.go -package media MediaService
type Service interface {
	UploadSticker(ctx context.Context, name string) error
}

type (
	mediaProvider interface {
		SendDynamicEmail(ctx context.Context, payload UploadPayload) error
	}

	// Handler object for email
	Handler struct {
		media  mediaProvider
		env    *environment.Env
		logger zerolog.Logger
	}
)

// UploadSticker uploads a sticker to the data base.
func (h *Handler) UploadSticker(ctx context.Context, name string) error {
	h.logger.Info().Msgf("UploadSticker ::: uploading sticker: %s", name)

	return nil
}

// New creates a new instance of email Handler
func New(z zerolog.Logger, ev *environment.Env) *Service {
	l := z.With().Str(helper.LogStrKeyLevel, name).Logger()

	// init talking to partner endpoint
	registeredGateways := initialize(l, ev)
	defaultMediaService := registeredGateways[defaultEmailGateway()]

	h := &Handler{
		media:  defaultMediaService,
		env:    ev,
		logger: l,
	}
	media := Service(h)
	return &media
}

// initializes sets up new providers accordingly.
// When new providers is added, this snippet needs to be updated
func initialize(l zerolog.Logger, env *environment.Env) map[string]mediaProvider {
	return map[string]mediaProvider{
		cloudinaryName: newCloudinary(l, env),
	}
}

func defaultEmailGateway() string {
	return cloudinaryName // env.Get("DEFAULT_MEDIA_GATEWAY")
}
