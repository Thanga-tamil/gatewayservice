package middleware

import (
	"time"
	"github.com/gin-gonic/gin"
	log "gateway/internal/utils"
)

func ReqFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		log.Infof("Incoming Req Method: '%s' Path ---> %s Params: %s", 
		c.Request.Method, 
		c.Request.URL, 
		c.Request.URL.Query())

		// Pre-handler phase
		c.Next()

		// Post-handler phase
		latency := time.Since(start)
		log.Debug("Request took %v", latency)

		code := c.Writer.Status()

		var x func(msg string, code int)
		switch {
		case code >= 200 && code <= 300:
			x = func(msg string, code int) {
				log.Info(msg, code)
			}
		case code >= 300 && code <= 400:
			x = func(msg string, code int) {
				log.Debug(msg, code)
			}
		case  code >= 400 && code <= 600:
			x = func(msg string, code int) {
				log.Error(msg, code)
			}
		}

		x("HTTP statusCode => ", code)

	}
}
