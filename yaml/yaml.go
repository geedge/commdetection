/*
 * @Author: EnderByEndera
 * @Date: 2020-12-09 16:44:44
 * @LastEditTime: 2021-01-06 12:40:04
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

// MongoSet is the setting for the mongo db connection
type MongoSet struct {
	Host    string `yaml:"host"`
	Port    uint16 `yaml:"port"`
	User    string `yaml:"user"`
	Pwd     string `yaml:"password"`
	Timeout string `yaml:"timeout"`
}

// RulePaths is the rule paths
type RulePaths struct {
	MultiRulesPath   string `yaml:"multirulepath"`
	CommRulePath     string `yaml:"commrulepath"`
	PathRulePath     string `yaml:"pathrulepath"`
	WebsiteRulesPath string `yaml:"webrulepath"`
}

// GinSet is the setting for the gin-gonic router
type GinSet struct {
	GinMode string `yaml:"ginmode"`
	GinPort uint16 `yaml:"port"`
}

// Conf defines the configuration from the yaml file
type Conf struct {
	Paths RulePaths `yaml:"path"`
	Gin   GinSet    `yaml:"gin"`
	Mongo MongoSet  `yaml:"mongo"`
}

var conf *Conf

func init() {
	conf = new(Conf)
	buf, err := ioutil.ReadFile(filepath.Join(os.Getenv("COMMDEPATH"), "conf.yaml"))
	if err != nil {
		logger.Warnln("read yaml file failed, return NULL string")
		return
	}
	err = yaml.Unmarshal(buf, conf)
	if err != nil {
		logger.Warnln("Unmarshal yaml file failed, please check conf.yaml, return NULL string")
	}
}

// GetMongoSetting returns the mongo db settings
func GetMongoSetting() MongoSet {
	if conf.Mongo.Host == "" {
		conf.Mongo.Host = "localhost"
	}
	if conf.Mongo.Port == 0 {
		conf.Mongo.Port = 27017
	}
	if conf.Mongo.Timeout == "" {
		conf.Mongo.Timeout = "30s"
	}
	return conf.Mongo
}

// GetPort returns port of the server, if port is undefined, then return :8080
func GetPort() string {
	if conf.Gin.GinPort != 0 {
		return ":" + strconv.Itoa(int(conf.Gin.GinPort))
	}
	return ":8080"

}

// GetGinMode returns ginmode(release, debug, test)
func GetGinMode() (string, error) {
	if conf.Gin.GinMode == "debug" || conf.Gin.GinMode == "release" || conf.Gin.GinMode == "test" {
		return conf.Gin.GinMode, nil
	}
	return "debug", fmt.Errorf("GinMode reads failed. Cannot find %s in any of gin modes", conf.Gin.GinMode)
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
