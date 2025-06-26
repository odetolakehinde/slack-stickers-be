// Package model defines all the models in the system
package model

const (
	// SlackShortcutCallbackID is the callback id used for shortcuts
	SlackShortcutCallbackID = "send_sticker_shortcut"

	// StickerActionBlockID is the block id for the sticker actions button
	StickerActionBlockID = "sticker_action_block-id"
	// StickerImageBlockID is the image block id
	StickerImageBlockID = "sticker_image_block-id"
	// StickerContextBlockID is the context block id
	StickerContextBlockID = "sticker_context_block-id"

	// ActionIDShuffleSticker shuffle sticker action ID
	ActionIDShuffleSticker = "shuffle-sticker"
	// ActionIDCancelSticker cancel sticker action ID
	ActionIDCancelSticker = "cancel-sticker"
	// ActionIDSendSticker send sticker action ID
	ActionIDSendSticker = "send-sticker"
	// ActionIDViewDocs view docs action ID
	ActionIDViewDocs = "view-docs"

	// MaxResults during search from cloudinary
	MaxResults = 50

	// SubmissionViewType for View submission interaction
	SubmissionViewType = "view_submission"

	// ShortcutType for shortcuts used
	ShortcutType = "shortcut"

	// MessageActionType for message shortcut or action triggered from message context
	MessageActionType = "message_action"

	// BlockActionsViewType for View block actions interaction
	BlockActionsViewType = "block_actions"

	// BlockTypeImage for image block type
	BlockTypeImage = "image"

	// ModalViewType for modal view type
	ModalViewType = "modal"

	// EventTypeAppMention for when the bot is mentioned
	EventTypeAppMention = "app_mention"

	// EventTypeAppUninstalled for when the bot is uninstalled/removed from a workspace
	EventTypeAppUninstalled = "app_uninstalled"

	// EventTypeTokensRevoked for when the bot token is revoked
	EventTypeTokensRevoked = "tokens_revoked"

	// SlackCallbackEventURLVerification for verifying the url for event listener handler
	// https://api.slack.com/events/url_verification
	SlackCallbackEventURLVerification = "url_verification"
)
