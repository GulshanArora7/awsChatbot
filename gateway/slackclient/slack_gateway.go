package slackclient

import (
	"fmt"
	"log"

	"github.com/GulshanArora7/awsChatbot/appcontext"
	"github.com/GulshanArora7/awsChatbot/config"
	"github.com/GulshanArora7/awsChatbot/domain"
	"github.com/slack-go/slack"
)

// Slack struct for our slackbot
type Slack struct {
	Name  string
	Token string

	User   string
	UserID string

	Client *slack.Client
}

// EC2ephemeralMessage func
func (repo Slack) EC2ephemeralMessage(channel string, message []domain.EC2Dictionary) error {
	for _, v := range message {
		// Header Section
		headerText := slack.NewTextBlockObject("mrkdwn", "*AWS EC2 DETAILS*", false, false)
		headerSection := slack.NewSectionBlock(headerText, nil, nil)

		// Fields
		instanceIDField := slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("*Instance ID:*\n%s", v["InstanceId"]), false, false)
		instancetypeField := slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("*Instance Type:*\n%s", v["InstanceType"]), false, false)
		statusField := slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("*Status:*\n%s", v["State"]), false, false)
		vpcField := slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("*Vpc ID:*\n%s", v["VpcID"]), false, false)
		subnetField := slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("*Subnet ID:*\n%s", v["SubnetID"]), false, false)
		privateipField := slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("*PrivateIP:*\n%s", v["PrivateIP"]), false, false)
		publicipField := slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("*PublicIP:*\n%s", v["PublicIP"]), false, false)
		amiField := slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("*AMI-ID:*\n%s", v["ImageID"]), false, false)

		fieldSlice := make([]*slack.TextBlockObject, 0)
		fieldSlice = append(fieldSlice, instanceIDField)
		fieldSlice = append(fieldSlice, instancetypeField)
		fieldSlice = append(fieldSlice, statusField)
		fieldSlice = append(fieldSlice, vpcField)
		fieldSlice = append(fieldSlice, subnetField)
		fieldSlice = append(fieldSlice, privateipField)
		fieldSlice = append(fieldSlice, publicipField)
		fieldSlice = append(fieldSlice, amiField)

		fieldsSection := slack.NewSectionBlock(nil, fieldSlice, nil)

		_, _, err := repo.Client.PostMessage(channel, slack.MsgOptionBlocks(headerSection, fieldsSection))
		if err != nil {
			log.Printf("[ERROR] Unable to send message to Slack Channel %s", err)
		}
	}
	return nil
}

// S3ephemeralMessage Func
func (repo Slack) S3ephemeralMessage(channel string, message []domain.S3Dictionary) error {
	for _, v := range message {
		// Header Section
		headerText := slack.NewTextBlockObject("mrkdwn", "*AWS S3 DETAILS*", false, false)
		headerSection := slack.NewSectionBlock(headerText, nil, nil)

		// Fields
		bucketnameField := slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("*Bucket Name:*\n%s", v["BucketName"]), false, false)
		bucketcreationField := slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("*Bucket Creation Date:*\n%s", v["CreationDate"]), false, false)

		fieldSlice := make([]*slack.TextBlockObject, 0)
		fieldSlice = append(fieldSlice, bucketnameField)
		fieldSlice = append(fieldSlice, bucketcreationField)

		fieldsSection := slack.NewSectionBlock(nil, fieldSlice, nil)

		_, _, err := repo.Client.PostMessage(channel, slack.MsgOptionBlocks(headerSection, fieldsSection))
		if err != nil {
			log.Printf("[ERROR] Unable to send message to Slack Channel %s", err)
		}
	}
	return nil
}

// SGephemeralMessage Func
func (repo Slack) SGephemeralMessage(channel string, message []domain.SGDictionary) error {
	for _, v := range message {
		// Header Section
		headerText := slack.NewTextBlockObject("mrkdwn", "*AWS SECURITY GROUP DETAILS*", false, false)
		headerSection := slack.NewSectionBlock(headerText, nil, nil)

		// Fields
		sgidField := slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("*SecurityGroup ID:*\n%s", v["SecurityGroupID"]), false, false)
		sgnameField := slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("*SecurityGroup Name:*\n%s", v["SecurityGroupName"]), false, false)
		fromportField := slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("*From Port:* %d", v["FromPort"]), false, false)
		toportField := slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("*To Port:* %d", v["ToPort"]), false, false)
		ingcidrField := slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("*Ingress CIDR:* %s", v["IngressCIDR"]), false, false)

		fieldSlice := make([]*slack.TextBlockObject, 0)
		fieldSlice = append(fieldSlice, sgidField)
		fieldSlice = append(fieldSlice, sgnameField)
		fieldSlice = append(fieldSlice, fromportField)
		fieldSlice = append(fieldSlice, toportField)
		fieldSlice = append(fieldSlice, ingcidrField)

		fieldsSection := slack.NewSectionBlock(nil, fieldSlice, nil)

		_, _, err := repo.Client.PostMessage(channel, slack.MsgOptionBlocks(headerSection, fieldsSection))
		if err != nil {
			log.Printf("[ERROR] Unable to send message to Slack Channel %s", err)
		}
	}
	return nil
}

