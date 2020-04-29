package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/GulshanArora7/awsChatbot/config"
	"github.com/GulshanArora7/awsChatbot/domain"
	"github.com/GulshanArora7/awsChatbot/gateway/slackclient"
	"github.com/GulshanArora7/awsChatbot/usecase"

	"github.com/labstack/echo"
	"github.com/slack-go/slack"
)

const (
	// Disabled const is used to check if a config variable is not configured
	disabled         string = "disabled"
	interActivemsg          = "interactive_message"
	dialogSubmission        = "dialog_submission"
)

// SlackMessageEvents func receive an post from slack slash integration
func SlackMessageEvents(context echo.Context) error {
	if config.SlackBotToken == disabled || config.SlackBotSigningSecret == disabled {
		return context.JSON(http.StatusNotImplemented, nil)
	}
	var bodyBytes []byte
	if context.Request().Body == nil {
		return context.JSON(http.StatusBadRequest, nil)
	}
	bodyBytes, _ = ioutil.ReadAll(context.Request().Body)
	// Restore the io.ReadCloser to its original state
	context.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	bodyString := string(bodyBytes)
	// Copy Forms to Struct
	data := new(slack.SlashCommand)
	data.Token = context.FormValue("token")
	data.TeamID = context.FormValue("team_id")
	data.TeamDomain = context.FormValue("team_domain")
	data.EnterpriseID = context.FormValue("enterprise_id")
	data.EnterpriseName = context.FormValue("enterprise_name")
	data.ChannelID = context.FormValue("channel_id")
	data.ChannelName = context.FormValue("channel_name")
	data.UserID = context.FormValue("user_id")
	data.UserName = context.FormValue("user_name")
	data.Command = context.FormValue("command")
	data.Text = context.FormValue("text")
	data.ResponseURL = context.FormValue("response_url")
	data.TriggerID = context.FormValue("trigger_id")
	// Slack Headers
	slackRequestTimestamp := context.Request().Header.Get("X-Slack-Request-Timestamp")
	slackSignature := context.Request().Header.Get("X-Slack-Signature")
	basestring := fmt.Sprintf("v0:%s:%s", slackRequestTimestamp, bodyString)
	verifier := usecase.ValidateSlackBot(slackSignature, basestring, config.SlackBotSigningSecret)
	if !verifier {
		return context.JSON(http.StatusForbidden, nil)
	}
	if data.ChannelID != config.SlackChanneliD {
		return context.JSON(http.StatusForbidden, nil)
	}
	if data.Command != config.SlackAWSSlashCommand {
		return context.JSON(http.StatusForbidden, nil)
	}
	_, err := usecase.ParseSlashCommand(data)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}
	return err
}

// SlackReplyEvents func
func SlackReplyEvents(context echo.Context) (err error) {
	repo := domain.GetSlackRepository()
	data := new(slack.InteractionCallback)
	payload := context.FormValue("payload")
	err = json.Unmarshal([]byte(payload), &data)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}
	switch data.Type {
	case interActivemsg:
		resp, err := slackclient.SlackActionEvent(data, context)
		if err != nil {
			return context.JSON(http.StatusBadRequest, err)
		}
		return context.JSON(http.StatusOK, resp)
	case dialogSubmission:
		t := "Your Request has been accepted..Please wait while processing..!!"
		e := repo.DialogueAckMessageFunc(data.Channel.ID, t)
		if e != nil {
			log.Fatal("Failed to Send Dialogue Submission Acknowledgement..!!!!", e)
		}
		err := slackclient.SlackDialogSubmissionEvent(data, context)
		if err != nil {
			return context.JSON(http.StatusBadRequest, err)
		}
	}
	return nil
}
