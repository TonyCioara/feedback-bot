package models

import "github.com/jinzhu/gorm"

// Feedback is used for storing feedback in the database
type Feedback struct {
	gorm.Model
	User   string
	Sender string
	Good   string
	Better string
	Best   string
	Type   string
}
