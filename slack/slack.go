package slack

import (
	"github.com/TonyCioara/feedback-bot/controllers"
	"github.com/TonyCioara/feedback-bot/server"
)

// CreateSlackClient creates and api, rtm and sets up events responding
func CreateSlackClient() {
	server.StartAndMigrateDB()
	server.InitializeAPIAndRTM()

	go controllers.SetUpEventsAPI()
	go server.RTM.ManageConnection()

	controllers.RespondToEvents()
}

// DeliverWeeklyFeedback delivers last week's feedback to
// all users who are subscribed to Weekly Feedback
func DeliverWeeklyFeedback() {
	server.StartAndMigrateDB()
	server.InitializeAPIAndRTM()

	controllers.DeliverWeeklyFeedback()
}
