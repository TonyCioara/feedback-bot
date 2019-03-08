package utils

import "github.com/nlopes/slack"

// GenerateHelpButtons generates help buttons
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
	action3 := slack.AttachmentAction{
		Name:  "option",
		Text:  "More Help",
		Type:  "button",
		Value: "moreHelp",
		Style: "danger",
	}

	attachment := slack.Attachment{
		Text:       "Pick an option",
		Color:      "#3AA3E3",
		Fallback:   "You are unable to select an option",
		CallbackID: "helpButton",
		Actions:    []slack.AttachmentAction{action1, action2, action3},
	}

	return attachment
}

// GenerateFeedbackSurvey generates a feedback survey
func GenerateFeedbackSurvey(triggerID, callbackID string) slack.Dialog {
	var dialogElement1 slack.DialogElement = map[string]string{
		"type":        "select",
		"label":       "Who is this feedback for?",
		"name":        "selectUser",
		"data_source": "users",
	}
	var dialogElement2 slack.DialogElement = map[string]interface{}{
		"type":  "select",
		"label": "What is this feedback for?",
		"name":  "feedbackType",
		"options": []map[string]string{
			{
				"label": "Intensives",
				"value": "intensives",
			},
			{
				"label": "Unconferences and Lightning talks",
				"value": "unconferences",
			},
			{
				"label": "Project",
				"value": "project",
			},
			{
				"label": "Class",
				"value": "class",
			},
			{
				"label": "Coaching",
				"value": "coaching",
			},
			{
				"label": "Other",
				"value": "other",
			},
		},
	}
	var dialogElement3 slack.DialogElement = map[string]string{
		"type":     "text",
		"label":    "If you selected Other please elaborate",
		"name":     "other",
		"optional": "true",
	}
	var dialogElement4 slack.DialogElement = map[string]string{
		"type":     "textarea",
		"label":    "What has the person done that was good?",
		"name":     "good",
		"optional": "true",
	}
	var dialogElement5 slack.DialogElement = map[string]string{
		"type":     "textarea",
		"label":    "What could the person do to improve?",
		"name":     "better",
		"optional": "true",
	}
	var dialogElement6 slack.DialogElement = map[string]string{
		"type":     "textarea",
		"label":    "What has the person done that was best?",
		"name":     "best",
		"optional": "true",
	}

	dialogElements := []slack.DialogElement{dialogElement1, dialogElement2,
		dialogElement3, dialogElement4, dialogElement5, dialogElement6}

	dialog := slack.Dialog{
		TriggerID:      triggerID,
		CallbackID:     callbackID,
		Title:          "Submit Feedback",
		SubmitLabel:    "Submit",
		NotifyOnCancel: true,
		Elements:       dialogElements,
	}
	return dialog
}
