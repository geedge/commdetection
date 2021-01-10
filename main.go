/*
 * @Author: EnderByEndera
 * @Date: 2020-12-04 15:03:24
 * @LastEditTime: 2021-01-06 09:50:51
 * @LastEditors: Please set LastEditors
 * @Description: Main Func Entry, use flags to give help
 * @FilePath: /commdetection/main.go
 */

package main

import (
	"commdetection/cmd"
	"commdetection/logger"
)

func main() {
	if err := cmd.Execute(); err != nil {
		logger.Fatalln(err)
	}
}
