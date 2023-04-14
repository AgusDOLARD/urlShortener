package controllers

import (
	"net/url"

	"github.com/AgusDOLARD/urlShortener/initializers"
	"github.com/AgusDOLARD/urlShortener/models"
	"github.com/gofiber/fiber/v2"
	"github.com/twharmon/gouid"
)

// URLSolve GET "/:slug" lookup the slug in the database and return the full url
func URLSolve(c *fiber.Ctx) error {
	var u models.URL

	slug := c.Params("slug")
	if slug == "" {
		return fiber.NewError(fiber.StatusBadRequest, "slug not provided")
	}

	initializers.DB.First(&u, "slug = ?", slug)
	if u.ID == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "failed to find the slug")
	}

	return c.Redirect(u.Full)
}

// URLCreate POST "/" given {full} generate a slug and save it to the database then return it
func URLCreate(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var u models.URL

	if err := c.BodyParser(&u); err != nil {
		return err
	}

	_, err := url.ParseRequestURI(u.Full)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "failed to parse the url")
	}

	u.Slug = gouid.String(6, gouid.MixedCaseAlphaNum)

	res := initializers.DB.Create(&u)
	if res.Error != nil {
		return fiber.NewError(fiber.StatusBadRequest, "failed to generate a new entry in the db")
	}

	return c.JSON(u)
}
