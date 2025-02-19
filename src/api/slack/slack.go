// Package media houses all media related APIs
package media

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"

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
	slackGroup.POST("/events", slack.eventListener())
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
		err := s.controller.SendMessage(context.Background(), "", "", "")
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
		log := s.logger.With().Str(helper.LogEndpointLevel, c.FullPath()).Logger()
		requestBody, err := io.ReadAll(c.Request.Body)
		if err != nil {
			log.Err(err).Msg("io.ReadAll failed")
			restModel.ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		parsedBody, err := url.ParseQuery(string(requestBody))
		if err != nil {
			log.Err(err).Msg("url.ParseQuery failed")
			restModel.ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		var (
			i   slack.InteractionCallback
			tag string
		)
		err = json.Unmarshal([]byte(parsedBody["payload"][0]), &i)
		if err != nil {
			log.Err(err).Msg("json.Unmarshal failed")
			restModel.ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		teamID := i.Team.ID
		channelID := i.Channel.ID
		responseURL := i.ResponseURL
		userID := i.User.ID

		switch i.Type {
		case model.ShortcutType:
			err = s.controller.ShowSearchModal(context.Background(), i.TriggerID, i.CallbackID, teamID)
			if err != nil {
				log.Err(err).Msg("controller.ShowSearchModal failed")
				restModel.ErrorResponse(c, http.StatusBadRequest, err.Error())
				return
			}

			c.String(http.StatusOK, "success")
		case model.SubmissionViewType:
			if len(i.View.Blocks.BlockSet) > 1 && i.View.Blocks.BlockSet[1].BlockType() == model.BlockTypeImage {
				// they actually wanna send the message. Let us proceed
				var details Block

				err = mapstructure.Decode(i.View.Blocks.BlockSet[1], &details)
				if err != nil {
					log.Err(err).Msg("mapstructure.Decode failed")
					restModel.ErrorResponse(c, http.StatusBadRequest, err.Error())
					return
				}

				channelToSendSticker := i.View.CallbackID
				sticker := model.StickerBlockMetadata{
					Tag:    details.Title.Text,
					ImgURL: details.ImageURL,
				}

				err = s.controller.SendSticker(context.Background(), teamID, userID, channelToSendSticker, responseURL, sticker)
				if err != nil {
					log.Err(err).Msg("controller.SendSticker failed")
					restModel.ErrorResponse(c, http.StatusBadRequest, err.Error())
					return
				}

				c.JSON(http.StatusOK, struct {
					ResponseAction string `json:"response_action"`
				}{
					ResponseAction: "clear",
				})

				return
			}

			// this is the initial search
			if i.View.PrivateMetadata == model.InitialDataSearchID {
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
						var sticker model.StickerBlockMetadata
						if err := json.Unmarshal([]byte(blockValue), &sticker); err != nil {
							log.Err(err).Msg("json.Unmarshal failed")
							restModel.ErrorResponse(c, http.StatusBadRequest, err.Error())
							return
						}

						err = s.controller.SendSticker(context.Background(), teamID, userID, channelID, responseURL, sticker)
						if err != nil {
							log.Err(err).Msg("controller.SendSticker failed")
							restModel.ErrorResponse(c, http.StatusBadRequest, err.Error())
							return
						}

						c.String(http.StatusOK, "success")
						return

					case model.ActionIDShuffleSticker:
						var sticker model.StickerBlockMetadata
						if err := json.Unmarshal([]byte(blockValue), &sticker); err != nil {
							log.Err(err).Msg("json.Unmarshal failed")
							restModel.ErrorResponse(c, http.StatusBadRequest, err.Error())
							return
						}
						err = s.controller.ShuffleSticker(context.Background(), teamID, userID, channelID, responseURL, sticker)
						if err != nil {
							log.Err(err).Msg("controller.ShuffleSticker failed")
							restModel.ErrorResponse(c, http.StatusBadRequest, err.Error())
							return
						}
						c.String(http.StatusOK, "success")
						return

					case model.ActionIDCancelSticker:
						if err := s.controller.CancelSticker(context.Background(), teamID, channelID, responseURL); err != nil {
							log.Err(err).Msg("controller.CancelSticker failed")
							restModel.ErrorResponse(c, http.StatusBadRequest, err.Error())
							return
						}
						c.String(http.StatusOK, "success")
						return
					}
				}
			} else {
				tag = i.View.PrivateMetadata
			}

		default:
			// Note there might be a better way to get this info, but I figured this structure out from looking at the json response
			tag = i.View.State.Values["Tag"]["tag"].Value
			// TODO: this fails when we try to run it in threads.
		}

		externalViewID := i.View.ExternalID
		teamID = i.View.TeamID
		var sticker model.StickerBlockMetadata
		for _, v := range i.ActionCallback.BlockActions {
			if err := json.Unmarshal([]byte(v.Value), &sticker); err != nil {
				log.Err(err).Msg("json.Unmarshal failed")
				restModel.ErrorResponse(c, http.StatusBadRequest, err.Error())
				return
			}
		}

		if sticker.Tag == "" {
			sticker.Tag = tag
		}

		err = s.controller.ShowSearchResultModal(context.Background(), i.TriggerID, i.View.CallbackID, teamID, sticker, &externalViewID)
		if err != nil {
			log.Err(err).Msg("controller.ShowSearchResultModal failed")
			restModel.ErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		c.String(http.StatusOK, "Shortcut initiated")
	}
}

