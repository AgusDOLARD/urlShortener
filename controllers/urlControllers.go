package controllers

import (
	"urlShortener/initializers"
	"urlShortener/models"

	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

// URLSolve lookup the slug in the database and return the full url
func URLSolve(c *gin.Context) {
	var url models.URL
  if err := c.ShouldBindUri(&url); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}
	initializers.DB.First(&url, "Slug = ?", url.Slug)
	c.JSON(200, gin.H{"full": url.Full})
}

// URLCreate given a url generate a slug and save it to the database
func URLCreate(c *gin.Context) {
	var url models.URL

	if err := c.BindJSON(&url); err != nil {
		return
	}

	url.Slug = gonanoid.Must(6)

	initializers.DB.Create(&url)

	c.JSON(200, gin.H{"slug": url.Slug})
}
