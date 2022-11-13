package model

import (
	"time"

	"github.com/google/uuid"
)

type (
	// Sticker struct representing all details about a sticker
	Sticker struct {
		ID           uuid.UUID `json:"id,omitempty"`
		Name         string    `json:"name,omitempty"`
		URL          string    `json:"url,omitempty"`
		Folder       string    `json:"folder,omitempty"`
		AssetFolder  string    `json:"asset_folder,omitempty"`
		ResourceType string    `json:"resource_type,omitempty"`
		Format       string    `json:"format,omitempty"`
		Timestamp    time.Time `json:"timestamp,omitempty"`
		Tags         []string  `json:"tags,omitempty"`
	}
)
