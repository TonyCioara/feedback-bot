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

	var dialogElement1 slack.DialogElement = map[string]string{
		"type":  "text",
		"label": "Pickup Location",
		"name":  "loc_origin",
	}
	var dialogElement2 slack.DialogElement = map[string]string{
		"type":  "text",
		"label": "Pickup Location",
		"name":  "loc_origi2n",
	}

	dialogElements := []slack.DialogElement{dialogElement1, dialogElement2}

	dialog := slack.Dialog{
		TriggerID:      action.TriggerID,
		CallbackID:     action.CallbackID,
		Title:          "Submit Feedback",
		SubmitLabel:    "SubmitLabel",
		NotifyOnCancel: true,
		Elements:       dialogElements,
	}

	err := api.OpenDialog(action.TriggerID, dialog)

	fmt.Println("32. Error:", err)
}
