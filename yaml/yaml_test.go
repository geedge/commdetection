/*
 * @Author: EnderByEndera
 * @Date: 2020-12-14 14:57:47
 * @LastEditTime: 2020-12-28 17:27:00
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
	if setting != "/root/go/src/commdetection/static/rules/commrules.json" {
		fmt.Println(setting)
		t.Errorf("get yaml setting failed")
	}
	setting = GetYamlPathsSetting("pathrules")
	if setting != "/root/go/src/commdetection/static/rules/pathrules.json" {
		fmt.Println(setting)
		t.Errorf("get yaml setting failed")
	}
	setting = GetYamlPathsSetting("rules")
	if setting != "/root/go/src/commdetection/static/rules/rules.json" {
		fmt.Println(setting)
		t.Errorf("get yaml setting failed")
	}
}

func BenchmarkGetYamlSetting(b *testing.B) {
	GetYamlPathsSetting("commrules")
	GetYamlPathsSetting("pathrules")
	GetYamlPathsSetting("rules")
}
