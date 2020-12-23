/*
 * @Author: EnderByEndera
 * @Date: 2020-12-23 13:17:08
 * @LastEditTime: 2020-12-23 13:43:48
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

	"github.com/gin-gonic/gin"
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
		Addr:    ":8080",
		Handler: router,
	}

	manageRouter()
}

// Start starts gin engine
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
	logger.Debugln("Server shutting down...")

	if err := srv.Shutdown(context.Background()); err != nil {
		logger.Fatalln("Server Shutdown with err: ", err)
	}
	logger.Debugln("Server shutted down")
}

func manageRouter() {
	router.Static("/static", staticRoute)
}
