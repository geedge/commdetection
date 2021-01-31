/*
 * @Author: EnderByEndera
 * @Date: 2021-01-16 08:14:22
 * @LastEditTime: 2021-02-01 06:27:20
 * @LastEditors: Please set LastEditors
 * @Description: comm_endpoint.go has the function for all the commands' endpoint
 * @FilePath: /commdetection/router/comm_endpoint.go
 */

package router

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// uploadCommEndpoint uploads the file storing commands to the system
func uploadCommEndpoint(c *gin.Context) {
	commfile, err := c.FormFile("commfile")
	if err != nil {
		handleErr(c, err)
		return
	}

	filename := filepath.Join(StaticRoute, filepath.Base(commfile.Filename))
	if err := c.SaveUploadedFile(commfile, filename); err != nil {
		handleErr(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
