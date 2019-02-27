package models

// DialogSubmission is used for parsing dialog submissions received by the api
type DialogSubmission struct {
	Type        string            `json:"type"`
	Token       string            `json:"token"`
	ActionTs    string            `json:"action_ts"`
	Team        map[string]string `json:"team"`
	User        map[string]string `json:"user"`
	Channel     map[string]string `json:"channel"`
	Submission  map[string]string `json:"submission"`
	CallbackID  string            `json:"callback_id"`
	ResponseURL string            `json:"response_url"`
	State       string            `json:"state"`
}
