/*
 * @Author: EnderByEndera
 * @Date: 2020-12-23 13:17:08
 * @LastEditTime: 2020-12-29 16:33:47
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

var (
	addr = yaml.GetPort()
	// StaticRoute sets the base route of the gin static filesystem
	StaticRoute = filepath.Join(os.Getenv("COMMDEPATH"), "static", "base")

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

// StartRouter starts and sets ending method for gin engine
func StartRouter() {
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
	// "/static" stores all the static file of the system
	router.Static("/static", StaticRoute)
	router.GET("/version", verEndpoint)
	evar := router.Group("/evaluation")
	{
		evar.POST("/upload", uploadCommEndpoint)
		evar.GET("/start", startEvaluateEndpoint)
	}
	commr := router.Group("/commrules")
	{
		commr.POST("/setrules", setCommrulesEndpoint)
		commr.GET("/getrules", getCommrulesEndpoint)
	}
	pathr := router.Group("/pathrules")
	{
		pathr.POST("/setrules", setPathrulesEndpoint)
		pathr.GET("/getrules", getPathrulesEndpoint)
	}
	rsr := router.Group("/rules")
	{
		rsr.POST("/setrules", setEvaluationRulesEndpoint)
		rsr.GET("/getrules", getEvaluationRulesEndpoint)
	}
}
