package controllers

import (
	"fmt"
	"os"
	"strings"

	"github.com/TonyCioara/feedback-bot/models"
	"github.com/TonyCioara/feedback-bot/utils"
	"github.com/nlopes/slack"
)

func MessageReceived(slackClient *slack.RTM, message string, slackEvent *slack.MessageEvent) {
	elements := strings.Split(message, " ")
	switch first := elements[0]; first {
	case "help":
		SendHelp(slackClient, slackEvent.Channel)
	case "find":
		FindFeedback(slackClient, slackEvent, elements)
	case "delete":
		DeleteFeedback(slackClient, slackEvent, elements)
	case "unsubscribe":
		Unsubscribe(slackClient, slackEvent)
	case "subscribe":
		Subscribe(slackClient, slackEvent)
	}

}

// SendHelp gets called when the user types in 'help'
func SendHelp(slackClient *slack.RTM, slackChannel string) {

	attachment := utils.GenerateHelpButtons()
	response := "*What can I help you with?*"
	fmt.Println("Channel:", slackChannel)
	slackClient.PostMessage(slackChannel, slack.MsgOptionText(response, false), slack.MsgOptionAttachments(attachment))
}

// Greet gets called when the user creates a new chat with the bot
func Greet(slackClient *slack.RTM, slackChannel string) {

	attachment := utils.GenerateHelpButtons()
	response := "*What can I help you with?*"
	slackClient.PostMessage(slackChannel, slack.MsgOptionText(response, false), slack.MsgOptionAttachments(attachment))
}

// FindFeedback sends a Feedback CSV with given queries
func FindFeedback(slackClient *slack.RTM, slackEvent *slack.MessageEvent, elements []string) {
	api := slack.New(os.Getenv("BOT_OAUTH_ACCESS_TOKEN"))
	SendFeedbackCSV(api, slackEvent.User, slackEvent.Username, elements)
}

// Subscribe subscribes a user to the weekly feedback
func Subscribe(slackClient *slack.RTM, slackEvent *slack.MessageEvent) {
	db := utils.StartAndMigrateDB()
	var user models.User
	db.Model(&user).Where("user_id = ?", slackEvent.User).Update("active_subscription", true)
	slackClient.PostMessage(slackEvent.Channel, slack.MsgOptionText("You have been subscribed to the weekly feedback!", false))

}

// Unsubscribe unsubscribes a user from the weekly feedback
func Unsubscribe(slackClient *slack.RTM, slackEvent *slack.MessageEvent) {
	db := utils.StartAndMigrateDB()
	var user models.User
	db.Model(&user).Where("user_id = ?", slackEvent.User).Update("active_subscription", false)
	slackClient.PostMessage(slackEvent.Channel, slack.MsgOptionText("You have been unsubscribed from the weekly feedback!", false))
}
