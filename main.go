package main

import (
	"log"

	"github.com/AgusDOLARD/urlShortener/controllers"
	"github.com/AgusDOLARD/urlShortener/database"
	"github.com/gofiber/fiber/v2"
)

func init() {
	db.Connect()
	db.Migrate()
}

func main() {
	router := fiber.New()

	router.Static("/", "./public")

	router.Post("/", controllers.URLCreate)

	router.Get("/:slug", controllers.URLSolve)

	log.Fatal(router.Listen(":8080"))
}
