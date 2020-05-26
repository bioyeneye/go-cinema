package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] %s %s %d %s\n",
			params.ClientIP,
			params.TimeStamp.Format(time.RFC1123),
			params.Method,
			params.Path,
			params.StatusCode,
			params.Latency)
	})
}