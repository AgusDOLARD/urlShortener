package main

import (
	"urlShortener/initializers"
	"urlShortener/models"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.URL{})
}
