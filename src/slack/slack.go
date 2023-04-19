// Package slack houses all the used slack implementations
package slack

import (
	"log"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog"
	"github.com/slack-go/slack"

	"github.com/odetolakehinde/slack-stickers-be/src/pkg/environment"
	"github.com/odetolakehinde/slack-stickers-be/src/pkg/helper"
)

const name = "messaging.slack"

// Provider object
type Provider struct {
	logger     zerolog.Logger
	env        *environment.Env
	client     *slack.Client
	httpClient *resty.Client
}

// New creates a new instance of the Slack api consumption
func New(z zerolog.Logger, e *environment.Env, token string) *Provider {
	l := z.With().Str(helper.LogStrKeyLevel, name).Logger()
	api := slack.New(
		token,
		slack.OptionAppLevelToken(e.Get("SLACK_APP_TOKEN")),
		slack.OptionLog(log.New(os.Stdout, "api: ", log.Lshortfile|log.LstdFlags)),
	)

	httpClient := resty.New()
	httpClient.SetTimeout(15 * time.Second)

	if e.IsSandbox() {
		httpClient.SetDebug(true)
		httpClient.EnableTrace()
		api.Debug()
	}

	return &Provider{
		logger:     l,
		env:        e,
		client:     api,
		httpClient: httpClient,
	}
}
