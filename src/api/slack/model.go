package media

type (
	uploadRequest struct {
		Name string `json:"name" validate:"required"`
	}

	// Block image schema
	Block struct {
		Type     string      `mapstructure:"Type" json:"Type"`
		ImageURL string      `json:"ImageURL"`
		AltText  string      `json:"AltText"`
		BlockID  string      `json:"BlockID"`
		Title    interface{} `json:"Title"`
	}
)
