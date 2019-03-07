package models

import "github.com/jinzhu/gorm"

// User is used for storing user preference
type User struct {
	gorm.Model
	UserID             string
	ActiveSubscription bool
}
