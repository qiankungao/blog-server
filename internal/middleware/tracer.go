package middleware

import "github.com/gin-gonic/gin"

func Tracing() func(c *gin.Context) {

	return func(c *gin.Context) {
		c.Set("X-Trace-ID", 1)
		c.Set("X-Span-ID", 2)
		c.Next()
	}
}
