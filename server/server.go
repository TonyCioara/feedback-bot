package server

import (
	"github.com/nlopes/slack"
)

// API handles events on the Slack API
var API *slack.Client

// RTM hadles real time messages on the Slack API
var RTM *slack.RTM
