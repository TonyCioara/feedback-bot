package utils

import (
	"fmt"

	"github.com/TonyCioara/feedback-bot/models"
	"github.com/jinzhu/gorm"
)

func StartAndMigrateDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "feedback-bot.db")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.Feedback{})

	fmt.Println("Got here!")
	return db
}
