package slack

import (
	"os"

	"github.com/TonyCioara/feedback-bot/controllers"
	"github.com/TonyCioara/feedback-bot/server"
	"github.com/nlopes/slack"
)

// CreateSlackClient creates and api, rtm and sets up events responding
func CreateSlackClient() {
	server.StartAndMigrateDB()

	apiKey := os.Getenv("BOT_OAUTH_ACCESS_TOKEN")
	server.API = slack.New(apiKey)

	go controllers.SetUpEventsAPI()

	server.RTM = server.API.NewRTM()
	go server.RTM.ManageConnection()

	controllers.RespondToEvents()
}
