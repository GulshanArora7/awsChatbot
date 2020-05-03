package usecase

import (
	"log"

	"github.com/GulshanArora7/awsChatbot/domain"
	"github.com/slack-go/slack"
)

const (
	// action is used for slack attament action.
	actionSelect = "select"
	actionCancel = "cancel"
)

// SendAwsmenuSlackPost func
func SendAwsmenuSlackPost(userid string, channel string, message map[string]string) {
	repo := domain.GetSlackRepository()
	switch message["awsservice"] {
	case "ec2":
		attachment := slack.Attachment{}
		attachment.Color = "#36a64f"
		attachment.Text = "AWS EC2 Actions"
		attachment.CallbackID = "AWS EC2"
		attachment.Actions = []slack.AttachmentAction{
			{
				Name: actionSelect,
				Type: "select",
				Options: []slack.AttachmentActionOption{
					{
						Text:  "List All Running EC2",
						Value: "List All Running EC2",
					},
					{
						Text:  "List All Stopped EC2",
						Value: "List All Stopped EC2",
					},
					{
						Text:  "List EC2 Instance By Tag",
						Value: "List EC2 Instance By Tag",
					},
				},
			},
			{
				Name:  actionCancel,
				Text:  "Cancel",
				Type:  "button",
				Style: "danger",
			},
		}
		err := repo.EphemeralMenuMessage(channel, attachment)
		if err != nil {
			log.Printf("[ERROR] Unable to send message to Slack Channel %s", err)
		}
	case "s3":
		attachment := slack.Attachment{}
		attachment.Color = "#36a64f"
		attachment.Text = "AWS S3 Actions"
		attachment.CallbackID = "AWS S3"
		attachment.Actions = []slack.AttachmentAction{
			{
				Name: actionSelect,
				Type: "select",
				Options: []slack.AttachmentActionOption{
					{
						Text:  "List All S3 Buckets",
						Value: "List All S3 Buckets",
					},
					{
						Text:  "List S3 Bucket By Name",
						Value: "List S3 Bucket By Name",
					},
				},
			},
			{
				Name:  actionCancel,
				Text:  "Cancel",
				Type:  "button",
				Style: "danger",
			},
		}
		err := repo.EphemeralMenuMessage(channel, attachment)
		if err != nil {
			log.Printf("[ERROR] Unable to send message to Slack Channel %s", err)
		}

	case "sg", "securitygroup":
		attachment := slack.Attachment{}
		attachment.Color = "#36a64f"
		attachment.Text = "AWS SG Actions"
		attachment.CallbackID = "AWS SecurityGroup"
		attachment.Actions = []slack.AttachmentAction{
			{
				Name: actionSelect,
				Type: "select",
				Options: []slack.AttachmentActionOption{
					{
						Text:  "List All Open SG",
						Value: "List All Open SG",
					},
					{
						Text:  "List Open SG By Tag",
						Value: "List Open SG By Tag",
					},
				},
			},
			{
				Name:  actionCancel,
				Text:  "Cancel",
				Type:  "button",
				Style: "danger",
			},
		}
		err := repo.EphemeralMenuMessage(channel, attachment)
		if err != nil {
			log.Printf("[ERROR] Unable to send message to Slack Channel %s", err)
		}

	case "elbv1":
		attachment := slack.Attachment{}
		attachment.Color = "#36a64f"
		attachment.Text = "AWS ELBv1 Actions"
		attachment.CallbackID = "AWS ELBv1"
		attachment.Actions = []slack.AttachmentAction{
			{
				Name: actionSelect,
				Type: "select",
				Options: []slack.AttachmentActionOption{
					{
						Text:  "List All ELBv1",
						Value: "List All ELBv1",
					},
					{
						Text:  "List ELBv1 By Name",
						Value: "List ELBv1 By Name",
					},
				},
			},
			{
				Name:  actionCancel,
				Text:  "Cancel",
				Type:  "button",
				Style: "danger",
			},
		}
		err := repo.EphemeralMenuMessage(channel, attachment)
		if err != nil {
			log.Printf("[ERROR] Unable to send message to Slack Channel %s", err)
		}
	case "elbv2":
		attachment := slack.Attachment{}
		attachment.Color = "#36a64f"
		attachment.Text = "AWS ELBv2 Actions"
		attachment.CallbackID = "AWS ELBv2"
		attachment.Actions = []slack.AttachmentAction{
			{
				Name: actionSelect,
				Type: "select",
				Options: []slack.AttachmentActionOption{
					{
						Text:  "List All ELBv2",
						Value: "List All ELBv2",
					},
					{
						Text:  "List ELBv2 By Name",
						Value: "List ELBv2 By Name",
					},
				},
			},
			{
				Name:  actionCancel,
				Text:  "Cancel",
				Type:  "button",
				Style: "danger",
			},
		}
		err := repo.EphemeralMenuMessage(channel, attachment)
		if err != nil {
			log.Printf("[ERROR] Unable to send message to Slack Channel %s", err)
		}
	case "rds":
		attachment := slack.Attachment{}
		attachment.Color = "#36a64f"
		attachment.Text = "AWS RDS Actions"
		attachment.CallbackID = "AWS RDS"
		attachment.Actions = []slack.AttachmentAction{
			{
				Name: actionSelect,
				Type: "select",
				Options: []slack.AttachmentActionOption{
					{
						Text:  "List All RDS DB",
						Value: "List All RDS DB",
					},
					{
						Text:  "List RDS DB By Name",
						Value: "List RDS DB By Name",
					},
				},
			},
			{
				Name:  actionCancel,
				Text:  "Cancel",
				Type:  "button",
				Style: "danger",
			},
		}
		err := repo.EphemeralMenuMessage(channel, attachment)
		if err != nil {
			log.Printf("[ERROR] Unable to send message to Slack Channel %s", err)
		}
	case "help":
		attachment := slack.Attachment{}
		attachment.Color = "#36a64f"
		attachment.Pretext = "Chatbot Help"
		attachment.Title = "Supported AWS Service - ec2, s3, sg or securitygroup, elbv1, elbv2, rds"
		attachment.Text = "Usage Example: /awsChatbot ec2"
		attachment.CallbackID = "Invalid"
		err := repo.EphemeralMenuMessage(channel, attachment)
		if err != nil {
			log.Printf("[ERROR] Unable to send message to Slack Channel %s", err)
		}
	default:
		attachment := slack.Attachment{}
		attachment.Color = "#36a64f"
		attachment.Pretext = "Invalid Option Choosen"
		attachment.Title = "Get Help from awsChatbot help Command"
		attachment.Text = "example: /awsChatbot help"
		attachment.CallbackID = "Invalid"
		err := repo.EphemeralMenuMessage(channel, attachment)
		if err != nil {
			log.Printf("[ERROR] Unable to send message to Slack Channel %s", err)
		}
	}
}
