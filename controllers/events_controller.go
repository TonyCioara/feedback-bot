package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/TonyCioara/feedback-bot/models"
	"github.com/TonyCioara/feedback-bot/utils"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

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

// DialogReceived is responsible for handling received dialogs
func DialogReceived(api *slack.Client, payloadString string) {
	byteString := []byte(payloadString)
	dialog := models.DialogSubmission{}
	err := json.Unmarshal(byteString, &dialog)
	if err != nil {
		fmt.Println("Dialog parsing error", err)
		return
	}

	// Save feedback in DB
	db, err := gorm.Open("sqlite3", "feedback-bot.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	db.AutoMigrate(&models.Feedback{})

	ftype := dialog.Submission["feedbackType"]
	if ftype == "other" {
		ftype = dialog.Submission["other"]
	}

	feedback := models.Feedback{
		UserID:       dialog.Submission["selectUser"],
		SenderID:     dialog.User["id"],
		SenderName:   dialog.User["name"],
		Good:         dialog.Submission["good"],
		Better:       dialog.Submission["better"],
		Best:         dialog.Submission["best"],
		FeedbackType: ftype,
	}

	db.Create(&feedback)

}

// SendFeedbackSurvey sends the user a feedback survey
func SendFeedbackSurvey(api *slack.Client, action slackevents.MessageAction) {

	dialog := utils.GenerateFeedbackSurvey(action.TriggerID, action.CallbackID)

	err := api.OpenDialog(action.TriggerID, dialog)

	fmt.Println("Error sending survey:", err)
}
