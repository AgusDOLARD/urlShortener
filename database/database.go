package db

import (
	"log"

	"github.com/AgusDOLARD/urlShortener/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Instance database
var Instance *gorm.DB

// Connect to database
func Connect() {
	var err error
	Instance, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
}

// Migrate URL model to database
func Migrate() {
	Instance.AutoMigrate(&models.URL{})
}
