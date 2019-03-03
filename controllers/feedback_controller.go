package controllers

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/TonyCioara/feedback-bot/models"
	"github.com/TonyCioara/feedback-bot/utils"
	"github.com/nlopes/slack"
	"github.com/nlopes/slack/slackevents"
)

// CreateFeedbackFromDialog creates feedback given a dialog
func CreateFeedbackFromDialog(dialog models.DialogSubmission) {
	db := utils.StartAndMigrateDB()
	defer db.Close()

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

// SendFeedbackCSV sends a user all of their feedback
func SendFeedbackCSV(api *slack.Client, action slackevents.MessageAction) {
	sendFeedback(api, action, "CSV")
}

// SendFeedbackSheet sends a user all of their feedback
func SendFeedbackSheet(api *slack.Client, action slackevents.MessageAction) {
	sendFeedback(api, action, "sheet")
}

// SendFeedback sends a user all of their feedback
func sendFeedback(api *slack.Client, action slackevents.MessageAction, feedbackType string) {

	db := utils.StartAndMigrateDB()
	defer db.Close()

	var feedbacks []models.Feedback

	db.Where("user_id = ?", action.User.ID).Find(&feedbacks)

	fileName := "Feedback_" + action.User.Name + "_" + time.Now().Format("2006-01-02")
	row1 := []string{"ID", "Created", "Sender", "Type", "Good", "Better", "Best"}
	rows := [][]string{row1}
	for _, feedback := range feedbacks {
		var id = strconv.FormatUint(uint64(feedback.ID), 10)
		row := []string{id, feedback.CreatedAt.Format("2006-01-02"), feedback.SenderName,
			feedback.FeedbackType, feedback.Good, feedback.Better, feedback.Best}
		rows = append(rows, row)
	}

	switch ft := feedbackType; ft {
	case "sheet":
		fmt.Println("GOT TO 1")
		CreateSpreadsheet(fileName, rows)
	case "CSV":
		fmt.Println("GOT TO 2")
		utils.WriteCSV(fileName, rows)

		params := slack.FileUploadParameters{
			Title:    fileName,
			File:     fileName,
			Filename: fileName,
			Channels: []string{action.User.ID},
		}
		var _, err = api.UploadFile(params)
		if err != nil {
			log.Fatalf("Error: %s\n", err)
			return
		}

		utils.DeleteFile("./" + fileName)
	}

}

// DeleteFeedbackWithID deletes feedback given an ID
func DeleteFeedbackWithID(ID string, userID string) {
	db := utils.StartAndMigrateDB()
	defer db.Close()

	var feedback models.Feedback

	db.Where("user_id = ? AND id >= ?", userID, ID).First(&feedback)

}
