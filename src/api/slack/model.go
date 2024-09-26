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

	// SlackSendMessageRequest struct to parse incoming Slack request
	SlackSendMessageRequest struct {
		Token       string `json:"token"`
		TeamID      string `json:"team_id"`
		TeamDomain  string `json:"team_domain"`
		ChannelID   string `json:"channel_id"`
		ChannelName string `json:"channel_name"`
		UserID      string `json:"user_id"`
		UserName    string `json:"user_name"`
		Command     string `json:"command"`
		Text        string `json:"text"`
		ResponseURL string `json:"response_url"`
		TriggerID   string `json:"trigger_id"`
		ThreadTS    string `json:"thread_ts"` // Present if the command is called in a thread
	}

	// SlackSendMessageResponse struct for responding back to Slack
	SlackSendMessageResponse struct {
		ResponseType string `json:"response_type"`
		Text         string `json:"text"`
		ThreadTS     string `json:"thread_ts,omitempty"`
	}
)
