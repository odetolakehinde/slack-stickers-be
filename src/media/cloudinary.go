package media

import (
	"context"
	"fmt"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/admin/search"
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

	//_, err := c.client.Upload.Upload(ctx, name, uploader.UploadParams{})
	//if err != nil {
	//	c.logger.Err(err).Msg("upload failed")
	//	return err
	//}

	return nil
}

// SearchByTag searches the list of stickers by a preferred tag
func (c *Cloudinary) SearchByTag(ctx context.Context, tag string) ([]*model.Sticker, error) {
	c.logger.Info().Msgf("SearchByTag ::: searching by tag: %s", tag)

	resp, err := c.client.Admin.Search(ctx, search.Query{
		Expression: fmt.Sprintf("resource_type:image AND tags:%s*", tag),
		WithField:  []string{"tags", "context"},
		SortBy:     []search.SortByField{{"public_id": "desc"}},
		MaxResults: 10})

	if err != nil {
		c.logger.Err(err).Msg("upload failed")
	}

	fmt.Println(resp.Response)

	return nil, nil
}
