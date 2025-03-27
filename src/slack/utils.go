package slack

import (
	"encoding/json"

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
		PrivateMetadata: model.SlackShortcutCallbackID,
		CallbackID:      channelID, // we use the channel ID
		// ClearOnClose:    false,
		// NotifyOnClose:   false,
		ExternalID: uuid.New().String(),
	}
}

func generateSearchResultModal(channelID string, sticker model.StickerBlockMetadata, isShuffle bool) slack.ModalViewRequest {
	if isShuffle {
		sticker.Index++
	} else {
		sticker.Index = 0
	}

	jsonByte, _ := json.Marshal(sticker)
	jsonString := string(jsonByte)

	titleText := slack.NewTextBlockObject("plain_text", "Slack Stickers ", false, false)
	closeText := slack.NewTextBlockObject("plain_text", "Close", false, false)
	submitText := slack.NewTextBlockObject("plain_text", "Send Sticker", false, false)
	headerText := slack.NewTextBlockObject("mrkdwn", "Not what you have in mind? Switch it", false, false)
	btnText := slack.NewTextBlockObject("plain_text", "Shuffle", false, false)

	btn := slack.NewButtonBlockElement(model.ActionIDShuffleSticker, jsonString, btnText)
	accessory := slack.Accessory{
		ButtonElement: btn,
	}
	headerSection := slack.NewSectionBlock(headerText, nil, &accessory)

	imageText := slack.NewTextBlockObject(slack.PlainTextType, sticker.Tag, false, false)
	image := slack.NewImageBlock(sticker.ImgURL, sticker.Tag, "image-block-id", imageText)

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
		PrivateMetadata: sticker.Tag,
		CallbackID:      channelID, // we are using the channel ID for this
		ExternalID:      uuid.New().String(),
	}
}

// createStickerPreviewBlock creates a Slack message block containing a sticker preview and action buttons (Send, Shuffle, Cancel).
//
// It also adjusts the sticker's index based on whether the shuffle action is triggered or not.
func createStickerPreviewBlock(sticker model.StickerBlockMetadata, isShuffle bool) slack.Blocks {
	if isShuffle {
		sticker.Index++
	} else {
		sticker.Index = 0
	}

	jsonByte, _ := json.Marshal(sticker)
	jsonString := string(jsonByte)

	blocks := []slack.Block{
		slack.NewImageBlock(
			sticker.ImgURL,
			sticker.Tag,
			model.StickerImageBlockID,
			slack.NewTextBlockObject(slack.PlainTextType, sticker.Tag, false, false),
		),
		slack.NewActionBlock(
			model.StickerActionBlockID,
			slack.NewButtonBlockElement(
				model.ActionIDSendSticker,
				jsonString,
				slack.NewTextBlockObject(slack.PlainTextType, "Send", false, false),
			).WithStyle(slack.StylePrimary),
			slack.NewButtonBlockElement(
				model.ActionIDShuffleSticker,
				jsonString,
				slack.NewTextBlockObject(slack.PlainTextType, "Shuffle", false, false),
			),
			slack.NewButtonBlockElement(
				model.ActionIDCancelSticker,
				"",
				slack.NewTextBlockObject(slack.PlainTextType, "Cancel", false, false),
			).WithStyle(slack.StyleDanger),
		),
	}

	return slack.Blocks{
		BlockSet: blocks,
	}
}

func generateHelpBlocks() slack.Blocks {
	headerText := slack.NewTextBlockObject(slack.MarkdownType, "*Need help using `Stickers For Slack`?*", false, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)

	usageText := slack.NewTextBlockObject(slack.MarkdownType,
		"*Slash Command:*\n"+
			"• `/sticker happy dance`\n\n"+
			"*Mention Format:*\n"+
			"this can be used in threads and channels:\n"+
			"• `@sticker search excited cat`\n"+
			"• `@sticker find dancing dog`\n"+
			"• `@sticker g celebration`\n"+
			"• `@sticker gif party time`\n\n"+
			"*Open Modal:*\n"+
			"• Type `/sticker` and press Enter",
		false, false,
	)

	usageSection := slack.NewSectionBlock(usageText, nil, nil)

	// Buttons: View Docs and Cancel
	viewDocsText := slack.NewTextBlockObject(slack.PlainTextType, "View Docs", false, false)
	viewDocsButton := slack.NewButtonBlockElement(model.ActionIDViewDocs, "", viewDocsText)
	viewDocsButton.URL = DocsURL
	viewDocsButton.WithStyle(slack.StylePrimary)

	cancelText := slack.NewTextBlockObject(slack.PlainTextType, "Cancel", false, false)
	cancelButton := slack.NewButtonBlockElement(model.ActionIDCancelSticker, "", cancelText)
	cancelButton.WithStyle(slack.StyleDanger)

	actionBlock := slack.NewActionBlock(
		model.StickerActionBlockID,
		viewDocsButton,
		cancelButton,
	)

	blocks := []slack.Block{
		headerSection,
		usageSection,
		actionBlock,
	}

	return slack.Blocks{
		BlockSet: blocks,
	}
}
