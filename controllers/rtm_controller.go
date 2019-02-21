package controllers

import (
	"strings"

	"github.com/TonyCioara/feedback-bot/utils"
	"github.com/nlopes/slack"
)

// SendHelp gets called when the user types in 'help'
func SendHelp(slackClient *slack.RTM, message, slackChannel string) {
	if strings.ToLower(message) != "help" {
		return
	}
	attachment := utils.GenerateHelpButtons()
	response := "What can I help you with?"
	slackClient.PostMessage(slackChannel, slack.MsgOptionText(response, false), slack.MsgOptionAttachments(attachment))
}

// Greet gets called when the user creates a new chat with the bot
func Greet(slackClient *slack.RTM, slackChannel string) {

	attachment := utils.GenerateHelpButtons()
	response := "What can I help you with?"
	slackClient.PostMessage(slackChannel, slack.MsgOptionText(response, false), slack.MsgOptionAttachments(attachment))
}
