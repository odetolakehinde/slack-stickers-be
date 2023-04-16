package model

type (
	// SlackAuthDetails schema
	SlackAuthDetails struct {
		TeamID    string `json:"team_id" validate:"required"`
		UserID    string `json:"user_id" validate:"required"`
		Token     string `json:"token" validate:"required"`
		TokenType string `json:"token_type" validate:"required"`
	}
)
