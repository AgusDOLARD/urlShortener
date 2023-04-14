package main

import (
	"log"

	"github.com/AgusDOLARD/urlShortener/controllers"
	"github.com/AgusDOLARD/urlShortener/initializers"
	"github.com/gofiber/fiber/v2"
)

func init() {
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	router := fiber.New()

	router.Static("/", "./public")

	router.Post("/", controllers.URLCreate)

	router.Get("/:slug", controllers.URLSolve)

	log.Fatal(router.Listen(":8080"))
}
