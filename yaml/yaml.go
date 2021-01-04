/*
 * @Author: EnderByEndera
 * @Date: 2020-12-09 16:44:44
 * @LastEditTime: 2021-01-04 19:11:44
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
	"strconv"

	"gopkg.in/yaml.v2"
)

// Conf defines the configuration from the yaml file
type Conf struct {
	Paths struct {
		MultiRulesPath   string `yaml:"multirulepath"`
		CommRulePath     string `yaml:"commrulepath"`
		PathRulePath     string `yaml:"pathrulepath"`
		WebsiteRulesPath string `yaml:"webrulepath"`
	} `yaml:"path"`
	GinMode string `yaml:"ginmode"`
	Port    uint16 `yaml:"port"`
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

// GetPort returns port of the server, if port is undefined, then return :8080
func GetPort() string {
	if conf.Port != 0 {
		return ":" + strconv.Itoa(int(conf.Port))
	}
	return ":8080"

}

// GetGinMode returns ginmode(release, debug, test)
func GetGinMode() (string, error) {
	if conf.GinMode == "debug" || conf.GinMode == "release" || conf.GinMode == "test" {
		return conf.GinMode, nil
	}
	return "debug", fmt.Errorf("GinMode reads failed. Cannot find %s in any of gin modes", conf.GinMode)
}

// GetCommRulesSetting returns the path storing sensitive command rules
func GetCommRulesSetting() string {
	if conf.Paths.CommRulePath == "" {
		return conf.Paths.CommRulePath
	}
	return filepath.Join(os.Getenv("COMMDEPATH"), conf.Paths.CommRulePath)
}

// GetPathRulesSetting returns the path storing sensitive path rules
func GetPathRulesSetting() string {
	if conf.Paths.CommRulePath == "" {
		return conf.Paths.PathRulePath
	}
	return filepath.Join(os.Getenv("COMMDEPATH"), conf.Paths.PathRulePath)
}

// GetMultiRulesSetting returns the path storing rules of the sensitive rules like command rules or path rules
func GetMultiRulesSetting() string {
	if conf.Paths.MultiRulesPath == "" {
		return conf.Paths.MultiRulesPath
	}
	return filepath.Join(os.Getenv("COMMDEPATH"), conf.Paths.MultiRulesPath)
}

// GetWebsiteRulesSetting returns the file path storing the net rules white list
func GetWebsiteRulesSetting() string {
	if conf.Paths.WebsiteRulesPath == "" {
		return conf.Paths.WebsiteRulesPath
	}
	return filepath.Join(os.Getenv("COMMDEPATH"), conf.Paths.WebsiteRulesPath)
}
