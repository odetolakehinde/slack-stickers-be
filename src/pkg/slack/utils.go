package slack

import "github.com/slack-go/slack"

func generateSearchModalRequest() slack.ModalViewRequest {
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

	var modalRequest slack.ModalViewRequest
	modalRequest.Type = slack.ViewType("modal")
	modalRequest.Title = titleText
	modalRequest.Close = closeText
	modalRequest.Submit = submitText
	modalRequest.Blocks = blocks
	return modalRequest
}

func generateSearchResultModal() slack.ModalViewRequest {
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

	var modalRequest slack.ModalViewRequest
	modalRequest.Type = slack.ViewType("modal")
	modalRequest.Title = titleText
	modalRequest.Close = closeText
	modalRequest.Submit = submitText
	modalRequest.Blocks = blocks
	return modalRequest
}
