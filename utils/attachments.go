package utils

import "github.com/nlopes/slack"

func GenerateHelpButtons() slack.Attachment {

	action1 := slack.AttachmentAction{
		Name:  "option",
		Text:  "Send Feedback",
		Type:  "button",
		Value: "sendFeedback",
	}
	action2 := slack.AttachmentAction{
		Name:  "option",
		Text:  "See My Feedback",
		Type:  "button",
		Value: "seeFeedback",
	}

	attachment := slack.Attachment{
		Text:       "Pick an option",
		Color:      "#3AA3E3",
		Fallback:   "You are unable to select at option",
		CallbackID: "helpButton",
		Actions:    []slack.AttachmentAction{action1, action2},
	}

	return attachment
}
