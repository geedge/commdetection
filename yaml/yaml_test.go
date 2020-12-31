/*
 * @Author: EnderByEndera
 * @Date: 2020-12-14 14:57:47
 * @LastEditTime: 2020-12-31 10:27:55
 * @LastEditors: Please set LastEditors
 * @Description: test yaml.go
 * @FilePath: /commdetection/yaml/yaml_test.go
 */

package yaml

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestGetYamlSetting(t *testing.T) {
	setting := GetYamlPathsSetting("commrules")
	if setting != filepath.Join(os.Getenv("COMMDEPATH"), "static", "rules", "commrules.json") {
		fmt.Println(setting)
		t.Errorf("get yaml setting failed")
	}
	setting = GetYamlPathsSetting("pathrules")
	if setting != filepath.Join(os.Getenv("COMMDEPATH"), "static", "rules", "pathrules.json") {
		fmt.Println(setting)
		t.Errorf("get yaml setting failed")
	}
	setting = GetYamlPathsSetting("rules")
	if setting != filepath.Join(os.Getenv("COMMDEPATH"), "static", "rules", "rules.json") {
		fmt.Println(setting)
		t.Errorf("get yaml setting failed")
	}
}

func BenchmarkGetYamlSetting(b *testing.B) {
	GetYamlPathsSetting("commrules")
	GetYamlPathsSetting("pathrules")
	GetYamlPathsSetting("rules")
}
