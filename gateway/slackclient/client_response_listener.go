package slackclient

import (
	"fmt"
	"log"
	"strings"

	"github.com/GulshanArora7/awsChatbot/config"
	"github.com/GulshanArora7/awsChatbot/domain"
	"github.com/GulshanArora7/awsChatbot/gateway/awsclient"
	"github.com/GulshanArora7/awsChatbot/utils"
	"github.com/labstack/echo"
	"github.com/slack-go/slack"
)

const (
	// action is used for slack
	actionSelect = "select"
	actionStart  = "start"
	actionCancel = "cancel"
)

var value string

// SlackActionEvent func
func SlackActionEvent(data *slack.InteractionCallback, context echo.Context) (slack.Message, error) {
	repo := domain.GetSlackRepository()
	var response slack.Message
	var answer string
	option := data.ActionCallback
	switch option.AttachmentActions[0].Name {
	case actionSelect:
		k := strings.Fields(option.AttachmentActions[0].SelectedOptions[0].Value)
		value = option.AttachmentActions[0].SelectedOptions[0].Value
		if utils.StringInSlice("Tag", k) && utils.StringInSlice("EC2", k) {
			originalMessage := data.OriginalMessage
			originalMessage.ReplaceOriginal = true
			originalMessage.Attachments[0].Text = fmt.Sprintf("%s", value)
			originalMessage.Attachments[0].Actions = []slack.AttachmentAction{}
			originalMessage.Attachments[0].Fields = []slack.AttachmentField{
				{
					Title: "Enter Details in Dialogue Box to Process..!!",
					Short: false,
				},
			}
			title := strings.Join(strings.Fields(value)[0:3], " ")

			textInput := slack.NewTextInput("TAGKEY", "Key", "Enter Tag Key")
			textInput1 := slack.NewTextInput("TAGVALUE", "Value", "Enter Tag Value")

			// Open a dialog
			elements := []slack.DialogElement{
				textInput,
				textInput1,
			}
			dialog := slack.Dialog{
				CallbackID:  data.CallbackID,
				State:       value,
				Title:       title,
				SubmitLabel: "Submit",
				Elements:    elements,
			}
			err := repo.OpenDialogMenu(data.TriggerID, dialog)
			if err != nil {
				log.Printf("[ERROR] Unable to send dialog to Slack Channel %s", err)
			}
			response = originalMessage
		} else if utils.StringInSlice("Name", k) && utils.StringInSlice("S3", k) {
			originalMessage := data.OriginalMessage
			originalMessage.ReplaceOriginal = true
			originalMessage.Attachments[0].Text = fmt.Sprintf("%s", value)
			originalMessage.Attachments[0].Actions = []slack.AttachmentAction{}
			originalMessage.Attachments[0].Fields = []slack.AttachmentField{
				{
					Title: "Enter Details in Dialogue Box to Process..!!",
					Short: false,
				},
			}
			title := strings.Join(strings.Fields(value)[0:3], " ")

			textInput := slack.NewTextInput("BUCKETNAME", "Bucket Name", "Enter S3 Bucket Name")

			// Open a dialog
			elements := []slack.DialogElement{
				textInput,
			}
			dialog := slack.Dialog{
				CallbackID:  data.CallbackID,
				State:       value,
				Title:       title,
				SubmitLabel: "Submit",
				Elements:    elements,
			}
			err := repo.OpenDialogMenu(data.TriggerID, dialog)
			if err != nil {
				log.Printf("[ERROR] Unable to send dialog to Slack Channel %s", err)
			}
			response = originalMessage
		} else if utils.StringInSlice("Tag", k) && utils.StringInSlice("SG", k) {
			originalMessage := data.OriginalMessage
			originalMessage.ReplaceOriginal = true
			originalMessage.Attachments[0].Text = fmt.Sprintf("%s", value)
			originalMessage.Attachments[0].Actions = []slack.AttachmentAction{}
			originalMessage.Attachments[0].Fields = []slack.AttachmentField{
				{
					Title: "Enter Details in Dialogue Box to Process..!!",
					Short: false,
				},
			}
			title := strings.Join(strings.Fields(value)[0:3], " ")

			textInput := slack.NewTextInput("TAGKEY", "Key", "Enter Tag Key")
			textInput1 := slack.NewTextInput("TAGVALUE", "Value", "Enter Tag Value")

			// Open a dialog
			elements := []slack.DialogElement{
				textInput,
				textInput1,
			}
			dialog := slack.Dialog{
				CallbackID:  data.CallbackID,
				State:       value,
				Title:       title,
				SubmitLabel: "Submit",
				Elements:    elements,
			}
			err := repo.OpenDialogMenu(data.TriggerID, dialog)
			if err != nil {
				log.Printf("[ERROR] Unable to send dialog to Slack Channel %s", err)
			}
			response = originalMessage
		} else if (utils.StringInSlice("Name", k) && utils.StringInSlice("ELBv1", k)) || (utils.StringInSlice("Name", k) && utils.StringInSlice("ELBv2", k)) {
			originalMessage := data.OriginalMessage
			originalMessage.ReplaceOriginal = true
			originalMessage.Attachments[0].Text = fmt.Sprintf("%s", value)
			originalMessage.Attachments[0].Actions = []slack.AttachmentAction{}
			originalMessage.Attachments[0].Fields = []slack.AttachmentField{
				{
					Title: "Enter Details in Dialogue Box to Process..!!",
					Short: false,
				},
			}
			title := strings.Join(strings.Fields(value)[0:2], " ")

			textInput := slack.NewTextInput("ELBNAME", "ELB Name", "Enter ELB Name")

			// Open a dialog
			elements := []slack.DialogElement{
				textInput,
			}
			dialog := slack.Dialog{
				CallbackID:  data.CallbackID,
				State:       value,
				Title:       title,
				SubmitLabel: "Submit",
				Elements:    elements,
			}
			err := repo.OpenDialogMenu(data.TriggerID, dialog)
			if err != nil {
				log.Printf("[ERROR] Unable to send dialog to Slack Channel %s", err)
			}
			response = originalMessage
		} else if utils.StringInSlice("Name", k) && utils.StringInSlice("DB", k) {
			originalMessage := data.OriginalMessage
			originalMessage.ReplaceOriginal = true
			originalMessage.Attachments[0].Text = fmt.Sprintf("%s", value)
			originalMessage.Attachments[0].Actions = []slack.AttachmentAction{}
			originalMessage.Attachments[0].Fields = []slack.AttachmentField{
				{
					Title: "Enter Details in Dialogue Box to Process..!!",
					Short: false,
				},
			}
			title := strings.Join(strings.Fields(value)[0:3], " ")

			textInput := slack.NewTextInput("RDSDBNAME", "RDS DB Name", "Enter RDS DB Name")

			// Open a dialog
			elements := []slack.DialogElement{
				textInput,
			}
			dialog := slack.Dialog{
				CallbackID:  data.CallbackID,
				State:       value,
				Title:       title,
				SubmitLabel: "Submit",
				Elements:    elements,
			}
			err := repo.OpenDialogMenu(data.TriggerID, dialog)
			if err != nil {
				log.Printf("[ERROR] Unable to send dialog to Slack Channel %s", err)
			}
			response = originalMessage
		} else {
			originalMessage := data.OriginalMessage
			originalMessage.ReplaceOriginal = true
			originalMessage.Attachments[0].Text = fmt.Sprintf("%s\n", value)
			originalMessage.Attachments[0].Text += fmt.Sprintf("Do you want to Perform this action.?")
			originalMessage.Attachments[0].Actions = []slack.AttachmentAction{
				{
					Name:  actionStart,
					Text:  "Yes",
					Type:  "button",
					Value: "start",
					Style: "primary",
				},
				{
					Name:  actionCancel,
					Text:  "No",
					Type:  "button",
					Style: "danger",
				},
			}
			response = originalMessage
		}
	case actionStart:
		t := "Your Request has been accepted..Please wait while processing..!!"
		v := "accepted"
		originalMessage := data.OriginalMessage
		originalMessage.ReplaceOriginal = true
		originalMessage.Attachments[0].Actions = []slack.AttachmentAction{}
		originalMessage.Attachments[0].Fields = []slack.AttachmentField{
			{
				Title: t,
				Short: false,
			},
		}
		answer += v
		response = originalMessage
	case actionCancel:
		t := fmt.Sprintf(":x: @%s %s", data.User.Name, "has cancel the request")
		v := ""
		originalMessage := data.OriginalMessage
		originalMessage.ReplaceOriginal = true
		originalMessage.Attachments[0].Actions = []slack.AttachmentAction{}
		originalMessage.Attachments[0].Fields = []slack.AttachmentField{
			{
				Title: t,
				Value: v,
				Short: false,
			},
		}
		response = originalMessage
	default:
		log.Printf("[ERROR] Invalid action was submitted: %s", option.AttachmentActions[0].Name)
	}
	status := strings.Fields(value)
	if strings.EqualFold(answer, "accepted") && status[2] == "ELBv1" {
		result := awsclient.AwsGetELBv1(config.AwsRegion)
		if len(result) > 0 {
			err := repo.ELBv1ephemeralMessage(config.SlackChanneliD, result)
			if err != nil {
				log.Printf("[ERROR] Unable to send message to Slack Channel %s", err)
			}
		} else {
			noResult := fmt.Sprintf("No v1 ELB Found in %s region in your AWS Account..!!", config.AwsRegion)
			err := repo.BlankResultSlackMsg(config.SlackChanneliD, noResult)
			if err != nil {
				log.Printf("[ERROR] Unable to send message to Slack Channel %s", err)
			}
		}
	} else if strings.EqualFold(answer, "accepted") && status[2] == "ELBv2" {
		result := awsclient.AwsGetELBv2(config.AwsRegion)
		if len(result) > 0 {
			err := repo.ELBv2ephemeralMessage(config.SlackChanneliD, result)
			if err != nil {
				log.Printf("[ERROR] Unable to send message to Slack Channel %s", err)
			}
		} else {
			noResult := fmt.Sprintf("No v2 ELB Found in %s region in your AWS Account..!!", config.AwsRegion)
			err := repo.BlankResultSlackMsg(config.SlackChanneliD, noResult)
			if err != nil {
				log.Printf("[ERROR] Unable to send message to Slack Channel %s", err)
			}
		}
	} else if strings.EqualFold(answer, "accepted") && status[3] == "EC2" {
		result := awsclient.AwsGetInstances(config.AwsRegion, strings.ToLower(status[2]))
		if len(result) > 0 {
			err := repo.EC2ephemeralMessage(config.SlackChanneliD, result)
			if err != nil {
				log.Printf("[ERROR] Unable to send message to Slack Channel %s", err)
			}
		} else {
			noResult := fmt.Sprintf("No EC2 Instances Found in %s region in your AWS Account..!!", config.AwsRegion)
			err := repo.BlankResultSlackMsg(config.SlackChanneliD, noResult)
			if err != nil {
				log.Printf("[ERROR] Unable to send message to Slack Channel %s", err)
			}
		}
	} else if strings.EqualFold(answer, "accepted") && status[3] == "Buckets" {
		result := awsclient.AwsGetS3Buckets(config.AwsRegion)
		if len(result) > 0 {
			err := repo.S3ephemeralMessage(config.SlackChanneliD, result)
			if err != nil {
				log.Printf("[ERROR] Unable to send message to Slack Channel %s", err)
			}
		} else {
			noResult := fmt.Sprintf("No S3 Buckets Found in your AWS Account..!!")
			err := repo.BlankResultSlackMsg(config.SlackChanneliD, noResult)
			if err != nil {
				log.Printf("[ERROR] Unable to send message to Slack Channel %s", err)
			}
		}
	} else if strings.EqualFold(answer, "accepted") && status[3] == "SG" {
		result := awsclient.AwsGetSecGroup(config.AwsRegion)
		if len(result) > 0 {
			err := repo.SGephemeralMessage(config.SlackChanneliD, result)
			if err != nil {
				log.Printf("[ERROR] Unable to send message to Slack Channel %s", err)
			}
		} else {
			noResult := fmt.Sprintf("No SecurityGroups Found in your AWS Account..!!")
			err := repo.BlankResultSlackMsg(config.SlackChanneliD, noResult)
			if err != nil {
				log.Printf("[ERROR] Unable to send message to Slack Channel %s", err)
			}
		}
	} else if strings.EqualFold(answer, "accepted") && status[3] == "DB" {
		result := awsclient.AwsGetRDS(config.AwsRegion)
		if len(result) > 0 {
			err := repo.RDSephemeralMessage(config.SlackChanneliD, result)
			if err != nil {
				log.Printf("[ERROR] Unable to send message to Slack Channel %s", err)
			}
		} else {
			noResult := fmt.Sprintf("No RDS DB Instances Found in %s in your AWS Account..!!", config.AwsRegion)
			err := repo.BlankResultSlackMsg(config.SlackChanneliD, noResult)
			if err != nil {
				log.Printf("[ERROR] Unable to send message to Slack Channel %s", err)
			}
		}
	}
	return response, nil
}

