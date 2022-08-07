package media

import (
	"context"

	"github.com/rs/zerolog"

	"github.com/odetolakehinde/slack-stickers-be/src/pkg/environment"
	"github.com/odetolakehinde/slack-stickers-be/src/pkg/helper"
)

const (
	cloudinaryName = "cloudinary"
)

type (
	cloudinaryObj struct {
		logger zerolog.Logger
		env    *environment.Env
	}
)

func newCloudinary(l zerolog.Logger, ev *environment.Env) cloudinaryObj {
	return cloudinaryObj{
		logger: l.With().Str(helper.LogStrKeyModule, cloudinaryName).Logger(),
		env:    ev,
	}
}

func (s cloudinaryObj) SendDynamicEmail(_ context.Context, _ UploadPayload) error {

	return nil
}
