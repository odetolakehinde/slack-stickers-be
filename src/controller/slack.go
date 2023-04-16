package controller

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/odetolakehinde/slack-stickers-be/src/model"
)

// SendSticker sends a sticker to the specified channel
func (c *Controller) SendSticker(ctx context.Context, channelID, imageURL string) error {
	return c.slackService.SendSticker(ctx, channelID, imageURL)
}

// ShowSearchModal shows up the search modal
func (c *Controller) ShowSearchModal(ctx context.Context, triggerID, channelID string) error {
	return c.slackService.ShowSearchModal(ctx, triggerID, channelID)
}

// SearchByTag shows up the search modal
func (c *Controller) SearchByTag(ctx context.Context, triggerID, tag, countToReturn, channelID string, externalViewID *string) error {
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
		err = c.slackService.ShowSearchResultModal(ctx, triggerID, response.URL, response.Name, tag, channelID, externalViewID, indexToReturn)
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
		panic(err)
	}

	// save to redis
	err = c.store.SetValue(ctx, authDetails.TeamID, string(bytes), 0)
	if err != nil {
		c.logger.Err(err).Msg("store.SetValue failed")
		return err
	}

	return nil
}
