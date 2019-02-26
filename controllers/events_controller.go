package controllers

import (
	"fmt"

	"github.com/TonyCioara/feedback-bot/utils"

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

	dialog := utils.GenerateFeedbackSurvey(action.TriggerID, action.CallbackID)

	err := api.OpenDialog(action.TriggerID, dialog)

	fmt.Println("32. Error:", err)
}
