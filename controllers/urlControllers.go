package controllers

import (
	"regexp"
	"urlShortener/initializers"
	"urlShortener/models"

	"github.com/gofiber/fiber/v2"
	"github.com/twharmon/gouid"
)

// URLSolve GET "/:slug" lookup the slug in the database and return the full url
func URLSolve(c *fiber.Ctx) error {
	var url models.URL

	slug := c.Params("slug")
	if slug == "" {
		return fiber.NewError(fiber.StatusBadRequest, "slug not provided")
	}

	initializers.DB.First(&url, "slug = ?", slug)
	if url.ID == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "failed to find the slug")
	}

	return c.Redirect(url.Full)
}

// URLCreate POST "/" given {full} generate a slug and save it to the database then return it
func URLCreate(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var url models.URL

	if err := c.BodyParser(&url); err != nil {
		return err
	}

	if matched, err := regexp.MatchString(`^(https?)://[^\s/$.?#].[^\s]*$`, url.Full); err != nil ||
		!matched {
		return fiber.NewError(fiber.StatusBadRequest, "failed to parse the url")
	}

	url.Slug = gouid.String(6, gouid.MixedCaseAlphaNum)

	res := initializers.DB.Create(&url)
	if res.Error != nil {
		return fiber.NewError(fiber.StatusBadRequest, "failed to generate a new entry in the db")
	}

	return c.JSON(url)
}
