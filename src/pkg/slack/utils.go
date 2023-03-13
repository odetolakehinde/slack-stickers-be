package slack

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/slack-go/slack"

	"github.com/odetolakehinde/slack-stickers-be/src/model"
)

func generateSearchModalRequest(channelID string) slack.ModalViewRequest {
	// Create a ModalViewRequest with a header and two inputs
	titleText := slack.NewTextBlockObject("plain_text", "Slack Stickers ", false, false)
	closeText := slack.NewTextBlockObject("plain_text", "Close", false, false)
	submitText := slack.NewTextBlockObject("plain_text", "Submit", false, false)

	headerText := slack.NewTextBlockObject("mrkdwn", "Your chats don't have to be boring! Send a sticker to make things fun!", false, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)

	tagText := slack.NewTextBlockObject("plain_text", "Search for a sticker", false, false)
	tagPlaceholder := slack.NewTextBlockObject("plain_text", "Search slack stickers", false, false)
	tagElement := slack.NewPlainTextInputBlockElement(tagPlaceholder, "tag")
	// Notice that blockID is a unique identifier for a block
	tag := slack.NewInputBlock("Tag", tagText, nil, tagElement)

	blocks := slack.Blocks{
		BlockSet: []slack.Block{
			headerSection,
			tag,
		},
	}

	return slack.ModalViewRequest{
		Type:            model.ModalViewType,
		Title:           titleText,
		Blocks:          blocks,
		Close:           closeText,
		Submit:          submitText,
		PrivateMetadata: model.InitialDataSearchID,
		CallbackID:      channelID, // we use the channel ID
		//ClearOnClose:    false,
		//NotifyOnClose:   false,
		ExternalID: uuid.New().String(),
	}
}

func generateSearchResultModal(imageURL, altText, tag, channelID string, indexToReturn int) slack.ModalViewRequest {
	// Create a ModalViewRequest with a header and two inputs
	titleText := slack.NewTextBlockObject("plain_text", "Slack Stickers ", false, false)
	closeText := slack.NewTextBlockObject("plain_text", "Close", false, false)
	submitText := slack.NewTextBlockObject("plain_text", "Send Sticker", false, false)

	headerText := slack.NewTextBlockObject("mrkdwn", "Not what you have in mind? Switch it", false, false)
	btnText := slack.NewTextBlockObject("plain_text", "Shuffle", false, false)
	btn := slack.NewButtonBlockElement(model.ActionIDShuffle, fmt.Sprintf("%d", indexToReturn+1), btnText)
	accessory := slack.Accessory{
		ButtonElement: btn,
	}
	headerSection := slack.NewSectionBlock(headerText, nil, &accessory)

	//imageText := slack.NewTextBlockObject("mrkdwn", tag, false, false)
	image := slack.NewImageBlock(imageURL, altText, "image-block-id", nil)

	blocks := slack.Blocks{
		BlockSet: []slack.Block{
			headerSection,
			image,
		},
	}

	return slack.ModalViewRequest{
		Type:            model.ModalViewType,
		Title:           titleText,
		Blocks:          blocks,
		Close:           closeText,
		Submit:          submitText,
		PrivateMetadata: tag,
		CallbackID:      channelID, // we are using the channel ID for this
		//ClearOnClose:    false,
		//NotifyOnClose:   false,
		ExternalID: uuid.New().String(),
	}
}
