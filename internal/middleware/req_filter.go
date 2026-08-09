package middleware

import (
	"time"
	"github.com/gin-gonic/gin"
	log "gateway/internal/utils"
)

func ReqFilter() gin.HandlerFunc {
  return func(c *gin.Context) {
    start := time.Now()

	log.Infof("Incoming Req Method: %s Path ---> %s QueryParam: %s", 
			  c.Request.Method, 
			  c.Request.URL, 
			  c.Request.URL.Query())

    // Pre-handler phase
    c.Next()

    // Post-handler phase
    latency := time.Since(start)
	log.Infof("Request took %v", latency)
  }
}
