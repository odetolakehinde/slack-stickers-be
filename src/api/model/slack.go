package model

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
	ApiAppID            string `schema:"api_app_id"`
	IsEnterpriseInstall string `schema:"is_enterprise_install"`
	ResponseURL         string `schema:"response_url"`
	CallbackID          string `schema:"callback_id"`
	TriggerID           string `schema:"trigger_id"`
}
