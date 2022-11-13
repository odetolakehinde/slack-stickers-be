// Package media houses all media related APIs
package media

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/slack-go/slack"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/schema"
	"github.com/rs/zerolog"

	restModel "github.com/odetolakehinde/slack-stickers-be/src/api/model"
	"github.com/odetolakehinde/slack-stickers-be/src/controller"
	"github.com/odetolakehinde/slack-stickers-be/src/pkg/environment"
)

type slackHandler struct {
	logger      zerolog.Logger
	controller  controller.Operations
	environment *environment.Env
}

// New creates a new instance of the auth rest handler
func New(r *gin.RouterGroup, l zerolog.Logger, c controller.Operations, env *environment.Env) {
	slack := slackHandler{
		logger:      l,
		controller:  c,
		environment: env,
	}

	slackGroup := r.Group("/slack") // ,slack.controller.Middleware().AuthMiddleware(),

	// Endpoints exposed under Media API Handler
	slackGroup.POST("/send-message", slack.sendMessage())
	slackGroup.POST("/interactivity", slack.interactivityUsed())
	slackGroup.POST("/slash-command", slack.slashCommandUsed())
}

// sendMessage handles authentication for users
// @Summary This uploads stickers into the database. Keep in mind that it checks to ensure that it dfors
// @Description Takes the user email and password and returns user and token details
// @Tags Auth
// @Accept json
// @Param uploadRequest body uploadRequest true "Upload Request"
// @Success 201 {object} model.GenericResponse{data=uploadRequest}
// @Failure 400,401,502 {object} model.GenericResponse{error=model.GenericResponse}
// @Router /api/v1/slack/send-message [post]
func (s slackHandler) sendMessage() gin.HandlerFunc {
	return func(c *gin.Context) {
		//var req uploadRequest
		//
		//// run the validation first
		//if err := c.ShouldBindJSON(&req); err != nil {
		//	s.logger.Error().Msgf("%v", err)
		//	restModel.ErrorResponse(c, http.StatusBadRequest, err.Error())
		//	return
		//}
		//
		//err := restModel.ValidateRequest(req)
		//if err != nil {
		//	s.logger.Error().Msgf("%v", err)
		//	restModel.ErrorResponse(c, http.StatusBadRequest, err.Error())
		//	return
		//}

		err := s.controller.SendSticker(context.Background(), "", "")
		if err != nil {
			s.logger.Error().Msgf("%v", err)
			restModel.ErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		restModel.OkResponse(c, http.StatusOK, "Message sent successfully", "response")
	}
}

func (s slackHandler) interactivityUsed() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			s.logger.Err(err).Msgf("error parsing response :: %v", err.Error())
		}

		parsedBody, err := url.ParseQuery(string(requestBody))
		if err != nil {
			restModel.ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		var i slack.InteractionCallback
		err = json.Unmarshal([]byte(parsedBody["payload"][0]), &i)
		if err != nil {
			restModel.ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		// Note there might be a better way to get this info, but I figured this structure out from looking at the json response
		tag := i.View.State.Values["Tag"]["tag"].Value
		fmt.Println(tag)

		err = s.controller.SearchByTag(context.Background(), tag)
		restModel.OkResponse(c, http.StatusOK, "Shortcut initiated", "response")
	}
}

func (s slackHandler) slashCommandUsed() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req restModel.ShortcutPayload

		requestBody, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			s.logger.Err(err).Msgf("error parsing response :: %v", err.Error())
		}

		parsedBody, err := url.ParseQuery(string(requestBody))
		if err != nil {
			restModel.ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		var decoder = schema.NewDecoder()
		err = decoder.Decode(&req, parsedBody)
		if err != nil {
			// Handle error;
			s.logger.Err(err).Msg("e don happen")
		}

		err = s.controller.ShowSearchModal(context.Background(), req.ChannelID, req.TriggerID)
		if err != nil {
			s.logger.Error().Msgf("%v", err)
			restModel.ErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		restModel.OkResponse(c, http.StatusOK, "Slash command initiated", "response")
	}
}
