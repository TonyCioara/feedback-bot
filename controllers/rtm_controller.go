package controllers

import (
	"strings"

	"github.com/TonyCioara/feedback-bot/models"
	"github.com/TonyCioara/feedback-bot/server"
	"github.com/TonyCioara/feedback-bot/utils"
	"github.com/nlopes/slack"
	"github.com/nlopes/slack/slackevents"
)

/*
RespondToEvents waits for messages on the Slack client's incomingEvents channel,
and sends a response when it detects the bot has been tagged in a message with @<botTag>.
*/
func RespondToEvents() {
	slackClient := server.RTM
	for msg := range slackClient.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			message := ev.Msg.Text
			go MessageReceived(message, ev)
		case *slack.IMCreatedEvent:
			go Greet(ev.Channel.ID)
		}
	}
}

// MessageReceived handles incoming messages
func MessageReceived(message string, slackEvent *slack.MessageEvent) {
	elements := strings.Split(message, " ")
	switch first := elements[0]; first {
	case "help":
		SendHelp(slackEvent.Channel)
	case "find":
		FindFeedback(slackEvent, elements)
	case "delete":
		DeleteFeedback(slackEvent, elements)
	case "unsubscribe":
		Unsubscribe(slackEvent)
	case "subscribe":
		Subscribe(slackEvent)
	}
}

// SendHelp gets called when the user types in 'help'
func SendHelp(slackChannel string) {
	slackClient := server.RTM

	attachment := utils.GenerateHelpButtons()
	response := "*What can I help you with?*"
	slackClient.PostMessage(slackChannel, slack.MsgOptionText(response, false), slack.MsgOptionAttachments(attachment))
}

// Greet gets called when the user creates a new chat with the bot
func Greet(slackChannel string) {
	slackClient := server.RTM

	attachment := utils.GenerateHelpButtons()
	response := "*What can I help you with?*"
	slackClient.PostMessage(slackChannel, slack.MsgOptionText(response, false), slack.MsgOptionAttachments(attachment))
}

// FindFeedback sends a Feedback CSV with given queries
func FindFeedback(slackEvent *slack.MessageEvent, elements []string) {
	SendFeedbackCSV(slackEvent.User, slackEvent.Username, elements)
}

// Subscribe subscribes a user to the weekly feedback
func Subscribe(slackEvent *slack.MessageEvent) {
	db := server.DB
	slackClient := server.RTM

	var user models.User
	db.Model(&user).Where("user_id = ?", slackEvent.User).Update("active_subscription", true)

	slackClient.PostMessage(slackEvent.Channel, slack.MsgOptionText("You have been subscribed to the weekly feedback!", false))

}

// Unsubscribe unsubscribes a user from the weekly feedback
func Unsubscribe(slackEvent *slack.MessageEvent) {
	db := server.DB
	slackClient := server.RTM

	var user models.User
	db.Model(&user).Where("user_id = ?", slackEvent.User).Update("active_subscription", false)

	slackClient.PostMessage(slackEvent.Channel, slack.MsgOptionText("You have been unsubscribed from the weekly feedback!", false))
}

// SendMoreHelp sends the user information about the commands
func SendMoreHelp(action slackevents.MessageAction) {
	slackClient := server.RTM

	response :=
		"*Here are a few useful commands:*\nTo query feedback use: \n   - `find param_name=param` \n   - Example: `find Type=intensives Sender=steve` \nTo delete feedback use: \n   - `delete ID` \n   - Example: `delete 28` \nTo subscribe to weekly feedback use: \n   - `subscribe` \nTo unsubscribe from weekly feedback use: \n   - `unsubscribe` "

	slackClient.PostMessage(action.Channel.ID, slack.MsgOptionText(response, false))
}
