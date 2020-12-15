/*
 * @Author: EnderByEndera
 * @Date: 2020-12-09 16:44:44
 * @LastEditTime: 2020-12-15 11:52:42
 * @LastEditors: Please set LastEditors
 * @Description: Init settings from yaml file
 * @FilePath: /commdetection/init/init.go
 */

package yaml

import (
	"io/ioutil"
	"log"
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
}

// GetYamlSetting gets the yaml settings from conf.yaml
func GetYamlSetting(settingName string) (c string) {
	conf := new(Conf)
	buf, err := ioutil.ReadFile(filepath.Join(os.Getenv("COMMDEPATH"), "conf.yaml"))
	if err != nil {
		log.Println("read yaml file failed, return NULL string")
		return ""
	}
	err = yaml.Unmarshal(buf, conf)
	if err != nil {
		log.Println("Unmarshal yaml file failed, please check conf.yaml, return NULL string")
		return ""
	}
	switch settingName {
	case "commrules":
		c = filepath.Join(os.Getenv("COMMDEPATH"), conf.Paths.CommRulePath)
	case "pathrules":
		c = filepath.Join(os.Getenv("COMMDEPATH"), conf.Paths.PathRulePath)
	case "rules":
		c = filepath.Join(os.Getenv("COMMDEPATH"), conf.Paths.MultiRulesPath)
	default:
		c = os.Getenv("COMMDEPATH")
	}
	return
}
