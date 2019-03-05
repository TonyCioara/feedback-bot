package controllers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/TonyCioara/feedback-bot/models"
	"github.com/TonyCioara/feedback-bot/utils"

	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/nlopes/slack"
	"github.com/nlopes/slack/slackevents"
)

// ButtonClicked is responsible for handling button clicking
func ButtonClicked(api *slack.Client, action slackevents.MessageAction) {
	switch value := action.Actions[0].Value; value {
	case "sendFeedback":
		SendFeedbackSurvey(api, action)
	case "seeFeedback":
		SendFeedbackCSV(api, action.User.ID, action.User.Name, []string{})
	case "moreHelp":
		SendMoreHelp(api, action)
	}
}

// DialogReceived is responsible for handling received dialogs
func DialogReceived(api *slack.Client, payloadString string) {
	byteString := []byte(payloadString)
	dialog := models.DialogSubmission{}
	err := json.Unmarshal(byteString, &dialog)
	if err != nil {
		log.Fatalf("Dialog parsing error: %s", err)
		return
	}

	CreateFeedbackFromDialog(dialog)

}

// SendFeedbackSurvey sends the user a feedback survey
func SendFeedbackSurvey(api *slack.Client, action slackevents.MessageAction) {

	dialog := utils.GenerateFeedbackSurvey(action.TriggerID, action.CallbackID)

	err := api.OpenDialog(action.TriggerID, dialog)

	if err != nil {
		log.Fatalf("Error sending survey: %s", err)
	}
}

// SendMoreHelp sends the user information about the commands
func SendMoreHelp(api *slack.Client, action slackevents.MessageAction) {

	response :=
		"To query feedback use: \n   - `find param_name=param` \n   - Example: `find Type=intensives Sender=steve` \nTo delete feedback by id use: \n   - `delete ID` \n   - Example: `delete 28` "

	fmt.Println("Channel:", action.User.ID)
	api.NewRTM().PostMessage(action.Channel.ID, slack.MsgOptionText(response, false))
}
