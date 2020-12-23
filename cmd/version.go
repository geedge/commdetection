/*
 * @Author: your name
 * @Date: 2020-12-19 11:49:14
 * @LastEditTime: 2020-12-22 16:12:16
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /commdetection/cmd/version.go
 */

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var verCmd = &cobra.Command{
	Use:   "version",
	Short: "show version",
	Long:  "show version of the application",
	Run: func(cmd *cobra.Command, args []string) {
		version()
	},
}

func version() {
	fmt.Println("ver 0.0.1")
}
