/*
 * @Author: your name
 * @Date: 2020-12-19 11:49:14
 * @LastEditTime: 2021-01-05 18:09:49
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /commdetection/cmd/version.go
 */

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Ver means version of the system
const Ver = "command detection ver 0.0.5"

var verCmd = &cobra.Command{
	Use:   "version",
	Short: "show version",
	Long:  "show version of the application",
	Run: func(cmd *cobra.Command, args []string) {
		version()
	},
}

func version() {
	fmt.Println(Ver)
}
