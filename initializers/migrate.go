package initializers

import "github.com/AgusDOLARD/urlShortener/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.URL{})
}
