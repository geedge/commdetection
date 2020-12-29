/*
 * @Author: EnderByEndera
 * @Date: 2020-12-04 15:03:24
 * @LastEditTime: 2020-12-23 16:11:29
 * @LastEditors: Please set LastEditors
 * @Description: Main Func Entry, use flags to give help
 * @FilePath: /commdetection/main.go
 */

package main

import (
	"commdetection/cmd"
	"commdetection/logger"
	"commdetection/router"

	"github.com/sirupsen/logrus"
)

func init() {
	logger.LogLevel = logrus.Level(cmd.GetLogLevel())
}

func main() {
	if err := cmd.Execute(); err != nil {
		logger.Fatalln(err)
	}
	router.StartRouter()
}
