package middleware

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func ProcessingTimeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		durationMicros := time.Since(start).Microseconds()
		c.Header(
			"X-Processing-Time-Micros",
			strconv.FormatInt(durationMicros, 10),
		)
	}
}
