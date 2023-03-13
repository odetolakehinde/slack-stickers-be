package media

import (
	"context"
	"fmt"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/admin/search"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
	"github.com/rs/zerolog"

	"github.com/odetolakehinde/slack-stickers-be/src/model"
	"github.com/odetolakehinde/slack-stickers-be/src/pkg/environment"
	"github.com/odetolakehinde/slack-stickers-be/src/pkg/helper"
)

const (
	cloudinaryName = "cloudinary"
)

type (
	Cloudinary struct {
		logger zerolog.Logger
		env    *environment.Env
		client *cloudinary.Cloudinary
	}
)

// NewCloudinary initializes a new instance of Cloudinary
func NewCloudinary(l zerolog.Logger, ev *environment.Env) *Cloudinary {
	cld, err := cloudinary.NewFromParams(ev.Get("CLOUDINARY_CLOUD_NAME"), ev.Get("CLOUDINARY_API_KEY"), ev.Get("CLOUDINARY_SECRET_KEY"))
	if err != nil {
		l.Err(err).Msg("unable to connect to cloudinary")
		return nil
	}

	return &Cloudinary{
		logger: l.With().Str(helper.LogStrKeyModule, cloudinaryName).Logger(),
		env:    ev,
		client: cld,
	}
}

// UploadSticker uploads a sticker to the database.
func (c *Cloudinary) UploadSticker(ctx context.Context, name, details string) error {
	c.logger.Info().Msgf("UploadSticker ::: uploading sticker: %s", name)

	_, err := c.client.Upload.Upload(ctx, name, uploader.UploadParams{
		PublicID:                 "",
		PublicIDPrefix:           "",
		PublicIDs:                nil,
		UseFilename:              nil,
		UniqueFilename:           nil,
		DisplayName:              "",
		UseFilenameAsDisplayName: nil,
		FilenameOverride:         "",
		Folder:                   "",
		AssetFolder:              "",
		Overwrite:                nil,
		ResourceType:             "",
		Type:                     "",
		Tags:                     nil,
		Context:                  nil,
		Metadata:                 nil,
		Transformation:           "",
		Format:                   "",
		AllowedFormats:           nil,
		Eager:                    "",
		ResponsiveBreakpoints:    nil,
		Eval:                     "",
		Async:                    nil,
		EagerAsync:               nil,
		Unsigned:                 nil,
		Proxy:                    "",
		Headers:                  "",
		Callback:                 "",
		NotificationURL:          "",
		EagerNotificationURL:     "",
		Faces:                    nil,
		ImageMetadata:            nil,
		Exif:                     nil,
		Colors:                   nil,
		Phash:                    nil,
		FaceCoordinates:          nil,
		CustomCoordinates:        nil,
		Backup:                   nil,
		ReturnDeleteToken:        nil,
		Invalidate:               nil,
		DiscardOriginalFilename:  nil,
		Moderation:               "",
		UploadPreset:             "",
		RawConvert:               "",
		Categorization:           "",
		AutoTagging:              0,
		BackgroundRemoval:        "",
		Detection:                "",
		OCR:                      "",
		QualityAnalysis:          nil,
		AccessibilityAnalysis:    nil,
		CinemagraphAnalysis:      nil,
	})
	if err != nil {
		c.logger.Err(err).Msg("upload failed")
		return err
	}

	return nil
}

// SearchByTag searches the list of stickers by a preferred tag
func (c *Cloudinary) SearchByTag(ctx context.Context, tag string) ([]*model.Sticker, int, error) {
	c.logger.Info().Msgf("SearchByTag ::: searching by tag: %s", tag)

	resp, err := c.client.Admin.Search(ctx, search.Query{
		Expression: fmt.Sprintf("resource_type:image AND tags:%s*", tag),
		WithField:  []string{"tags", "context"},
		SortBy:     []search.SortByField{{"public_id": "desc"}},
		MaxResults: model.MaxResults})

	if err != nil {
		c.logger.Err(err).Msg("search by tag failed")
		return nil, 0, err
	}

	response := make([]*model.Sticker, 0)

	for _, img := range resp.Assets {
		response = append(response, &model.Sticker{
			ID:           uuid.New(), // TODO(Kehinde): This needs attention,
			PublicID:     img.PublicID,
			Name:         img.Filename,
			URL:          img.URL,
			Folder:       img.Folder,
			ResourceType: img.ResourceType,
			Format:       img.Format,
			Status:       img.Status,
			Timestamp:    img.CreatedAt,
			Tags:         img.Tags,
		})
	}

	return response, resp.TotalCount, nil
}
