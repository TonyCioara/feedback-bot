package main

import (
	"fmt"
	"os"

	"github.com/TonyCioara/feedback-bot/slack"

	_ "github.com/joho/godotenv/autoload"
)

// main is our entrypoint, where the application initializes the Slackbot.
// DO NOT EDIT THIS FUNCTION. This is a fully complete implementation.
// export GO111MODULE=on; go run main.go
func main() {
	if len(os.Args) >= 2 && os.Args[1] == "my_job_param" {
		// Execute job.
		fmt.Println("Reload")
		return
	}
	// Set up my web server with port, router, etc.
	botToken := os.Getenv("BOT_OAUTH_ACCESS_TOKEN")
	slackClient := slack.CreateSlackClient(botToken)
	slack.RespondToEvents(slackClient)
}
