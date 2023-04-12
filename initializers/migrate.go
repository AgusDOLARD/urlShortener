package initializers

import "urlShortener/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.URL{})
}
