package slack

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/TonyCioara/feedback-bot/controllers"
	"github.com/nlopes/slack"
	"github.com/nlopes/slack/slackevents"
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
	go SetUpEventsAPI(api)
	rtm := api.NewRTM()
	go rtm.ManageConnection() // goroutine!
	return rtm
}

type EType struct {
	EType string `json:"type"`
}

func SetUpEventsAPI(api *slack.Client) {
	fmt.Println("1) Seting up")
	http.HandleFunc("/events-endpoint", func(w http.ResponseWriter, r *http.Request) {
		token := os.Getenv("VERIFICATION_TOKEN")

		data, _ := ioutil.ReadAll(r.Body)
		body := string(data)
		body, _ = url.QueryUnescape(body)
		body = strings.Replace(body, "payload=", "", 1)

		// Update to check type, then parse into custom model for dialogs and actionEvent for actionEvent
		actionEvent, e := slackevents.ParseActionEvent(body, slackevents.OptionVerifyToken(&slackevents.TokenComparator{VerificationToken: token}))
		if e != nil {
			fmt.Println("Something went wrong: ", e)
			w.WriteHeader(http.StatusInternalServerError)
		}
		fmt.Println("actionEventHappened:", actionEvent.Type)

		switch et := actionEvent.Type; et {
		case "interactive_message":
			controllers.ButtonClicked(api, actionEvent)
		case "dialog_submission":
			controllers.DialogReceived(api, body)
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

			controllers.SendHelp(slackClient, message, ev.Channel)
		case *slack.IMCreatedEvent:
			controllers.Greet(slackClient, ev.Channel.ID)
		}
	}
}