// SlackDialogSubmissionEvent func
func SlackDialogSubmissionEvent(data *slack.InteractionCallback, context echo.Context) error {
	repo := domain.GetSlackRepository()
	choosedOption := strings.Fields(data.State)
	if strings.EqualFold(choosedOption[0], "List") && choosedOption[1] == "EC2" {
		result := awsclient.AwsGetInstancesTag(config.AwsRegion, data.Submission)
		if len(result) > 0 {
			err := repo.EC2ephemeralMessage(config.SlackChanneliD, result)
			if err != nil {
				log.Printf("[ERROR] Unable to send message to Slack Channel %s", err)
			}
		} else {
			noResult := fmt.Sprintf("No Such EC2 Instance in %s region exists in your AWS Account..!! Please check your Input..!!", config.AwsRegion)
			err := repo.BlankResultSlackMsg(config.SlackChanneliD, noResult)
			if err != nil {
				log.Printf("[ERROR] Unable to send message to Slack Channel %s", err)
			}
		}
	} else if strings.EqualFold(choosedOption[0], "List") && choosedOption[1] == "S3" {
		result := awsclient.AwsGetS3BucketsTag(config.AwsRegion, data.Submission)
		if len(result) > 0 {
			err := repo.S3ephemeralMessage(config.SlackChanneliD, result)
			if err != nil {
				log.Printf("[ERROR] Unable to send message to Slack Channel %s", err)
			}
		} else {
			noResult := fmt.Sprintf("No Such S3 Bucket Exists in your AWS Account..!! Please check your Input..!!")
			err := repo.BlankResultSlackMsg(config.SlackChanneliD, noResult)
			if err != nil {
				log.Printf("[ERROR] Unable to send message to Slack Channel %s", err)
			}
		}
	} else if strings.EqualFold(choosedOption[0], "List") && choosedOption[2] == "SG" {
		result := awsclient.AwsGetSecGroupTag(config.AwsRegion, data.Submission)
		if len(result) > 0 {
			err := repo.SGephemeralMessage(config.SlackChanneliD, result)
			if err != nil {
				log.Printf("[ERROR] Unable to send message to Slack Channel %s", err)
			}
		} else {
			noResult := fmt.Sprintf("No Such SecurityGroup Exists in your AWS Account..!! Please check your Input..!!")
			err := repo.BlankResultSlackMsg(config.SlackChanneliD, noResult)
			if err != nil {
				log.Printf("[ERROR] Unable to send message to Slack Channel %s", err)
			}
		}
	} else if strings.EqualFold(choosedOption[0], "List") && choosedOption[1] == "ELBv1" {
		result := awsclient.AwsGetELBv1Tag(config.AwsRegion, data.Submission)
		if len(result) > 0 {
			err := repo.ELBv1ephemeralMessage(config.SlackChanneliD, result)
			if err != nil {
				log.Printf("[ERROR] Unable to send message to Slack Channel %s", err)
			}
		} else {
			noResult := fmt.Sprintf("No Such v1 ELB in %s region exists in your AWS Account..!! Please check your Input..!!", config.AwsRegion)
			err := repo.BlankResultSlackMsg(config.SlackChanneliD, noResult)
			if err != nil {
				log.Printf("[ERROR] Unable to send message to Slack Channel %s", err)
			}
		}
	} else if strings.EqualFold(choosedOption[0], "List") && choosedOption[1] == "ELBv2" {
		result := awsclient.AwsGetELBv2Tag(config.AwsRegion, data.Submission)
		if len(result) > 0 {
			err := repo.ELBv2ephemeralMessage(config.SlackChanneliD, result)
			if err != nil {
				log.Printf("[ERROR] Unable to send message to Slack Channel %s", err)
			}
		} else {
			noResult := fmt.Sprintf("No Such v2 ELB in %s region exists in your AWS Account..!! Please check your Input..!!", config.AwsRegion)
			err := repo.BlankResultSlackMsg(config.SlackChanneliD, noResult)
			if err != nil {
				log.Printf("[ERROR] Unable to send message to Slack Channel %s", err)
			}
		}
	} else if strings.EqualFold(choosedOption[0], "List") && choosedOption[2] == "DB" {
		result := awsclient.AwsGetRDSTag(config.AwsRegion, data.Submission)
		if len(result) > 0 {
			err := repo.RDSephemeralMessage(config.SlackChanneliD, result)
			if err != nil {
				log.Printf("[ERROR] Unable to send message to Slack Channel %s", err)
			}
		} else {
			noResult := fmt.Sprintf("No Such RDS DB in %s exists in your AWS Account..!! Please check your Input..!!", config.AwsRegion)
			err := repo.BlankResultSlackMsg(config.SlackChanneliD, noResult)
			if err != nil {
				log.Printf("[ERROR] Unable to send message to Slack Channel %s", err)
			}
		}
	} else {
		log.Printf("[ERROR] Invalid Action has been choosen..Please Check..!!")
	}
	return nil
}
