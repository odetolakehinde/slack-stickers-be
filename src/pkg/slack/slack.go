// Package slack houses all the used slack implementations
package slack

import (
	"log"
	"os"

	"github.com/rs/zerolog"
	"github.com/slack-go/slack"

	"github.com/odetolakehinde/slack-stickers-be/src/pkg/environment"
	"github.com/odetolakehinde/slack-stickers-be/src/pkg/helper"
)

const name = "messaging.slack"

// Provider object
type Provider struct {
	logger zerolog.Logger
	env    *environment.Env
	client *slack.Client
}

// New creates a new instance of the Slack api consumption
func New(z zerolog.Logger, e *environment.Env) *Provider {
	l := z.With().Str(helper.LogStrKeyLevel, name).Logger()
	api := slack.New(
		SLACK_BOT_TOKEN,
		slack.OptionDebug(true),
		slack.OptionAppLevelToken(SLACK_APP_TOKEN),
		slack.OptionLog(log.New(os.Stdout, "api: ", log.Lshortfile|log.LstdFlags)),
	)

	return &Provider{
		logger: l,
		env:    e,
		client: api,
	}
}
