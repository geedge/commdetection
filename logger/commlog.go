/*
 * @Author: EnderByEndera
 * @Date: 2020-12-18 16:16:09
 * @LastEditTime: 2020-12-19 10:47:01
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /commdetection/commlog/commlog.go
 */

package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

// Init initializes
func Init(logLevel logrus.Level, formatter logrus.Formatter) {
	logger.SetFormatter(formatter)
	logger.SetLevel(logLevel)
	logger.SetOutput(os.Stdout)
}

// Warnln uses logrus.Warnln to set warn info
func Warnln(args ...interface{}) {
	logger.Warnln(args)
}

// Debugln uses logrus.Debugln to set debug info
func Debugln(args ...interface{}) {
	logger.Debugln(args)
}

// Fatalln uses logrus.Fatalln to set fatal info
func Fatalln(args ...interface{}) {
	logger.Fatalln(args)
}

// Warnf uses logrus.Warnf to set warning formatted info
func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args)
}
