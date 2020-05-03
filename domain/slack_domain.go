package domain

import (
	"github.com/GulshanArora7/awsChatbot/appcontext"
	"github.com/slack-go/slack"
)

// SlackRepository interface
type SlackRepository interface {
	appcontext.Component

	// EphemeralMenuMessage func send a message with attachment using channelID, and return an error
	EphemeralMenuMessage(channel string, attachment slack.Attachment) error

	// EC2ephemeralMessage func send a message with attachment using channelID, and textMessage and return an error
	EC2ephemeralMessage(channel string, message []EC2Dictionary) error

	// S3ephemeralMessage func send a message with attachment using channelID, and textMessage and return an error
	S3ephemeralMessage(channel string, message []S3Dictionary) error

	// SGephemeralMessage func send a message with attachment using channelID, and textMessage and return an error
	SGephemeralMessage(channel string, message []SGDictionary) error

	// ELBv1ephemeralMessage func send a message with attachment using channelID, and textMessage and return an error
	ELBv1ephemeralMessage(channel string, message []ELBv1Dictionary) error

	// ELBv2ephemeralMessage func send a message with attachment using channelID, and textMessage and return an error
	ELBv2ephemeralMessage(channel string, message []ELBv2Dictionary) error

	// RDSephemeralMessage func send a message with attachment using channelID, and textMessage and return an error
	RDSephemeralMessage(channel string, message []RDSDictionary) error

	// OpenDialogMenu func send a message with triggerID and dialogue box
	OpenDialogMenu(triggerID string, dialogue slack.Dialog) error

	// DialogueAckMessageFunc func send a ack message to opened dialogue box
	DialogueAckMessageFunc(channel string, text string) error

	// BlankResultSlackMsg func send a slack message if in case result is blank
	BlankResultSlackMsg(channel string, message string) error
}

// GetSlackRepository func return SlackRepository interface
func GetSlackRepository() SlackRepository {
	return appcontext.Current.Get(appcontext.SlackRepository).(SlackRepository)
}

// EC2Dictionary interface
type EC2Dictionary map[string]interface{}

// S3Dictionary interface
type S3Dictionary map[string]interface{}

// SGDictionary interface
type SGDictionary map[string]interface{}

// ELBv1Dictionary interface
type ELBv1Dictionary map[string]interface{}

// ELBv2Dictionary interface
type ELBv2Dictionary map[string]interface{}

// RDSDictionary interface
type RDSDictionary map[string]interface{}
