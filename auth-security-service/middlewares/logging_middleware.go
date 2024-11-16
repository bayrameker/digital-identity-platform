package middlewares

import (
    "github.com/gin-gonic/gin"
    "log"
    "time"
)

func LoggingMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        startTime := time.Now()
        c.Next()
        latency := time.Since(startTime)
        status := c.Writer.Status()

        log.Printf("%s %s %d %s", c.Request.Method, c.Request.URL.Path, status, latency)
    }
}
