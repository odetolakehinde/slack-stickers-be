// Package cron defines all cron jobs
package cron

import (
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/go-resty/resty/v2"
	"github.com/odetolakehinde/slack-stickers-be/src/controller"
	"github.com/odetolakehinde/slack-stickers-be/src/pkg/environment"
	"github.com/odetolakehinde/slack-stickers-be/src/pkg/helper"
	"github.com/odetolakehinde/slack-stickers-be/src/store"
	"github.com/rs/zerolog"
)

const (
	packageName                   = "cron"
	sixHourDuration time.Duration = 6 * time.Hour
)

// Cron object
type Cron struct {
	Logger     zerolog.Logger
	controller controller.Operations
	env        *environment.Env

	store       store.Store
	restyClient *resty.Client
}

// New creates new instance of cron
func New(z zerolog.Logger, e *environment.Env, s store.Store, c controller.Operations) gocron.Scheduler {
	log := z.With().Str(helper.LogStrPackageLevel, packageName).Logger()

	restyClient := resty.New()
	cron := Cron{
		Logger:      log,
		env:         e,
		controller:  c,
		store:       s,
		restyClient: restyClient,
	}

	// initialize cron jobs
	return cron.cronJobs()
}

func (c *Cron) cronJobs() gocron.Scheduler {
	cronScheduler, err := gocron.NewScheduler()
	if err != nil {
		c.Logger.Err(err).Msg("gocron.NewScheduler failed")
	}
	//
	if _, err := cronScheduler.NewJob(
		gocron.DurationJob(sixHourDuration), // runs every 6 hours
		gocron.NewTask(c.handleTokenRefresh())); err != nil {
		c.Logger.Err(err).Msg("c.handleRefreshingAccessTokens job failed")
	}

	return cronScheduler
}
