/*
 * @Author: EnderByEndera
 * @Date: 2020-12-20 23:07:26
 * @LastEditTime: 2020-12-22 14:24:26
 * @LastEditors: Please set LastEditors
 * @Description: middleware.go makes a logger middleware for gin framework
 * @FilePath: /commdetection/logger/middleware.go
 */

package logger

import (
	"time"

	"github.com/gin-gonic/gin"
)

// Middleware creates a logger middleware for gin to use
func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := c.Request.Method
		reqURI := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		logger.Debugf("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqURI,
		)
	}
}
