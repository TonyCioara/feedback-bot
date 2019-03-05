package controllers

import (
	"fmt"
	"os"
	"strings"

	"github.com/TonyCioara/feedback-bot/utils"
	"github.com/nlopes/slack"
)

func MessageReceived(slackClient *slack.RTM, message string, slackEvent *slack.MessageEvent) {
	elements := strings.Split(message, " ")
	switch first := elements[0]; first {
	case "help":
		go SendHelp(slackClient, slackEvent.Channel)
	case "find":
		go FindFeedback(slackClient, slackEvent, elements)
	case "delete":
		go DeleteFeedback(slackClient, slackEvent, elements)
	}

}

// SendHelp gets called when the user types in 'help'
func SendHelp(slackClient *slack.RTM, slackChannel string) {

	attachment := utils.GenerateHelpButtons()
	response := "What can I help you with?"
	fmt.Println("Channel:", slackChannel)
	slackClient.PostMessage(slackChannel, slack.MsgOptionText(response, false), slack.MsgOptionAttachments(attachment))
}

// Greet gets called when the user creates a new chat with the bot
func Greet(slackClient *slack.RTM, slackChannel string) {

	attachment := utils.GenerateHelpButtons()
	response := "What can I help you with?"
	slackClient.PostMessage(slackChannel, slack.MsgOptionText(response, false), slack.MsgOptionAttachments(attachment))
}

func FindFeedback(slackClient *slack.RTM, slackEvent *slack.MessageEvent, elements []string) {

	api := slack.New(os.Getenv("BOT_OAUTH_ACCESS_TOKEN"))
	SendFeedbackCSV(api, slackEvent.User, slackEvent.Username, elements)
}
