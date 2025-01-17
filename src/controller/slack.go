package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/odetolakehinde/slack-stickers-be/src/model"
	"github.com/odetolakehinde/slack-stickers-be/src/pkg/helper"
)

// SendMessage sends a sticker to the specified channel
func (c *Controller) SendMessage(ctx context.Context, channelID, imageURL, teamID string) error {
	log := c.logger.With().Str(helper.LogStrKeyMethod, "SendMessage").Logger()
	slackService, err := c.getSlackService(ctx, teamID)
	if err != nil {
		log.Err(err).Str("teamID", teamID).Msg("failed to get Slack service")
		return err
	}
	log.Info().Str("imageURL", imageURL).Str("channelID", channelID).Msg("sending sticker")
	return slackService.SendMessage(ctx, channelID, imageURL)
}

// ShowSearchModal shows up the search modal
func (c *Controller) ShowSearchModal(ctx context.Context, triggerID, channelID, teamID string) error {
	log := c.logger.With().Str(helper.LogStrKeyMethod, "ShowSearchModal").Logger()
	slackService, err := c.getSlackService(ctx, teamID)
	if err != nil {
		log.Err(err).Str("teamID", teamID).Msg("failed to get Slack service")
		return err
	}
	return slackService.ShowSearchModal(ctx, triggerID, channelID)
}

// SearchByTag shows up the search modal
func (c *Controller) SearchByTag(ctx context.Context, triggerID, tag, countToReturn, channelID, teamID string, externalViewID *string) error {
	log := c.logger.With().Str(helper.LogStrKeyMethod, "SearchByTag").Logger()
	slackService, err := c.getSlackService(ctx, teamID)
	if err != nil {
		log.Err(err).Str("teamID", teamID).Msg("failed to get Slack service")
		return err
	}

	result, totalCount, err := c.cloudinary.SearchByTag(ctx, tag)
	if err != nil {
		log.Err(err).Msg("cloudinary.SearchByTag failed")
		return err
	}

	indexToReturn, _ := strconv.Atoi(countToReturn)

	if indexToReturn >= totalCount {
		// we are at the last one, go back to zero
		indexToReturn = 0
	}

	// for now, send the first result
	if len(result) > 0 {
		response := result[indexToReturn]
		err = slackService.ShowSearchResultModal(ctx, triggerID, response.URL, response.Name, tag, channelID, externalViewID, indexToReturn)
		if err != nil {
			log.Err(err).Msg("slackService.ShowSearchResultModal failed")
			return err
		}
	}

	return nil
}

// SaveAuthDetails handles function to safe the authorization details to redis
func (c *Controller) SaveAuthDetails(ctx context.Context, authDetails model.SlackAuthDetails) error {
	log := c.logger.With().Str(helper.LogStrKeyMethod, "SaveAuthDetails").Logger()

	bytes, err := json.Marshal(authDetails)
	if err != nil {
		log.Err(err).Msg("json.Marshal failed")
		return err
	}

	// save to redis
	err = c.store.SetValue(ctx, authDetails.Team.ID, string(bytes), 0)
	if err != nil {
		log.Err(err).Msg("store.SetValue failed")
		return err
	}

	return nil
}

// GetStickerSearchResult shows up the search result
func (c *Controller) GetStickerSearchResult(ctx context.Context, channelID, teamID, userID, text string) error {
	log := c.logger.With().Str(helper.LogStrKeyMethod, "GetStickerSearchResult").Logger()
	slackService, err := c.getSlackService(ctx, teamID)
	if err != nil {
		log.Err(err).Str("teamID", teamID).Msg("failed to get Slack service")
		return err
	}

	response, err := c.tenor.SearchGifsByQuery(ctx, text, "")
	if err != nil {
		log.Err(err).Msg("tenor.SearchGifsByQuery failed")
		return err
	}

	if len(response.Results) == 0 {
		err := fmt.Errorf("not found")
		log.Err(err).Msg("sticker not found")
		return err
	}

	imageURL := response.Results[0].MediaFormats.Gif.URL

	return slackService.ShowStickerPreview(ctx, userID, channelID, text, imageURL)
}

// CancelSticker to close sticker preview block
func (c *Controller) CancelSticker(ctx context.Context, teamID, channelID, responseURL string) error {
	log := c.logger.With().Str(helper.LogStrKeyMethod, "CancelSticker").Logger()
	slackService, err := c.getSlackService(ctx, teamID)
	if err != nil {
		log.Err(err).Str("teamID", teamID).Msg("failed to get Slack service")
		return err
	}

	return slackService.CancelStickerPreview(ctx, channelID, responseURL)
}

// SendSticker to send sticker
func (c *Controller) SendSticker(ctx context.Context, teamID, userID, channelID, responseURL string, sticker model.StickerBlockActionValue) error {
	log := c.logger.With().Str(helper.LogStrKeyMethod, "SendSticker").Logger()
	slackService, err := c.getSlackService(ctx, teamID)
	if err != nil {
		log.Err(err).Str("teamID", teamID).Msg("failed to get Slack service")
		return err
	}

	log.Info().Str("channelID", channelID).Msg("sending sticker")
	return slackService.SendStickerToChannel(ctx, userID, channelID, responseURL, sticker)
}

// ShuffleSticker to shuffle sticker
func (c *Controller) ShuffleSticker(ctx context.Context, teamID, userID, channelID, responseURL string, sticker model.StickerBlockActionValue) error {
	log := c.logger.With().Str(helper.LogStrKeyMethod, "ShuffleSticker").Logger()
	slackService, err := c.getSlackService(ctx, teamID)
	if err != nil {
		log.Err(err).Str("teamID", teamID).Msg("failed to get Slack service")
		return err
	}

	response, err := c.tenor.SearchGifsByQuery(ctx, sticker.Tag, sticker.Pos)
	if err != nil {
		log.Err(err).Msg("tenor.SearchGifsByQuery failed")
		return err
	}

	totalCount := len(response.Results)

	// If the index is equal to the total count, fetch the next page of results
	if sticker.Index == totalCount {
		response, err = c.tenor.SearchGifsByQuery(ctx, sticker.Tag, response.Next)
		if err != nil {
			log.Err(err).Msg("tenor.SearchGifsByQuery failed when index is == totalCount")
			return err
		}
		sticker.Index = 0           // reset index to 0
		sticker.Pos = response.Next // Set position for the next page of results
	}

	if len(response.Results) > 0 {
		sticker.ImgURL = response.Results[sticker.Index].MediaFormats.Gif.URL
		if err := slackService.ShuffleStickerPreview(ctx, userID, channelID, responseURL, sticker); err != nil {
			log.Err(err).Msg("ShuffleStickerPreview failed")
			return err
		}
	}

	return nil
}
