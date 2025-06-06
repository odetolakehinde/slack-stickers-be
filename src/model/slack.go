package model

// RedisSlackAuthPrefix is the key prefix used for storing Slack authentication details in Redis.
// Each team's authentication data is stored with a key format of "RedisSlackAuthPrefix:team_id".
const RedisSlackAuthPrefix = "slack_auth"

type (
	// SlackAuthDetails schema
	//SlackAuthDetails struct {
	//	TeamID    string `json:"team_id" validate:"required"`
	//	UserID    string `json:"user_id" validate:"required"`
	//	Token     string `json:"token" validate:"required"`
	//	TokenType string `json:"token_type" validate:"required"`
	//}

	// SlackAuthDetails schema
	SlackAuthDetails struct {
		Ok         bool   `json:"ok"`
		Error      string `json:"error,omitempty"`
		AppID      string `json:"app_id"`
		AuthedUser struct {
			ID           string `json:"id"`
			Scope        string `json:"scope"`
			AccessToken  string `json:"access_token"`
			TokenType    string `json:"token_type"`
			RefreshToken string `json:"refresh_token"`
			ExpiresIn    int    `json:"expires_in"`
		} `json:"authed_user"`
		Scope        string `json:"scope"`
		TokenType    string `json:"token_type"`
		AccessToken  string `json:"access_token"`
		BotUserID    string `json:"bot_user_id"`
		RefreshToken string `json:"refresh_token"`
		ExpiresIn    int    `json:"expires_in"`
		Team         struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"team"`
		Enterprise          interface{} `json:"enterprise"`
		IsEnterpriseInstall bool        `json:"is_enterprise_install"`
	}

	// StickerBlockMetadata is a sticker block metadata
	StickerBlockMetadata struct {
		Tag      string  `json:"tag"`
		Index    int     `json:"index"`
		ImgURL   string  `json:"imgURL"`
		Pos      string  `json:"pos"`
		ThreadTS *string `json:"threadTS"`
	}
)
