package model

import "github.com/slack-go/slack"

// ShortcutPayload represents the payload received when a slack global shortcut is used:
type ShortcutPayload struct {
	Type                string `schema:"type"`
	Token               string `schema:"token"`
	TeamID              string `schema:"team_id"`
	TeamDomain          string `schema:"team_domain"`
	ChannelID           string `schema:"channel_id"`
	ChannelName         string `schema:"channel_name"`
	UserID              string `schema:"user_id"`
	UserName            string `schema:"user_name"`
	Command             string `schema:"command"`
	Text                string `schema:"text"`
	APIAppID            string `schema:"api_app_id"`
	IsEnterpriseInstall string `schema:"is_enterprise_install"`
	ResponseURL         string `schema:"response_url"`
	CallbackID          string `schema:"callback_id"`
	TriggerID           string `schema:"trigger_id"`
}

// SlackEventCallback represents the top-level structure for the Slack event payload.
type SlackEventCallback struct {
	Token              string                     `json:"token"`
	Challenge          *string                    `json:"challenge"`
	TeamID             string                     `json:"team_id"`
	APIAppID           string                     `json:"api_app_id"`
	Event              Event                      `json:"event"`
	Type               string                     `json:"type"`
	EventID            string                     `json:"event_id"`
	EventTime          int64                      `json:"event_time"`
	Authorizations     []slack.EventAuthorization `json:"authorizations"`
	IsExtSharedChannel bool                       `json:"is_ext_shared_channel"`
	EventContext       string                     `json:"event_context"`
}

// Event event
type Event struct {
	Type         string `json:"type"`
	User         string `json:"user"`
	Text         string `json:"text"`
	Timestamp    string `json:"ts"`
	Channel      string `json:"channel"`
	EventTS      string `json:"event_ts"`
	ClientMsgID  string `json:"client_msg_id"`
	Team         string `json:"team"`
	ThreadTS     string `json:"thread_ts"`
	ParentUserID string `json:"parent_user_id"`
}
