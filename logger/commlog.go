/*
 * @Author: EnderByEndera
 * @Date: 2020-12-18 16:16:09
 * @LastEditTime: 2020-12-23 08:55:37
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /commdetection/commlog/commlog.go
 */

package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// logger is the logrus's logger engine
var (
	logger   = logrus.New()
	LogLevel = logrus.DebugLevel
)

func init() {
	formatter := &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	}
	logger.SetFormatter(formatter)
	logger.SetLevel(LogLevel)
	logger.SetOutput(os.Stdout)
}

// Debugln uses logrus.Debugln to set debug info
func Debugln(args ...interface{}) {
	logger.Debugln(args...)
}

// Warnln uses logrus.Warnln to set warn info
func Warnln(args ...interface{}) {
	logger.Warnln(args...)
}

// Fatalln uses logrus.Fatalln to set fatal info
func Fatalln(args ...interface{}) {
	logger.Fatalln(args...)
}

// Debugf uses logrus.Debugf to set debug formatted info
func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

// Warnf uses logrus.Warnf to set warning formatted info
func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}
