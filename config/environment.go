package config

import (
	"os"
)

var (
	// AwsRegion to be used
	AwsRegion string
	// Port to be listened by application
	Port string
	// SlackBotToken string
	SlackBotToken string
	// SlackBotSigningSecret string
	SlackBotSigningSecret string
	// SlackChanneliD string
	SlackChanneliD string
	// SlackAWSSlashCommand string
	SlackAWSSlashCommand string
	// DebugAWSRequests string
	DebugAWSRequests string
)

// GetEnv func return a default value if missing an Environment variable
func GetEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

func init() {
	AwsRegion = GetEnv("AWS_REGION", "eu-central-1")
	Port = GetEnv("AWSBOT_PORT", "9090")
	DebugAWSRequests = GetEnv("AWSBOT_DEBUG_AWS_REQUESTS", "false")
	SlackAWSSlashCommand = GetEnv("AWSBOT_SLASH_COMMAND", "/awschatbot")
	SlackBotToken = GetEnv("AWSBOT_SLACK_TOKEN", "disabled")
	SlackBotSigningSecret = GetEnv("AWSBOT_SLACK_SIGNING_SECRET", "disabled")
	SlackChanneliD = GetEnv("AWSBOT_SLACK_CHANNEL", "disabled")
}
