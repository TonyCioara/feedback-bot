package server

import (
	"os"
	"log"

	"github.com/TonyCioara/feedback-bot/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB a reference to the PostgresQl DB
var DB *gorm.DB

// StartAndMigrateDB starts and migrates the database
func StartAndMigrateDB() {

	password := os.Getenv("DBPASSWORD")
	user := os.Getenv("DBUSER")
	host := os.Getenv("DBHOST")
	port := os.Getenv("DBPORT")
	name := os.Getenv("DBNAME")

	params := "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + name + " sslmode=disable"

	db, err := gorm.Open("postgres", params)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %s", err)
	}

	db.AutoMigrate(&models.Feedback{})
	db.AutoMigrate(&models.User{})

	DB = db
}
