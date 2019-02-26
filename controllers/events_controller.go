package controllers

import (
	"fmt"

	"github.com/nlopes/slack"
	"github.com/nlopes/slack/slackevents"
)

// ButtonClicked is responsible for handling button clicking
func ButtonClicked(api *slack.Client, action slackevents.MessageAction) {
	switch value := action.Actions[0].Value; value {
	case "sendFeedback":
		SendFeedbackSurvey(api, action)
	}
}

// SendFeedbackSurvey sends the user a feedback survey
func SendFeedbackSurvey(api *slack.Client, action slackevents.MessageAction) {
	fmt.Println("------GOT HERE-----")
	dialog := slack.Dialog{
		TriggerID:      action.TriggerID,
		CallbackID:     action.CallbackID,
		Title:          "Submit Feedback",
		SubmitLabel:    "SubmitLabel",
		NotifyOnCancel: true,
		Elements:       []slack.DialogElement{},
	}
e
	err := api.OpenDialog(action.TriggerID, dialog)

	fmt.Println("123:", err)
}
