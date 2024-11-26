// Package media houses all media related APIs
package media

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/schema"
	"github.com/mitchellh/mapstructure"
	"github.com/rs/zerolog"
	"github.com/slack-go/slack"

	restModel "github.com/odetolakehinde/slack-stickers-be/src/api/model"
	"github.com/odetolakehinde/slack-stickers-be/src/controller"
	"github.com/odetolakehinde/slack-stickers-be/src/model"
	"github.com/odetolakehinde/slack-stickers-be/src/pkg/environment"
	"github.com/odetolakehinde/slack-stickers-be/src/pkg/helper"
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
	slackGroup.POST("/auth", slack.saveAuthDetails())
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
func (s *slackHandler) sendMessage() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := s.controller.SendSticker(context.Background(), "", "", "")
		if err != nil {
			s.logger.Error().Msgf("%v", err)
			restModel.ErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		restModel.OkResponse(c, http.StatusOK, "Message sent successfully", "response")
	}
}

func (s *slackHandler) interactivityUsed() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody, err := io.ReadAll(c.Request.Body)
		if err != nil {
			s.logger.Err(err).Msgf("error parsing response :: %v", err.Error())
		}

		parsedBody, err := url.ParseQuery(string(requestBody))
		if err != nil {
			restModel.ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		var (
			i             slack.InteractionCallback
			tag           string
			indexToReturn = "0"
		)
		err = json.Unmarshal([]byte(parsedBody["payload"][0]), &i)
		if err != nil {
			restModel.ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		teamID := i.Team.ID
		channelID := i.Channel.ID
		responseURL := i.ResponseURL
		userID := i.User.ID

		switch i.Type {
		case model.ShortcutType:
			err = s.controller.ShowSearchModal(context.Background(), i.TriggerID, i.CallbackID, i.Team.ID)
			if err != nil {
				s.logger.Error().Msgf("%v", err)
				restModel.ErrorResponse(c, http.StatusBadRequest, err.Error())
				return
			}
		case model.SubmissionViewType:
			if len(i.View.Blocks.BlockSet) > 1 && i.View.Blocks.BlockSet[1].BlockType() == model.BlockTypeImage {
				// they actually wanna send the message. Let us proceed
				var details Block

				err = mapstructure.Decode(i.View.Blocks.BlockSet[1], &details)
				if err != nil {
					s.logger.Error().Msgf("%v", err)
					restModel.ErrorResponse(c, http.StatusBadRequest, err.Error())
					return
				}

				channelToSendSticker := i.View.CallbackID
				err = s.controller.SendSticker(context.Background(), channelToSendSticker, details.ImageURL, i.View.TeamID)
				if err != nil {
					s.logger.Error().Msgf("%v", err)
					restModel.ErrorResponse(c, http.StatusBadRequest, err.Error())
					return
				}

				c.JSON(http.StatusOK, struct {
					ResponseAction string `json:"response_action"`
				}{
					ResponseAction: "clear",
				})

				// c.String(http.StatusOK, "Hurray! You've sent your sticker")
				return
			}

			// this is the initial search
			if i.View.CallbackID == model.InitialDataSearchID {
				// Note there might be a better way to get this info, but I figured this structure out from looking at the json response
				tag = i.View.State.Values["Tag"]["tag"].Value
			}
		case model.BlockActionsViewType:
			blockActions := i.ActionCallback.BlockActions
			if len(blockActions) < 1 {
				restModel.ErrorResponse(c, http.StatusBadRequest, "invalid block action")
				return
			}

			if blockActions[0].BlockID == model.StickerActionBlockID {
				blockValue := blockActions[0].Value

				for _, v := range blockActions {
					switch v.ActionID {
					case model.ActionIDSendSticker:
						var sticker model.StickerBlockActionValue
						if err := json.Unmarshal([]byte(blockValue), &sticker); err != nil {
							s.logger.Error().Msgf("%v", err)
							restModel.ErrorResponse(c, http.StatusBadRequest, err.Error())
							return
						}

						err = s.controller.SendSticker(context.Background(), teamID, userID, channelID, responseURL, sticker)
						if err != nil {
							s.logger.Error().Msgf("%v", err)
							restModel.ErrorResponse(c, http.StatusBadRequest, err.Error())
							return
						}

						c.String(http.StatusOK, "action is send")
						return

					case model.ActionIDShuffleSticker:
						var sticker model.StickerBlockActionValue
						if err := json.Unmarshal([]byte(blockValue), &sticker); err != nil {
							s.logger.Error().Msgf("%v", err)
							restModel.ErrorResponse(c, http.StatusBadRequest, err.Error())
							return
						}
						err = s.controller.ShuffleSticker(context.Background(), teamID, userID, channelID, responseURL, sticker)
						if err != nil {
							s.logger.Error().Msgf("%v", err)
							restModel.ErrorResponse(c, http.StatusBadRequest, err.Error())
							return
						}
						c.String(http.StatusOK, "action is shuffle")
						return

					case model.ActionIDCancelSticker:
						if err := s.controller.CancelSticker(context.Background(), teamID, channelID, responseURL); err != nil {
							restModel.ErrorResponse(c, http.StatusBadRequest, err.Error())
							return
						}
						c.String(http.StatusOK, "action is cancel")
						return
					}
				}
			} else {
				indexToReturn = i.ActionCallback.BlockActions[0].Value
				tag = i.View.PrivateMetadata
			}

		default:
			// Note there might be a better way to get this info, but I figured this structure out from looking at the json response
			tag = i.View.State.Values["Tag"]["tag"].Value
			// TODO: this fails when we try to run it in threads.
		}

		externalViewID := i.View.ExternalID
		teamID := i.View.TeamID

		err = s.controller.SearchByTag(context.Background(), i.TriggerID, tag, indexToReturn, i.View.CallbackID, teamID, &externalViewID)
		if err != nil {
			s.logger.Error().Msgf("%v", err)
			restModel.ErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		c.String(http.StatusOK, "Shortcut initiated")
	}
}

func (s *slackHandler) slashCommandUsed() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req restModel.ShortcutPayload

		requestBody, err := io.ReadAll(c.Request.Body)
		if err != nil {
			s.logger.Err(err).Msgf("error parsing response :: %v", err.Error())
		}

		parsedBody, err := url.ParseQuery(string(requestBody))
		if err != nil {
			restModel.ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		decoder := schema.NewDecoder()
		err = decoder.Decode(&req, parsedBody)
		if err != nil {
			// Handle error;
			s.logger.Err(err).Msg("e don happen")
		}

		if len(req.Text) < 1 {
			// they did not pass anything else asides the slash command
			err = s.controller.ShowSearchModal(context.Background(), req.TriggerID, req.ChannelID, req.TeamID)
		} else {
			// something else was passed asides the slash command
			tag := req.Text
			err = s.controller.SearchByTag(context.Background(), req.TriggerID, tag, "0", req.ChannelID, req.TeamID, nil)
		}
		if err != nil {
			s.logger.Error().Msgf("%v", err)
			// c.String(http.StatusBadRequest, err.Error())
			return
		}

		// restModel.OkResponse(c, http.StatusOK, "Slash command initiated", "response")
	}
}

func (s *slackHandler) saveAuthDetails() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			err error
			req model.SlackAuthDetails
		)

		requestID := requestid.Get(c)
		endpoint := c.FullPath()

		log := s.logger.With().Str(helper.LogEndpointLevel, endpoint).
			Str(helper.LogStrRequestIDLevel, requestID).Logger()

		if err = c.ShouldBind(&req); err != nil {
			log.Err(err).Msg("bad request")
			restModel.ErrorResponse(c, http.StatusBadRequest, err.Error())
			c.Abort()
			return
		}

		err = restModel.ValidateRequest(req)
		if err != nil {
			log.Err(err).Msg("saveAuthDetails ValidateRequest failed")
			restModel.ErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		log.Info().Interface(helper.LogStrPayloadLevel, req).Msg("payload received")

		err = s.controller.SaveAuthDetails(context.Background(), req)
		if err != nil {
			// there is an error.
			restModel.ErrorResponse(c, http.StatusBadRequest, err.Error())
			c.Abort()
			return
		}

		log.Info().Interface(helper.LogStrResponseLevel, true).Msg("response returned")

		c.Abort()
		restModel.OkResponse(c, http.StatusOK, "Details saved successfully", true)
	}
}
