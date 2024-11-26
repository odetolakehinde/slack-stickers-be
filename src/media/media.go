// Package media defines the media related operations and activities
package media

import (
	"context"

	"github.com/odetolakehinde/slack-stickers-be/src/model"
)

const name = "media" //nolint:unused

// Service enlist all possible operations for media in the platform
//
//go:generate mockgen -source media.go -destination ./mock_media.go -package media Service
type Service interface {
	UploadSticker(ctx context.Context, name, details string) error
	SearchByTag(ctx context.Context, tag string) ([]*model.Sticker, error)
}
