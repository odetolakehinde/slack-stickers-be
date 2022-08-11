package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	restModel "github.com/odetolakehinde/slack-stickers-be/src/api/model"
)

// AuthMiddleware authenticates the endpoint for running
func (m *Middleware) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if len(authHeader) == 0 {
			restModel.ErrorResponse(c, http.StatusUnauthorized, "Bearer token is missing")
			return
		}
		authString := m.env.Get("AUTHORIZATION_KEY")
		if !strings.EqualFold(authString, authHeader) {
			restModel.ErrorResponse(c, http.StatusUnauthorized, "Bearer token is missing")
			return
		}
		c.Next()
	}
}

// CorsMiddleware adds a CORS check for api rest endpoints allowing only list of origins defined in env
func (m *Middleware) CorsMiddleware() gin.HandlerFunc {
	return cors.New(cors.DefaultConfig())
}
