/*
 * @Author: EnderByEndera
 * @Date: 2020-12-20 23:27:48
 * @LastEditTime: 2020-12-22 15:09:24
 * @LastEditors: Please set LastEditors
 * @Description: test middleware.go
 * @FilePath: /commdetection/logger/middleware_test.go
 */

package logger

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func init() {
	r := gin.New()
	r.Use(Middleware())
	r.Use(gin.Recovery())
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "world",
		})
	})
	go r.Run()
}

func BenchmarkMiddleware(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		http.Get("127.0.0.1:8080/hello")
	}
}
