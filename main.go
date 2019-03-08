package main

import (
	"github.com/TonyCioara/feedback-bot/slack"

	_ "github.com/joho/godotenv/autoload"
)

// main is our entrypoint, where the application initializes the Slackbot.
// export GO111MODULE=on; go run main.go
func main() {
	slack.CreateSlackClient()
}
