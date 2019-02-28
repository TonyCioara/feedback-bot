package controllers

import (
	"encoding/json"
	"log"
	"time"

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
	case "seeFeedback":
		SendFeedbackCSV(api, action)
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

	// Save feedback in DB
	db, err := gorm.Open("sqlite3", "feedback-bot.db")
	if err != nil {
		log.Fatalf("failed to connect database")
		return
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

	if err != nil {
		log.Fatalf("Error sending survey: %s", err)
	}
}

// SendFeedbackCSV sends a user all of their feedback
func SendFeedbackCSV(api *slack.Client, action slackevents.MessageAction) {

	db, err := gorm.Open("sqlite3", "feedback-bot.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	db.AutoMigrate(&models.Feedback{})

	var feedbacks []models.Feedback

	db.Where("user_id = ?", action.User.ID).Find(&feedbacks)

	csvName := "Feedback_" + action.User.Name + "_" + time.Now().Format("2006-01-02")
	row1 := []string{"Created", "Sender", "Type", "Good", "Better", "Best"}
	rows := [][]string{row1}
	for _, feedback := range feedbacks {
		row := []string{feedback.CreatedAt.Format("2006-01-02"), feedback.SenderName,
			feedback.FeedbackType, feedback.Good, feedback.Better, feedback.Best}
		rows = append(rows, row)
	}
	utils.WriteCSV(csvName, rows)

	params := slack.FileUploadParameters{
		Title:    csvName,
		File:     csvName,
		Filename: csvName,
		Channels: []string{action.User.ID},
	}
	_, err = api.UploadFile(params)
	if err != nil {
		log.Fatalf("%s\n", err)
		return
	}

	utils.DeleteFile("./" + csvName)
}
