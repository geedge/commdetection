/*
 * @Author: EnderByEndera
 * @Date: 2020-12-23 13:17:08
 * @LastEditTime: 2020-12-28 16:10:14
 * @LastEditors: Please set LastEditors
 * @Description: router.go defines gin.Engine and its router
 * @FilePath: /commdetection/route/router.go
 */

package router

import (
	"commdetection/logger"
	"commdetection/yaml"
	"context"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	addr = ":8080"
)

var (
	staticRoute = filepath.Join(os.Getenv("COMMDEPATH"), "static")

	router *gin.Engine
	srv    *http.Server
)

func init() {
	// initilize gin-gonic backend engine
	ginmode, err := yaml.GetGinMode()
	if err != nil {
		logger.Warnln(err)
	}
	gin.SetMode(ginmode)
	logger.Debugf("Sets Gin mode to \"%s\" mode", ginmode)

	router = gin.New()
	router.Use(gin.Recovery())
	router.Use(logger.Middleware())

	srv = &http.Server{
		Addr:    addr,
		Handler: router,
	}

	manageRouter()
}

// Start starts and sets ending method for gin engine
func Start() {
	go func() {
		// engine start on http://127.0.0.1:8000
		logger.Debugln("Server started at", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalln(err)
		}
	}()

	// Make a quit channel to receive SIGINT and SIGTERM signal
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Debugln("Server starts shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatalln("Server Shutdown with err: ", err)
	}
	logger.Debugln("Server shutted down")
	cancel()
}

func manageRouter() {
	// `/static` stores all the static file of the system
	router.Static("/static", staticRoute)
	router.GET("/version", verEndpoint)
	eva := router.Group("/eva")
	{
		eva.POST("/upload", uploadCommEndpoint)
		eva.GET("/start", startEvaluateEndpoint)
	}
	comm := router.Group("/comm")
	{
		comm.POST("/rules", setCommrulesEndpoint)
	}
}
