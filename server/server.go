package server

import (
	"os"

	"github.com/nlopes/slack"
)

// API handles events on the Slack API
var API *slack.Client

// RTM hadles real time messages on the Slack API
var RTM *slack.RTM

func InitializeAPIAndRTM() {
	apiKey := os.Getenv("BOT_OAUTH_ACCESS_TOKEN")
	API = slack.New(apiKey)
	RTM = API.NewRTM()
}
