package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/nlopes/slack/slackevents"

	"github.com/TonyCioara/feedback-bot/utils"
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
	go SetUpEventsAPI(apiKey)
	rtm := api.NewRTM()
	go rtm.ManageConnection() // goroutine!
	return rtm
}

func SetUpEventsAPI(apiKey string) {
	fmt.Println("1) Seting up")
	http.HandleFunc("/events-endpoint", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("2) Receiving")
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		body := buf.String()

		eventsAPIEvent, e := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionVerifyToken(&slackevents.TokenComparator{VerificationToken: apiKey}))
		if e != nil {
			fmt.Println("Something went wrong: ", e, "\n", eventsAPIEvent)
			w.WriteHeader(http.StatusInternalServerError)
		}
		if eventsAPIEvent.Type == slackevents.CallbackEvent {
			innerEvent := eventsAPIEvent.InnerEvent
			fmt.Println("Callback:", innerEvent)
		}
	})
	fmt.Println("[INFO] Server listening")
	http.ListenAndServe(":3000", nil)
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
	attachment := utils.GenerateHelpButtons()
	response := "What can I help you with?"
	slackClient.PostMessage(slackChannel, slack.MsgOptionText(response, false), slack.MsgOptionAttachments(attachment))
}

func sendResponse(slackClient *slack.RTM, message, slackChannel string) {
	command := strings.ToLower(message)
	println("[RECEIVED] sendResponse:", command)

	slackClient.SendMessage(slackClient.NewOutgoingMessage("I Like pizza", slackChannel))
}

func greet(slackClient *slack.RTM, slackChannel string) {

	attachment := utils.GenerateHelpButtons()
	response := "What can I help you with?"
	slackClient.PostMessage(slackChannel, slack.MsgOptionText(response, false), slack.MsgOptionAttachments(attachment))
}

func sendSurvey() {

}
