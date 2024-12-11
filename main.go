package main

import (
	"github.com/gin-gonic/gin"
	"github.com/webgoprojects/controllers"
	"github.com/webgoprojects/middleware"
)

func main() {
	r := gin.Default()

	// Middleware
	r.Use(middleware.Logger())

	// Routes
	r.GET("/", controllers.Index)
	r.GET("/about", controllers.About)

	// Serve static files
	r.Static("/static", "./static")

	// Start the server
	r.Run(":8080")
}
