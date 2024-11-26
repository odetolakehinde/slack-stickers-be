package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/odetolakehinde/slack-stickers-be/src/model"
)

// SendMessage sends a sticker to the specified channel
func (c *Controller) SendMessage(ctx context.Context, channelID, imageURL, teamID string) error {
	c.logger.Info().
		Str("imageURL", imageURL).
		Str("channelID", channelID).
		Msg("sending sticker")
	slackService := c.getSlackService(ctx, teamID)
	return slackService.SendMessage(ctx, channelID, imageURL)
}

// ShowSearchModal shows up the search modal
func (c *Controller) ShowSearchModal(ctx context.Context, triggerID, channelID, teamID string) error {
	slackService := c.getSlackService(ctx, teamID)
	return slackService.ShowSearchModal(ctx, triggerID, channelID)
}

// SearchByTag shows up the search modal
func (c *Controller) SearchByTag(ctx context.Context, triggerID, tag, countToReturn, channelID, teamID string, externalViewID *string) error {
	slackService := c.getSlackService(ctx, teamID)
	result, totalCount, err := c.cloudinary.SearchByTag(ctx, tag)
	if err != nil {
		c.logger.Err(err).Msg("cloudinary.SearchByTag failed")
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
			c.logger.Err(err).Msg("slackService.ShowSearchResultModal failed")
			return err
		}
	}

	return nil
}

// SaveAuthDetails handles function to safe the authorization details to redis
func (c *Controller) SaveAuthDetails(ctx context.Context, authDetails model.SlackAuthDetails) error {
	// change to string
	bytes, err := json.Marshal(authDetails)
	if err != nil {
		c.logger.Err(err).Msg("json.Marshal failed")
		return err
	}

	// save to redis
	err = c.store.SetValue(ctx, authDetails.Team.ID, string(bytes), 0)
	if err != nil {
		c.logger.Err(err).Msg("store.SetValue failed")
		return err
	}

	return nil
}

// GetStickerSearchResult shows up the search result
func (c *Controller) GetStickerSearchResult(ctx context.Context, channelID, teamID, userID, text string) error {
	slackService := c.getSlackService(ctx, teamID)

	result, totalCount, err := c.cloudinary.SearchByTag(ctx, text)
	if err != nil {
		c.logger.Err(err).Msg("cloudinary.SearchByTag failed")
		return err
	}

	if totalCount < 1 {
		return fmt.Errorf("not found")
	}

	imageURL := result[0].URL

	return slackService.ShowStickerPreview(ctx, userID, channelID, text, imageURL)
}

// CancelSticker to close sticker preview block
func (c *Controller) CancelSticker(ctx context.Context, teamID, channelID, responseURL string) error {
	slackService := c.getSlackService(ctx, teamID)
	return slackService.CancelStickerPreview(ctx, channelID, responseURL)
}

// SendSticker to send sticker
func (c *Controller) SendSticker(ctx context.Context, teamID, userID, channelID, responseURL string, sticker model.StickerBlockActionValue) error {
	c.logger.Info().
		Str("channelID", channelID).
		Msg("sending sticker")
	slackService := c.getSlackService(ctx, teamID)
	return slackService.SendStickerToChannel(ctx, userID, channelID, responseURL, sticker)
}

// ShuffleSticker to shuffle sticker
func (c *Controller) ShuffleSticker(ctx context.Context, teamID, userID, channelID, responseURL string, sticker model.StickerBlockActionValue) error {
	slackService := c.getSlackService(ctx, teamID)
	result, totalCount, err := c.cloudinary.SearchByTag(ctx, sticker.Tag)
	if err != nil {
		c.logger.Err(err).Msg("cloudinary.SearchByTag failed")
		return err
	}

	if sticker.Index >= totalCount {
		sticker.Index = 0
	}

	if len(result) > 0 {
		sticker.ImgURL = result[sticker.Index].URL
		if err := slackService.ShuffleStickerPreview(ctx, userID, channelID, responseURL, sticker); err != nil {
			return err
		}
	}

	return nil
}
