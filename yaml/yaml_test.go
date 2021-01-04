/*
 * @Author: EnderByEndera
 * @Date: 2020-12-14 14:57:47
 * @LastEditTime: 2021-01-04 15:34:16
 * @LastEditors: Please set LastEditors
 * @Description: test yaml.go
 * @FilePath: /commdetection/yaml/yaml_test.go
 */

package yaml

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

func TestGetYamlSetting(t *testing.T) {
	t.Run("Test GetPort", func(t *testing.T) {
		setting := GetPort()
		if !strings.HasPrefix(setting, ":") {
			t.Errorf("Wrong setting: \"%s\"", setting)
		}
		if port, err := strconv.Atoi(setting[1:]); err != nil {
			t.Error(err)
		} else if port <= 0 || port >= 65536 {
			t.Errorf("Wrong port: %d", port)
		}
	})
	t.Run("Test GetGinmode", func(t *testing.T) {
		setting, err := GetGinMode()
		if err != nil {
			t.Error(err)
		}
		if setting != "release" && setting != "debug" && setting != "test" {
			t.Errorf("Wrong setting: %s", setting)
		}
	})
	t.Run("Test GetCommrulesSetting", func(t *testing.T) {
		setting := GetCommRulesSetting()
		if setting != "" && !filepath.HasPrefix(setting, os.Getenv("COMMDEPATH")) {
			t.Errorf("Wrong path: %s", setting)
		}
	})
	t.Run("Test GetPathrulesSetting", func(t *testing.T) {
		setting := GetPathRulesSetting()
		if setting != "" && !filepath.HasPrefix(setting, os.Getenv("COMMDEPATH")) {
			t.Errorf("Wrong path: %s", setting)
		}
	})
	t.Run("Test GetMultirulesSetting", func(t *testing.T) {
		setting := GetMultiRulesSetting()
		if setting != "" && !filepath.HasPrefix(setting, os.Getenv("COMMDEPATH")) {
			t.Errorf("Wrong path: %s", setting)
		}
	})
}

func BenchmarkGetYamlSetting(b *testing.B) {
	GetCommRulesSetting()
	GetMultiRulesSetting()
	GetPathRulesSetting()
}
