package slack

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/TonyCioara/feedback-bot/controllers"
	"github.com/nlopes/slack"
	"github.com/nlopes/slack/slackevents"
)

const helpMessage = "type in '@feedback-bot'"

/*
CreateSlackClient sets up the slack RTM (real-timemessaging) client library,
initiating the socket connection and returning the client.
*/

func CreateSlackClient(apiKey string) *slack.RTM {
	api := slack.New(apiKey)
	go SetUpEventsAPI(api)
	rtm := api.NewRTM()
	go rtm.ManageConnection() // goroutine!
	return rtm
}

// SetUpEventsAPI sets up the events api
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
			log.Fatalf("Something went wrong: %s", e)
			w.WriteHeader(http.StatusInternalServerError)
		}

		switch et := actionEvent.Type; et {
		case "interactive_message":
			go controllers.ButtonClicked(api, actionEvent)
		case "dialog_submission":
			go controllers.DialogReceived(api, body)
		}

	})

	port := ":" + os.Getenv("PORT")
	if port == ":" {
		port = ":3000"
	}

	fmt.Println("[INFO] Server listening on port", port)
	http.ListenAndServe(port, nil)
}

/*
RespondToEvents waits for messages on the Slack client's incomingEvents channel,
and sends a response when it detects the bot has been tagged in a message with @<botTag>.
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

			go controllers.MessageReceived(slackClient, message, ev)
		case *slack.IMCreatedEvent:
			go controllers.Greet(slackClient, ev.Channel.ID)
		}
	}
}
