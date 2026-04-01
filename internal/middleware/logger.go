package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// process request
		c.Next()

		duration := time.Since(start)

		requestID, _ := c.Get("request_id")

		// basic request log
		log.Printf(
			"[%s] %s | %d | %s | %s %s",
			requestID,
			c.ClientIP(),
			c.Writer.Status(),
			duration,
			c.Request.Method,
			c.Request.URL.Path,
		)

		// log errors if any
		if len(c.Errors) > 0 {
			log.Printf(
				"[%s] ERROR: %s",
				requestID,
				c.Errors.String(),
			)
		}
	}
}
