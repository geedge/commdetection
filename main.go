/*
 * @Author: your name
 * @Date: 2020-12-04 15:03:24
 * @LastEditTime: 2020-12-15 11:46:55
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /commdetection/main.go
 */
package main

import (
	"commdetection/yaml"
	"fmt"
)

func main() {
	fmt.Println(yaml.GetYamlSetting("commrule"))
}
