package media

type (
	uploadRequest struct {
		Name string `json:"name" validate:"required"`
	}
)
