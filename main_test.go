/*
 * @Author: EnderByEndera
 * @Date: 2020-12-04 15:03:42
 * @LastEditTime: 2020-12-15 11:51:34
 * @LastEditors: Please set LastEditors
 * @Description: Test main.go
 * @FilePath: /commdetection/main_test.go
 */
package main

import (
	"commdetection/yaml"
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	fmt.Println(yaml.GetYamlSetting("commrule"))
}
