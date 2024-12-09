package main

import (
        "github.com/gin-gonic/gin"
        "your-project-name/controllers"
        "your-project-name/middleware"
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