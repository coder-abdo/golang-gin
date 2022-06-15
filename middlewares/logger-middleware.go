package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - %s - [%s] - %s %d %s",
			params.ClientIP,
			params.Method,
			params.TimeStamp.Format(time.RFC822Z),
			params.Path,
			params.StatusCode,
			params.Latency,
		)
	})
}
