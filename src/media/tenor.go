package media

import (
	"context"
	"encoding/json"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog"

	"github.com/odetolakehinde/slack-stickers-be/src/model/env"
	"github.com/odetolakehinde/slack-stickers-be/src/pkg/environment"
	"github.com/odetolakehinde/slack-stickers-be/src/pkg/helper"
)

const (
	tenorName = "tenor"
)

type (
	// Tenor schema
	Tenor struct {
		logger zerolog.Logger
		env    *environment.Env
		client resty.Client
	}
)

// NewTenor initializes a new instance of Cloudinary
func NewTenor(l zerolog.Logger, ev *environment.Env, client resty.Client) *Tenor {
	return &Tenor{
		logger: l.With().Str(helper.LogStrKeyModule, tenorName).Logger(),
		env:    ev,
		client: client,
	}
}

// SearchGifsByQuery searches for GIFs on the Tenor API based on a provided query.
//
// Api Doc: https://developers.google.com/tenor/guides/endpoints
func (t *Tenor) SearchGifsByQuery(ctx context.Context, query string, pos string) (TenorAPIResponse, error) {
	log := t.logger.With().
		Str(helper.LogStrRequestIDLevel, getRequestID(ctx)).
		Str(helper.LogStrKeyMethod, "tenor.SearchTenorGifByQuery").Logger()

	resp, err := t.client.R().SetQueryParams(map[string]string{
		"key":          t.env.Get(env.TenorAPIKey),
		"media_filter": "gif",
		"limit":        "15",
		"random":       "true",
		"q":            query,
		"pos":          pos,
	}).Get("https://tenor.googleapis.com/v2/search")

	if err != nil || resp.StatusCode() != 200 {
		log.Err(err).Msgf("failed to fetch gif from tenor %s", resp.String())
		return TenorAPIResponse{}, err
	}

	var response TenorAPIResponse
	if err = json.Unmarshal(resp.Body(), &response); err != nil {
		log.Err(err).Msg("json.Unmarshal failed")
		return TenorAPIResponse{}, err
	}

	return response, nil
}

func getRequestID(ctx context.Context) string {
	rID := ctx.Value(helper.GinContextKey)
	if rID != nil {
		return rID.(string)
	}

	return helper.ZeroUUID
}
