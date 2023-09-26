package middleware

import (
	"github.com/alexcesaro/log/stdlog"
	"github.com/gin-gonic/gin"
)

func RequestLoggerActivity() gin.HandlerFunc {
	return func(c *gin.Context) {
		log := stdlog.GetFromFlags()
		log.Info("HTTP Request Method", "Method=", c.Request.Method, "URL=", c.Request.URL)
	}
}
