package cron

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/odetolakehinde/slack-stickers-be/src/model"
	"github.com/odetolakehinde/slack-stickers-be/src/model/env"
	"github.com/odetolakehinde/slack-stickers-be/src/pkg/helper"
)

// handleTokenRefresh handles refreshing authed user access tokens
func (c *Cron) handleTokenRefresh() func() {
	return func() {
		ctx := context.WithValue(context.Background(), helper.GinContextKey, generateRequestID())
		log := c.Logger.With().
			Str(helper.LogStrRequestIDLevel, getRequestID(ctx)).
			Str(helper.LogStrKeyMethod, "cron.handleTokenRefresh").
			Logger()

		c.Logger.Info().Msg("[start] running job: handleTokenRefresh")

		keys, err := c.store.ScanKeys(ctx, fmt.Sprintf("%s:*", model.RedisSlackAuthPrefix), 0)
		if err != nil {
			log.Err(err).Msg("Failed to scan Redis for Slack auth keys")
			return
		}

		for _, key := range keys {
			var authDetails model.SlackAuthDetails

			// Retrieve the stored authentication details
			err := c.store.GetJSONValue(ctx, key, &authDetails)
			if err != nil {
				log.Err(err).Str("key", key).Msg("Failed to retrieve Slack auth details")
				continue
			}

			// Refresh the authed user token
			newAccessToken, newRefreshToken, err := c.refreshSlackToken(ctx, authDetails.RefreshToken)
			if err != nil {
				log.Err(err).Str("team_id", authDetails.Team.ID).Msg("Failed to refresh Slack token")
				continue
			}

			// Update the authentication details
			authDetails.AccessToken = newAccessToken
			authDetails.RefreshToken = newRefreshToken

			// Save the updated details back to Redis
			bytes, err := json.Marshal(authDetails)
			if err != nil {
				log.Err(err).Msg("json.Marshal failed")
			}

			err = c.store.SetValue(ctx, key, string(bytes), 0)
			if err != nil {
				log.Err(err).Str("key", key).Msg("Failed to update Slack auth details in Redis")
			}
		}

		c.Logger.Info().Msg("[end] finished job: handleTokenRefresh")
	}
}

func (c *Cron) refreshSlackToken(ctx context.Context, refreshToken string) (newAccessToken string, newRefreshToken string, err error) {
	log := c.Logger.With().
		Str(helper.LogStrRequestIDLevel, getRequestID(ctx)).
		Str(helper.LogStrKeyMethod, "cron.refreshSlackToken").
		Logger()

	resp, err := c.restyClient.R().
		SetFormData(map[string]string{
			"client_id":     c.env.Get(env.SlackClientID),
			"client_secret": c.env.Get(env.SlackClientSecret),
			"grant_type":    "refresh_token",
			"refresh_token": refreshToken,
		}).
		Post("https://slack.com/api/oauth.v2.access")
	if err != nil {
		log.Err(err).Msg("Failed to send request to Slack API for token refresh")
		return "", "", err
	}
	if resp.StatusCode() != 200 {
		log.Error().Int("status_code", resp.StatusCode()).Str("response", resp.String()).Msg("Slack API returned an unexpected status code")
		return "", "", fmt.Errorf("unexpected status code from Slack API: %d", resp.StatusCode())
	}

	var response SlackTokenRefreshResponse
	if err = json.Unmarshal(resp.Body(), &response); err != nil {
		log.Err(err).Msg("Failed to unmarshal Slack API response")
		return "", "", err
	}

	if !response.Ok {
		log.Error().Str("error", response.Error).Msg("Slack API returned an error")
		return "", "", fmt.Errorf("slack API error: %s", response.Error)
	}

	return response.AccessToken, response.RefreshToken, nil
}
