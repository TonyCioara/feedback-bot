package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/TonyCioara/feedback-bot/models"

	"github.com/nlopes/slack/slackevents"
)

// SetUpEventsAPI sets up the events api
func SetUpEventsAPI() {
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
			go ButtonClicked(actionEvent)
		case "dialog_submission":
			go DialogReceived(body)
		}
	})

	port := ":" + os.Getenv("PORT")
	if port == ":" {
		port = ":3000"
	}

	fmt.Println("[INFO] Server listening on port", port)
	http.ListenAndServe(port, nil)
}

// ButtonClicked is responsible for handling button clicking
func ButtonClicked(action slackevents.MessageAction) {
	switch value := action.Actions[0].Value; value {
	case "sendFeedback":
		SendFeedbackSurvey(action)
	case "seeFeedback":
		SendFeedbackCSV(action.User.ID, action.User.Name, []string{})
	case "moreHelp":
		SendMoreHelp(action)
	}
}

// DialogReceived is responsible for handling received dialogs
func DialogReceived(payloadString string) {
	byteString := []byte(payloadString)
	dialog := models.DialogSubmission{}
	err := json.Unmarshal(byteString, &dialog)
	if err != nil {
		log.Fatalf("Dialog parsing error: %s", err)
		return
	}

	CreateFeedbackFromDialog(dialog)
}
