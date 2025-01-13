package media

import "github.com/slack-go/slack"

type (
	uploadRequest struct { //nolint:unused
		Name string `json:"name" validate:"required"`
	}

	// Block image schema
	Block struct {
		Type     string                `mapstructure:"Type" json:"Type"`
		ImageURL string                `json:"ImageURL"`
		AltText  string                `json:"AltText"`
		BlockID  string                `json:"BlockID"`
		Title    slack.TextBlockObject `json:"Title"`
	}
)
