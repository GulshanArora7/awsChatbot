package usecase

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"log"
	"regexp"
	"strings"

	"github.com/slack-go/slack"
)

// ValidateSlackBot func to validate auth from Slack Bot
func ValidateSlackBot(signing string, message string, mysigning string) bool {
	mac := hmac.New(sha256.New, []byte(mysigning))
	if _, err := mac.Write([]byte(message)); err != nil {
		log.Printf("[ERROR] ValidateBot mac.Write(%v) failed\n", message)
		return false
	}
	calculatedMAC := "v0=" + hex.EncodeToString(mac.Sum(nil))
	return hmac.Equal([]byte(calculatedMAC), []byte(signing))
}

// ParseSlashCommand func to parse Slash command
func ParseSlashCommand(data *slack.SlashCommand) (slack.Msg, error) {
	var res slack.Msg
	initialText := data.Text
	space := regexp.MustCompile(`\s+`)
	secondaryText := space.ReplaceAllString(initialText, " ")
	parseddata := strings.Split(strings.ToLower(secondaryText), " ")
	requestMap := make(map[string]string)
	requestMap["awsservice"] = parseddata[0]
	var errSlack error
	if errSlack != nil {
		log.Printf("[ERROR] ParseSlashCommand %s", errSlack)
	}
	go SendAwsmenuSlackPost(data.UserID, data.ChannelID, requestMap)
	return res, nil
}
