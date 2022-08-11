// Package middleware defines the middlewares and jwt related operations
package middleware

import (
	"github.com/rs/zerolog"

	"github.com/odetolakehinde/slack-stickers-be/src/pkg/environment"
	"github.com/odetolakehinde/slack-stickers-be/src/pkg/helper"
)

const (
	// RequestBodyInContext context key holder
	RequestBodyInContext = "request_body_in_context"
	// ActorIDInContext context key holder
	ActorIDInContext = "actor_id_in_context"
	// ActorTypeInContext context key holder
	ActorTypeInContext = "actor_type_in_context"
	// UserInContext context key holder
	UserInContext = "user_in_context"
	// packageName name of this package
	packageName = "middleware"
)

type (
	// ActorType is the type of user performing the action
	ActorType string

	// Middleware object
	Middleware struct {
		logger zerolog.Logger
		env    environment.Env
	}
)

// NewMiddleware new instance of our custom ginJwt middleware
func NewMiddleware(z zerolog.Logger, env environment.Env) *Middleware {
	l := z.With().Str(helper.LogStrKeyModule, packageName).Logger()
	return &Middleware{
		logger: l,
		env:    env,
	}
}
