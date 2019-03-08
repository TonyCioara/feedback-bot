package slack

import (
	"fmt"

	"github.com/TonyCioara/feedback-bot/controllers"
	"github.com/TonyCioara/feedback-bot/server"
	"github.com/jasonlvhit/gocron"
)

// CreateSlackClient creates and api, rtm and sets up events responding
func CreateSlackClient() {
	server.StartAndMigrateDB()
	server.InitializeAPIAndRTM()

	go controllers.SetUpEventsAPI()
	go server.RTM.ManageConnection()
	go ScheduleWeeklyFeedback()

	controllers.RespondToEvents()
}

// ScheduleWeeklyFeedback schedules the weekly feedback
func ScheduleWeeklyFeedback() {
	fmt.Println("I'm here!")
	s := gocron.NewScheduler()
	s.Every(1).Monday().At("10:00").Do(controllers.DeliverWeeklyFeedback)
	<-s.Start()
}
