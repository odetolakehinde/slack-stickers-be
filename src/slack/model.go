package slack

const (
	// HeaderText text to be used in the header of the slack message
	HeaderText = "Stickers, made for Slack!"
	// FooterText text to be used in the footer of the slack message
	FooterText = "powered by slackstickers.com"
	// ColorText text to be used in the color of the slack message
	ColorText = "#F26722"
	// IconURL url to the icon of the slack message
	IconURL = ""
)

type (
	// AuthResponse schema for auth response
	AuthResponse struct {
		Ok         bool   `json:"ok" mapstructure:"ok"`
		Error      string `json:"error,omitempty" mapstructure:"error,omitempty"`
		AppID      string `json:"app_id,omitempty" mapstructure:"app_id,omitempty"`
		AuthedUser struct {
			ID          string `json:"id,omitempty" mapstructure:"id,omitempty"`
			Scope       string `json:"scope,omitempty" mapstructure:"scope,omitempty"`
			AccessToken string `json:"access_token,omitempty" mapstructure:"access_token,omitempty"`
			TokenType   string `json:"token_type,omitempty" mapstructure:"token_type,omitempty"`
		} `json:"authed_user,omitempty" mapstructure:"authed_user,omitempty"`
		Team struct {
			ID string `json:"id,omitempty" mapstructure:"id,omitempty"`
		} `json:"team,omitempty" mapstructure:"team,omitempty"`
		Enterprise          interface{} `json:"enterprise,omitempty" mapstructure:"enterprise,omitempty"`
		IsEnterpriseInstall bool        `json:"is_enterprise_install,omitempty" mapstructure:"is_enterprise_install,omitempty"`
	}
)