func (s *slackHandler) slashCommandUsed() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := requestid.Get(c)
		log := s.logger.With().
			Str(helper.LogEndpointLevel, c.FullPath()).
			Str(helper.LogStrRequestIDLevel, requestID).
			Logger()

		var req restModel.ShortcutPayload

		requestBody, err := io.ReadAll(c.Request.Body)
		if err != nil {
			log.Err(err).Msg("io.ReadAll failed")
			restModel.ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		parsedBody, err := url.ParseQuery(string(requestBody))
		if err != nil {
			log.Err(err).Msg("url.ParseQuery failed")
			restModel.ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		decoder := schema.NewDecoder()
		err = decoder.Decode(&req, parsedBody)
		if err != nil {
			log.Err(err).Msg("decoder.Decode failed.")
			return
		}

		if len(req.Text) < 1 {
			// they did not pass anything else asides the slash command
			if err = s.controller.ShowSearchModal(context.Background(), req.TriggerID, req.ChannelID, req.TeamID); err != nil {
				log.Err(err).Msg("controller.ShowSearchModal failed")
				c.String(http.StatusBadRequest, err.Error())
				return
			}
		} else {
			if err = s.controller.GetStickerSearchResult(context.Background(), req.ChannelID, req.TeamID, req.UserID, req.Text, nil); err != nil {
				log.Err(err).Msg("controller.GetStickerSearchResult failed.")
				c.String(http.StatusBadRequest, err.Error())
				return
			}
		}

		// restModel.OkResponse(c, http.StatusOK, "Slash command initiated", "response")
	}
}

func (s *slackHandler) saveAuthDetails() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := requestid.Get(c)
		log := s.logger.With().
			Str(helper.LogEndpointLevel, c.FullPath()).
			Str(helper.LogStrRequestIDLevel, requestID).
			Logger()

		var req model.SlackAuthDetails

		if err := c.ShouldBind(&req); err != nil {
			log.Err(err).Msg("bad request")
			restModel.ErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		if err := restModel.ValidateRequest(req); err != nil {
			log.Err(err).Msg("saveAuthDetails ValidateRequest failed")
			restModel.ErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		log.Info().Interface(helper.LogStrPayloadLevel, true).Msg("payload received")

		if err := s.controller.SaveAuthDetails(context.Background(), req); err != nil {
			log.Err(err).Msg("SaveAuthDetails failed")
			restModel.ErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		log.Info().Interface(helper.LogStrResponseLevel, true).Msg("response returned")
		restModel.OkResponse(c, http.StatusOK, "Details saved successfully", true)
	}
}

func (s *slackHandler) eventListener() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := requestid.Get(c)
		log := s.logger.With().
			Str(helper.LogEndpointLevel, c.FullPath()).
			Str(helper.LogStrRequestIDLevel, requestID).
			Logger()

		var req restModel.SlackEventCallback

		requestBody, err := io.ReadAll(c.Request.Body)
		if err != nil {
			log.Err(err).Msg("io.ReadAll failed")
			restModel.ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		if err := json.Unmarshal(requestBody, &req); err != nil {
			log.Err(err).Msg("json.Unmarshal failed")
			restModel.ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		if req.Type == "url_verification" {
			c.JSON(http.StatusOK, gin.H{"challenge": req.Challenge})
			return
		}

		if req.Event.Type == "app_mention" {
			// Example message text -> "<@U0LAN0Z89> blah blah"
			// regexp to match mentions in the format <@USER_ID>
			re := regexp.MustCompile(`<@([A-Za-z0-9]+)>`)

			// replace all mentions with an empty string
			textWithoutMention := re.ReplaceAllString(req.Event.Text, "")

			// Optionally, trim any extra spaces that might be left
			textWithoutMention = strings.TrimSpace(textWithoutMention)

			channelID := req.Event.Channel
			teamID := req.TeamID
			userID := req.Event.User

			if err = s.controller.GetStickerSearchResult(context.Background(), channelID, teamID, userID, textWithoutMention, &req.Event.ThreadTS); err != nil {
				log.Err(err).Msg("controller.GetStickerSearchResult failed.")
				c.String(http.StatusBadRequest, err.Error())
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{"status": "event received"})
	}
}
