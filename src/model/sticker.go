package model

import "github.com/google/uuid"

type (
	// Sticker struct representing all details about a sticker
	Sticker struct {
		ID   uuid.UUID
		Name string
		URL  string
	}
)
