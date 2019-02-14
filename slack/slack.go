package slack

import (
	"fmt"
	"strings"

	"github.com/nlopes/slack"
)

/*
   TODO: Change @BOT_NAME to the same thing you entered when creating your Slack application.
   NOTE: command_arg_1 and command_arg_2 represent optional parameteras that you define
   in the Slack API UI
*/
const helpMessage = "type in '@feedback-bot'"

/*
   CreateSlackClient sets up the slack RTM (real-timemessaging) client library,
   initiating the socket connection and returning the client.
   DO NOT EDIT THIS FUNCTION. This is a fully complete implementation.
*/
func CreateSlackClient(apiKey string) *slack.RTM {
	api := slack.New(apiKey)
	rtm := api.NewRTM()
	go rtm.ManageConnection() // goroutine!
	return rtm
}

/*
   RespondToEvents waits for messages on the Slack client's incomingEvents channel,
   and sends a response when it detects the bot has been tagged in a message with @<botTag>.

   EDIT THIS FUNCTION IN THE SPACE INDICATED ONLY!
*/
func RespondToEvents(slackClient *slack.RTM) {
	for msg := range slackClient.IncomingEvents {
		fmt.Println("Event Received: ", msg.Type)
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			// botTagString := fmt.Sprintf("<@%s> ", slackClient.GetInfo().User.ID)
			// if !strings.Contains(ev.Msg.Text, botTagString) {
			// 	continue
			// }
			// message := strings.Replace(ev.Msg.Text, botTagString, "", -1)
			message := ev.Msg.Text

			sendHelp(slackClient, message, ev.Channel)
		case *slack.IMCreatedEvent:
			greet(slackClient, ev.Channel.ID)
		}
	}
}

// sendHelp is a working help message, for reference.
func sendHelp(slackClient *slack.RTM, message, slackChannel string) {
	if strings.ToLower(message) != "help" {
		return
	}
	// attachment := utils.generateHelpButtons()
	response := "What can I help you with?"
	slackClient.PostMessage(slackChannel, slack.MsgOptionText(response, false), slack.MsgOptionAttachments(attachment))
}

// sendResponse is NOT unimplemented --- write code in the function body to complete!

func sendResponse(slackClient *slack.RTM, message, slackChannel string) {
	command := strings.ToLower(message)
	println("[RECEIVED] sendResponse:", command)

	slackClient.SendMessage(slackClient.NewOutgoingMessage("I Like pizza", slackChannel))
}

func greet(slackClient *slack.RTM, slackChannel string) {

	// attachment := utils.generateHelpButtons()
	response := "What can I help you with?"
	slackClient.PostMessage(slackChannel, slack.MsgOptionText(response, false), slack.MsgOptionAttachments(attachment))
}
