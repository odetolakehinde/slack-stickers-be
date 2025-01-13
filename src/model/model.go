// Package model defines all the models in the system
package model

const (
	// InitialDataSearchID for when the initial search is done
	InitialDataSearchID = "initial-data-search"

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

	// MaxResults during search from cloudinary
	MaxResults = 50

	// SubmissionViewType for View submission interaction
	SubmissionViewType = "view_submission"

	// ShortcutType for shortcuts used
	ShortcutType = "shortcut"

	// BlockActionsViewType for View block actions interaction
	BlockActionsViewType = "block_actions"

	// BlockTypeImage for image block type
	BlockTypeImage = "image"

	// ModalViewType for modal view type
	ModalViewType = "modal"
)
