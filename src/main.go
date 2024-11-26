// Package main
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	ginzerolog "github.com/dn365/gin-zerolog"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	// This import is needed for swagger to work
	"github.com/odetolakehinde/slack-stickers-be/src/api/model"
	_ "github.com/odetolakehinde/slack-stickers-be/src/docs"
	e "github.com/odetolakehinde/slack-stickers-be/src/model/env"

	"github.com/odetolakehinde/slack-stickers-be/src/api"
	"github.com/odetolakehinde/slack-stickers-be/src/controller"
	"github.com/odetolakehinde/slack-stickers-be/src/pkg/environment"
	"github.com/odetolakehinde/slack-stickers-be/src/pkg/helper"
	"github.com/odetolakehinde/slack-stickers-be/src/pkg/middleware"
	"github.com/odetolakehinde/slack-stickers-be/src/store"
)

// @title Slack Stickers API documentation
// @version 1.0.0
// @description This documents all rest endpoints exposed by this application. Please support us @ buymeacoffee.com/slackstickers

// @contact.name Slack Stickers
// @contact.email slackstickers@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:6001
// @BasePath /api/v1/
func main() {
	// set global application timezone
	_ = os.Setenv("TZ", "Africa/Lagos")
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	applicationLogger := logger.With().Str(helper.LogStrKeyModule, "app").Logger()

	env, err := environment.New()
	if err != nil {
		applicationLogger.Fatal().Err(err)
		panic(err) // panic - this service should not start up
	}

	r := gin.New()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowHeaders = []string{"*"}
	corsConfig.AllowAllOrigins = true
	r.Use(cors.New(corsConfig), gin.Recovery())
	r.Use(ginzerolog.Logger("rest"))
	r.Use(GinContextToContextMiddleware())
	r.Use(requestid.New())
	r.NoRoute(func(c *gin.Context) {
		model.ErrorResponse(c, http.StatusNotFound, "404 page not found")
	})

	// init our custom middleware
	newMiddleware := middleware.NewMiddleware(logger, *env)

	// init the storage
	redisConn := store.ConnectionInfo{
		Address:  env.Get(e.RedisServerAddress),
		Password: env.Get(e.RedisServerPassword),
		Username: env.Get(e.RedisServerUsername),
	}
	db := store.NewRedis(env, applicationLogger, redisConn)
	application := controller.New(logger, env, newMiddleware, db)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Slack is about to get lit with stickers!",
			"rest":    true,
			"graphql": false,
		})
	})

	h := api.New(logger, env, r, *application)
	h.Build()

	port := env.Get(e.ServerPort)
	if strings.EqualFold(port, "") {
		port = "6001"
	}
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: r,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			applicationLogger.Fatal().Msgf("listen: %s", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	applicationLogger.Info().Msgf("Shutdown Server ... %v", <-quit)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		applicationLogger.Fatal().Msgf("Server Shutdown: %v", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		applicationLogger.Info().Msgf("timeout of 5 seconds.")
	default:
	}

	applicationLogger.Info().Msgf("Server exiting")
}

// GinContextToContextMiddleware middleware for gin context
func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), helper.GinContextKey, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
