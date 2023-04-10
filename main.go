package main

import (
	"urlShortener/controllers"
	"urlShortener/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	router := gin.Default()

	router.POST("/", controllers.URLCreate)

	router.GET("/:slug", controllers.URLSolve)

	router.Run(":3000")
}
