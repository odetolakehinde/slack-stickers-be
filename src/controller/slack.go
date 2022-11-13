package controller

import (
	"context"
	"fmt"
)

// SendSticker sends a sticker to the specified channel
func (c *Controller) SendSticker(_ context.Context, channelID, imageURL string) error {
	return c.slackService.SendSticker(channelID, imageURL)
}

// ShowSearchModal shows up the search modal
func (c *Controller) ShowSearchModal(_ context.Context, channelID, triggerID string) error {
	return c.slackService.ShowSearchModal(channelID, triggerID)
}

// SearchByTag shows up the search modal
func (c *Controller) SearchByTag(ctx context.Context, tag string) error {
	result, err := c.cloudinary.SearchByTag(ctx, tag)
	if err != nil {
		c.logger.Err(err).Msg("cloudinary.SearchByTag failed")
	}

	fmt.Println(result)
	return nil
}
