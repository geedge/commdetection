/*
 * @Author: EnderByEndera
 * @Date: 2020-12-28 15:53:37
 * @LastEditTime: 2021-01-16 08:15:11
 * @LastEditors: Please set LastEditors
 * @Description: endpoint.go contains all the endpoints
 * @FilePath: /commdetection/router/endpoint.go
 */

package router

import (
	"commdetection/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

// handleErr handles all the error, set warning info in the command line, returns the error to the frontend
// and exit the endpoint
func handleErr(c *gin.Context, err error) {
	logger.Warnln(err)
	c.JSON(http.StatusOK, gin.H{
		"message": err.Error(),
	})
}
