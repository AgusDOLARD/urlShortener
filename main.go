package main

import (
	"log"
	"urlShortener/controllers"
	"urlShortener/initializers"

	"github.com/gofiber/fiber/v2"
)

func init() {
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	router := fiber.New()

	router.Post("/", controllers.URLCreate)

	router.Get("/:slug", controllers.URLSolve)

	log.Fatal(router.Listen(":8080"))
}
