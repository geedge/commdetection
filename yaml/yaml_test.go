/*
 * @Author: EnderByEndera
 * @Date: 2020-12-14 14:57:47
 * @LastEditTime: 2020-12-17 10:02:22
 * @LastEditors: Please set LastEditors
 * @Description: test yaml.go
 * @FilePath: /commdetection/yaml/yaml_test.go
 */

package yaml

import (
	"fmt"
	"testing"
)

func TestGetYamlSetting(t *testing.T) {
	setting := GetYamlPathsSetting("commrules")
	if setting != "/root/go/src/commdetection/rules/rulesjson/commrules.json" {
		fmt.Println(setting)
		t.Errorf("get yaml setting failed")
	}
	setting = GetYamlPathsSetting("pathrules")
	if setting != "/root/go/src/commdetection/rules/rulesjson/pathrules.json" {
		fmt.Println(setting)
		t.Errorf("get yaml setting failed")
	}
	setting = GetYamlPathsSetting("rules")
	if setting != "/root/go/src/commdetection/rules/rulesjson/rules.json" {
		fmt.Println(setting)
		t.Errorf("get yaml setting failed")
	}
}

func BenchmarkGetYamlSetting(b *testing.B) {
	GetYamlPathsSetting("commrules")
	GetYamlPathsSetting("pathrules")
	GetYamlPathsSetting("rules")
}
