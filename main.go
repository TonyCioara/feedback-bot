package main

import (
	"fmt"
	"os"

	"github.com/TonyCioara/feedback-bot/controllers"
	"github.com/TonyCioara/feedback-bot/slack"

	_ "github.com/joho/godotenv/autoload"
)

// main is our entrypoint, where the application initializes the Slackbot.
// export GO111MODULE=on; go run main.go
func main() {
	// Set up my web server with port, router, etc.

	if len(os.Args) >= 2 {
		if os.Args[1] == "weekly_feedback" {
			fmt.Println("Sending out Weekly Feedback")
			controllers.DeliverWeeklyFeedback()
			return
		}
	}
	slack.CreateSlackClient()
}