// ELBv1ephemeralMessage Func
func (repo Slack) ELBv1ephemeralMessage(channel string, message []domain.ELBv1Dictionary) error {
	for _, v := range message {
		// Header Section
		headerText := slack.NewTextBlockObject("mrkdwn", "*AWS ELBV1 DETAILS*", false, false)
		headerSection := slack.NewSectionBlock(headerText, nil, nil)

		// Fields
		elbv1nameField := slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("*ELB Name:*\n%s", v["Elbv1Name"]), false, false)
		elbv1dnsnameField := slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("*ELB DNS Name:*\n%s", v["Elbv1DNSName"]), false, false)
		elbv1schemeField := slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("*ELB Scheme:*\n%s", v["Elbv1Scheme"]), false, false)
		elbv1dateField := slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("*ELB Creation Date:*\n%s", v["Elbv1CreationDate"]), false, false)

		fieldSlice := make([]*slack.TextBlockObject, 0)
		fieldSlice = append(fieldSlice, elbv1nameField)
		fieldSlice = append(fieldSlice, elbv1dnsnameField)
		fieldSlice = append(fieldSlice, elbv1schemeField)
		fieldSlice = append(fieldSlice, elbv1dateField)

		fieldsSection := slack.NewSectionBlock(nil, fieldSlice, nil)

		_, _, err := repo.Client.PostMessage(channel, slack.MsgOptionBlocks(headerSection, fieldsSection))
		if err != nil {
			log.Printf("[ERROR] Unable to send message to Slack Channel %s", err)
		}
	}
	return nil
}

// ELBv2ephemeralMessage Func
func (repo Slack) ELBv2ephemeralMessage(channel string, message []domain.ELBv2Dictionary) error {
	for _, v := range message {
		// Header Section
		headerText := slack.NewTextBlockObject("mrkdwn", "*AWS ELBV2 DETAILS*", false, false)
		headerSection := slack.NewSectionBlock(headerText, nil, nil)

		// Fields
		elbv2nameField := slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("*ELB Name:*\n%s", v["Elbv2Name"]), false, false)
		elbv2dnsnameField := slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("*ELB DNS Name:*\n%s", v["Elbv2DNSName"]), false, false)
		elbv2schemeField := slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("*ELB Scheme:*\n%s", v["Elbv2Scheme"]), false, false)
		elbv2statusField := slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("*ELB Status:*\n%s", v["ELBv2Status"]), false, false)
		elbv2dateField := slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("*ELB Creation Date:*\n%s", v["Elbv2CreationDate"]), false, false)

		fieldSlice := make([]*slack.TextBlockObject, 0)
		fieldSlice = append(fieldSlice, elbv2nameField)
		fieldSlice = append(fieldSlice, elbv2dnsnameField)
		fieldSlice = append(fieldSlice, elbv2schemeField)
		fieldSlice = append(fieldSlice, elbv2statusField)
		fieldSlice = append(fieldSlice, elbv2dateField)

		fieldsSection := slack.NewSectionBlock(nil, fieldSlice, nil)

		_, _, err := repo.Client.PostMessage(channel, slack.MsgOptionBlocks(headerSection, fieldsSection))
		if err != nil {
			log.Printf("[ERROR] Unable to send message to Slack Channel %s", err)
		}
	}
	return nil
}

// EphemeralMenuMessage func
func (repo Slack) EphemeralMenuMessage(channel string, attachment slack.Attachment) error {
	_, _, err := repo.Client.PostMessage(channel, slack.MsgOptionAttachments(attachment))
	if err != nil {
		log.Printf("[ERROR] Unable to send message to Slack Channel %s", err)
	}
	return nil
}

// DialogueAckMessageFunc func
func (repo Slack) DialogueAckMessageFunc(channel string, text string) error {
	_, _, err := repo.Client.PostMessage(channel, slack.MsgOptionText(text, false))
	if err != nil {
		log.Printf("[ERROR] Unable to send message to Slack Channel %s", err)
	}
	return nil
}

// OpenDialogMenu func
func (repo Slack) OpenDialogMenu(triggerID string, dialogue slack.Dialog) error {
	err := repo.Client.OpenDialog(triggerID, dialogue)
	if err != nil {
		log.Printf("[ERROR] Unable to send message to Slack Channel %s", err)
	}
	return nil
}

// BlankResultSlackMsg func
func (repo Slack) BlankResultSlackMsg(channel string, message string) error {
	attachment := slack.Attachment{}
	attachment.Color = "#36a64f"
	attachment.Text = message
	attachment.Footer = "AWS"
	attachment.FooterIcon = "https://a0.awsstatic.com/libra-css/images/logos/aws_logo_smile_1200x630.png"
	_, _, err := repo.Client.PostMessage(channel, slack.MsgOptionAttachments(attachment))
	if err != nil {
		log.Printf("[ERROR] Unable to send message to Slack Channel %s", err)
	}
	return nil
}

// New func to initializate Slack Client
func New() (*Slack, error) {
	return &Slack{Client: slack.New(config.SlackBotToken), Token: config.SlackBotToken, Name: "Slack Client"}, nil
}

func init() {
	if config.GetEnv("TESTRUN", "false") == "true" {
		return
	}
	SlackRepository, err := New()
	if err != nil {
		log.Println("[ERROR] Slack Repository not initiated")
	}
	appcontext.Current.Add(appcontext.SlackRepository, SlackRepository)
	if appcontext.Current.Count() != 0 {
		log.Println("[INFO] Slack Repository initiated")
	}
}
