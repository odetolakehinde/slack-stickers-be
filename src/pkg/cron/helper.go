package cron

import (
	"context"

	"github.com/google/uuid"
	"github.com/odetolakehinde/slack-stickers-be/src/pkg/helper"
)

// SlackTokenRefreshResponse represents the tokens and their status when refreshing Slack authentication tokens
type SlackTokenRefreshResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Ok           bool   `json:"ok"`
	Error        string `json:"error,omitempty"`
}

func getRequestID(ctx context.Context) string {
	rID := ctx.Value(helper.GinContextKey)
	if rID != nil {
		return rID.(string)
	}

	return helper.ZeroUUID
}

func generateRequestID() string {
	return uuid.New().String()
}
