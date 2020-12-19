/*
 * @Author: EnderByEndera
 * @Date: 2020-12-16 13:57:06
 * @LastEditTime: 2020-12-19 10:48:00
 * @LastEditors: Please set LastEditors
 * @Description: Includes various marshalling ways
 * @FilePath: /commdetection/rules/marshalling.go
 */

package rules

import (
	"commdetection/logger"
	"commdetection/yaml"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)

//TODO: Temporarily use the funcMap to map the stringVal and the funcRule, will use a func to map later
var (
	ruleFuncMap = map[string]evaluation{
		"EvaluateCommandRule": EvaluateCommandRule,
		"EvaluatePathRule":    EvaluatePathRule,
	}
)

// GetFuncName returns a function's name
func GetFuncName(i interface{}, seps ...rune) string {
	fn := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	fields := strings.FieldsFunc(fn, func(sep rune) bool {
		for _, s := range seps {
			if sep == s {
				return true
			}
		}
		return false
	})
	if size := len(fields); size > 0 {
		return fields[size-1]
	}
	return ""
}

func marshalRules(rs Rules) (err error) {
	fileName := yaml.GetYamlSetting("multirules")
	if fileName == "" {
		fileName = filepath.Join(os.Getenv("COMMDEPATH"), "rules", "rulesjson", "rules.json")
	}
	err = marshalSetting(rs, fileName)
	return
}

// unmarshalRules unmarshal rules from a json file
func unmarshalRules() (r Rules, err error) {
	fileName := yaml.GetYamlSetting("rules")
	if fileName == "" {
		fileName = filepath.Join(os.Getenv("COMMDEPATH"), "rules", "rulesjson", "rules.json")
	}
	err = unmarshalSetting(fileName, r)
	if err != nil {
		fmt.Printf("json file settings conversion to %s failed, please check json file %s is correct or not", reflect.TypeOf(r).Name(), fileName)
	}
	return
}

func marshalSensitivePathSetting(spaths SPaths, fileName string) error {
	return marshalSetting(spaths, fileName)
}

func unmarshalSensitivePathSetting(fileName string) (spaths SPaths, err error) {
	err = unmarshalSetting(fileName, &spaths)
	if err != nil {
		logger.Warnf("json file settings conversion to %s failed, please check json file %s is correct or not\n", reflect.TypeOf(spaths).Name(), fileName)
	}
	return
}

func marshalSensitiveCommSetting(scomms []SComm, fileName string) error {
	return marshalSetting(scomms, fileName)
}

func unmarshalSensitiveCommSetting(fileName string) (scomms SComms, err error) {
	err = unmarshalSetting(fileName, &scomms)
	if err != nil {
		logger.Warnf("json file settings conversion to %s failed, please check json file %s is correct or not\n", reflect.TypeOf(scomms).Name(), fileName)
	}
	return
}

func marshalSetting(settings interface{}, fileName string) error {
	buf, err := json.Marshal(settings)
	if err != nil {
		return err
	}
	ioutil.WriteFile(fileName, buf, os.ModeAppend)
	if err != nil {
		return err
	}
	return nil
}

func unmarshalSetting(fileName string, settings interface{}) error {
	buf, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	err = json.Unmarshal(buf, settings)
	return err
}
