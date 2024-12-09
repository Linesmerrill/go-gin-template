package middleware

import (
        "time"

        "github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
        return func(c *gin.Context) {
                // Start timer
                start := time.Now()

                // Process request
                c.Next()

                // Log the request
                latency := time.Since(start)
                clientIP := c.ClientIP()
                method := c.Request.Method
                path := c.Request.URL.Path
                statusCode := c.Writer.Status()

                logger.Printf("[%s] %s %s %d %v\n", clientIP, method, path, statusCode, latency)
        }
}