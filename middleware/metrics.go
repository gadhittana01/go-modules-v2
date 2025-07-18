package middleware

import (
	"strconv"
	"time"

	"github.com/gadhittana01/go-modules-v2/metrics"
	"github.com/gin-gonic/gin"
)

// MetricsMiddleware tracks HTTP requests and their duration
func MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Process request
		c.Next()

		// Calculate duration
		duration := time.Since(start).Seconds()

		// Get status code
		status := strconv.Itoa(c.Writer.Status())

		// Record metrics
		metrics.HttpRequestsTotal.WithLabelValues(
			c.Request.Method,
			c.FullPath(),
			status,
		).Inc()

		metrics.HttpRequestDuration.WithLabelValues(
			c.Request.Method,
			c.FullPath(),
		).Observe(duration)
	}
}
