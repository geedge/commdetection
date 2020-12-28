/*
 * @Author: EnderByEndera
 * @Date: 2020-12-09 16:44:44
 * @LastEditTime: 2020-12-28 17:16:23
 * @LastEditors: Please set LastEditors
 * @Description: Init settings from yaml file
 * @FilePath: /commdetection/init/init.go
 */

package yaml

import (
	"commdetection/logger"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Conf defines the configuration from the yaml file
type Conf struct {
	Paths struct {
		MultiRulesPath string `yaml:"multirulespath"`
		CommRulePath   string `yaml:"commrulepath"`
		PathRulePath   string `yaml:"pathrulepath"`
	}
	GinMode string `yaml:"ginmode"`
}

var conf *Conf

func init() {
	conf = new(Conf)
	buf, err := ioutil.ReadFile(filepath.Join(os.Getenv("COMMDEPATH"), "conf.yaml"))
	if err != nil {
		logger.Warnln("read yaml file failed, return NULL string")
	}
	err = yaml.Unmarshal(buf, conf)
	if err != nil {
		logger.Warnln("Unmarshal yaml file failed, please check conf.yaml, return NULL string")
	}
}

// GetGinMode returns ginmode(release, debug, test)
func GetGinMode() (string, error) {
	if conf.GinMode == "debug" || conf.GinMode == "release" || conf.GinMode == "test" {
		return conf.GinMode, nil
	}
	return "release", fmt.Errorf("GinMode reads failed. Cannot find %s in any of gin modes", conf.GinMode)
}

// GetYamlPathsSetting gets the paths: settings from conf.yaml
func GetYamlPathsSetting(settingName string) (c string) {
	switch settingName {
	case "commrules":
		c = filepath.Join(os.Getenv("COMMDEPATH"), conf.Paths.CommRulePath)
	case "pathrules":
		c = filepath.Join(os.Getenv("COMMDEPATH"), conf.Paths.PathRulePath)
	case "rules":
		c = filepath.Join(os.Getenv("COMMDEPATH"), conf.Paths.MultiRulesPath)
	default:
		c = os.Getenv("COMMDEPATH")
		logger.Warnln("Didn't get any setting, use default value: " + os.Getenv("COMMDEPATH"))
	}
	return
}
