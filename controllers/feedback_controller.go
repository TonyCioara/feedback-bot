package controllers

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/TonyCioara/feedback-bot/models"
	"github.com/TonyCioara/feedback-bot/server"
	"github.com/TonyCioara/feedback-bot/utils"
	"github.com/nlopes/slack"
	"github.com/nlopes/slack/slackevents"
)

// CreateFeedbackFromDialog creates feedback given a dialog
func CreateFeedbackFromDialog(dialog models.DialogSubmission) {
	db := server.DB

	ftype := dialog.Submission["feedbackType"]
	if ftype == "other" {
		ftype = dialog.Submission["other"]
	}

	feedback := models.Feedback{
		UserID:     dialog.Submission["selectUser"],
		SenderID:   dialog.User["id"],
		Sender:     dialog.User["name"],
		Good:       dialog.Submission["good"],
		Better:     dialog.Submission["better"],
		Best:       dialog.Submission["best"],
		Type:       ftype,
		SentWeekly: false,
	}

	db.Create(&feedback)

	var user models.User

	db.Where("user_id = ?", dialog.Submission["selectUser"]).First(&user)

	if user.ID != 0 {
		return
	}
	user = models.User{
		UserID:             dialog.Submission["selectUser"],
		ActiveSubscription: true,
	}

	db.Create(&user)
}

// SendFeedbackSurvey sends the user a feedback survey
func SendFeedbackSurvey(action slackevents.MessageAction) {
	api := server.API

	dialog := utils.GenerateFeedbackSurvey(action.TriggerID, action.CallbackID)

	err := api.OpenDialog(action.TriggerID, dialog)

	if err != nil {
		log.Fatalf("Error sending survey: %s", err)
	}
}

// SendFeedbackCSV sends a user all of their feedback
func SendFeedbackCSV(userID, userName string, queryParams []string) {
	db := server.DB
	api := server.API

	queryFeedback := map[string]interface{}{"user_id": userID}

	for index, param := range queryParams {
		if index == 0 {
			continue
		}
		elements := strings.Split(param, "=")
		if len(elements) != 2 {
			continue
		}
		key := strings.ToLower(elements[0])
		queryFeedback[key] = elements[1]
	}

	var feedbacks []models.Feedback

	db.Where(queryFeedback).Find(&feedbacks)

	fileName := "Feedback_" + userName + "_" + time.Now().Format("2006-01-02") + ".csv"
	row1 := []string{"ID", "Created", "Sender", "Type", "Good", "Better", "Best"}
	rows := [][]string{row1}
	for _, feedback := range feedbacks {
		var id = strconv.FormatUint(uint64(feedback.ID), 10)
		row := []string{id, feedback.CreatedAt.Format("2006-01-02"), feedback.Sender,
			feedback.Type, feedback.Good, feedback.Better, feedback.Best}
		rows = append(rows, row)
	}

	utils.WriteCSV(fileName, rows)

	params := slack.FileUploadParameters{
		Title:    fileName,
		File:     fileName,
		Filename: fileName,
		Channels: []string{userID},
	}
	var _, err = api.UploadFile(params)
	if err != nil {
		log.Fatalf("Error: %s\n", err)
		return
	}

	utils.DeleteFile("./" + fileName)

}

// DeleteFeedback deletes feedback
func DeleteFeedback(slackEvent *slack.MessageEvent, elements []string) {
	db := server.DB
	slackClient := server.RTM

	userID := slackEvent.User
	ID := elements[1]
	fmt.Println("id:", userID)

	var feedback models.Feedback

	db.Where("id >= ?", ID).First(&feedback)

	fmt.Println("feedback:", feedback, "UID:", userID)
	if feedback.UserID != userID {
		slackClient.PostMessage(slackEvent.Channel, slack.MsgOptionText("Invalid input", false))
		return
	}

	db.Delete(&feedback)
	slackClient.PostMessage(slackEvent.Channel, slack.MsgOptionText("Success", false))
}

// DeliverWeeklyFeedback sends users all feedback from the past week
func DeliverWeeklyFeedback() {
	db := server.DB
	api := server.RTM

	var feedbacks []models.Feedback

	db.Where("sent_weekly = ?", false).Find(&feedbacks)

	feedbackMap := make(map[string][][]string)

	row1 := []string{"ID", "Created", "Sender", "Type", "Good", "Better", "Best"}

	for _, feedback := range feedbacks {
		var id = strconv.FormatUint(uint64(feedback.ID), 10)
		row := []string{id, feedback.CreatedAt.Format("2006-01-02"), feedback.Sender,
			feedback.Type, feedback.Good, feedback.Better, feedback.Best}

		if feedbackMap[feedback.UserID] == nil {
			feedbackMap[feedback.UserID] = [][]string{row1, row}
		} else {
			feedbackMap[feedback.UserID] = append(feedbackMap[feedback.UserID], row)
		}
		feedback.SentWeekly = true
		db.Save(&feedback)
	}

	done := make(chan bool)

	for key, value := range feedbackMap {
		go func(key string, value [][]string) {
			var user models.User
			db.Where("user_id = ?", key).First(&user)

			if user.ActiveSubscription == false {
				return
			}

			fileName := "Weekly_Feedback_" + key + "_" + time.Now().Format("2006-01-02") + ".csv"

			utils.WriteCSV(fileName, value)

			params := slack.FileUploadParameters{
				Title:          fileName,
				File:           fileName,
				Filename:       fileName,
				InitialComment: "*Here is your weekly feedback!* \n   - To unsubscribe from weekly feedback type `unsubscribe` \n   - For more information type `help`",
				Channels:       []string{user.UserID},
			}
			var _, err = api.UploadFile(params)
			if err != nil {
				log.Fatalf("Error: %s\n", err)
				return
			}

			utils.DeleteFile("./" + fileName)

			done <- true
		}(key, value)
	}

	for i := 0; i < len(feedbackMap); i++ {
		<-done
	}

}
