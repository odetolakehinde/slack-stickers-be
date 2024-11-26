// Package env list all environment variables key in the app
package env

const (

	// ServerPort server port
	ServerPort = "SERVER_PORT"

	// IsSandboxMode if app is on Sandbox mode
	IsSandboxMode = "IS_SANDBOX_MODE"

	// RedisServerAddress redis connection address
	RedisServerAddress = "REDIS_SERVER_ADDRESS"
	// RedisServerUsername redis connection username
	RedisServerUsername = "REDIS_SERVER_USERNAME"
	// RedisServerPassword redis connection password
	RedisServerPassword = "REDIS_SERVER_PASSWORD"
	// RedisTLSEnabled to check if tls is enabled
	RedisTLSEnabled = "REDIS_TLS_ENABLED"

	// CloudinaryCloudName cloudinary cloud name
	CloudinaryCloudName = "CLOUDINARY_CLOUD_NAME"
	// CloudinaryAPIKey cloudinary api key
	CloudinaryAPIKey = "CLOUDINARY_API_KEY"
	// CloudinarySecretKey cloudinary api secret
	CloudinarySecretKey = "CLOUDINARY_SECRET_KEY"

	// SlackToken slack messaging token
	SlackToken = "SLACK_TOKEN"
	// SlackAppToken for app level token
	SlackAppToken = "SLACK_APP_TOKEN"
)
