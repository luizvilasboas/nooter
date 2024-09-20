package middlewares

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Next()

		latency := time.Since(t)
		log.Printf("[%s] %s - Latência: %v", c.Request.Method, c.Request.URL, latency)
	}
}
