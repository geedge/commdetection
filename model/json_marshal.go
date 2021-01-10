/*
 * @Author: your name
 * @Date: 2021-01-05 10:37:28
 * @LastEditTime: 2021-01-10 09:39:58
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /commdetection/model/marshal_and_unmarshal.go
 */
/*
 * @Author: EnderByEndera
 * @Date: 2020-12-16 13:57:06
 * @LastEditTime: 2021-01-06 09:56:56
 * @LastEditors: Please set LastEditors
 * @Description: Includes various marshalling ways
 * @FilePath: /commdetection/rules/marshalling.go
 */

package model

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

// getFuncName returns a function's name
func getFuncName(i interface{}, seps ...rune) string {
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

// MarshalEvaluationRules marshals rules.Rules to json file
func MarshalEvaluationRules(rs Rules) (err error) {
	fileName := yaml.GetMultiRulesSetting()
	if fileName == "" {
		fileName = filepath.Join(os.Getenv("COMMDEPATH"), "static", "rules", "rules.json")
	}
	err = marshalSetting(fileName, rs)
	return
}

// UnmarshalEvaluationRules unmarshal rules from a json file
func UnmarshalEvaluationRules() (r Rules, err error) {
	fileName := yaml.GetMultiRulesSetting()
	if fileName == "" {
		fileName = filepath.Join(os.Getenv("COMMDEPATH"), "static", "rules", "rules.json")
	}
	err = unmarshalSetting(fileName, r)
	if err != nil {
		fmt.Printf("json file settings conversion to %s failed, please check json file %s is correct or not", reflect.TypeOf(r).Name(), fileName)
	}
	return
}

// MarshalSensitivePathSetting marshals spaths setting to a json file
func MarshalSensitivePathSetting(spaths SPaths) error {
	fileName := yaml.GetPathRulesSetting()
	if fileName == "" {
		fileName = filepath.Join(os.Getenv("COMMDEPATH"), "static", "rules", "pathrules.json")
	}
	return marshalSetting(fileName, spaths)
}

// UnmarshalSensitivePathSetting unmarshal spaths setting json file to spaths variable
func UnmarshalSensitivePathSetting() (spaths SPaths, err error) {
	fileName := yaml.GetPathRulesSetting()
	if fileName == "" {
		fileName = filepath.Join(os.Getenv("COMMDEPATH"), "static", "rules", "pathrules.json")
	}
	err = unmarshalSetting(fileName, &spaths)
	if err != nil {
		logger.Warnf("json file settings conversion to %s failed, please check json file %s is correct or not\n", reflect.TypeOf(spaths).Name(), fileName)
	}
	return
}

// MarshalSensitiveCommSetting marshals scomms setting to a json file
func MarshalSensitiveCommSetting(scomms []SComm) error {
	fileName := yaml.GetCommRulesSetting()
	if fileName == "" {
		fileName = filepath.Join(os.Getenv("COMMDEPATH"), "static", "rules", "commrules.json")
	}
	return marshalSetting(fileName, scomms)
}

// UnmarshalSensitiveCommSetting unmarshal scomms setting json file to scomms variable
func UnmarshalSensitiveCommSetting() (scomms SComms, err error) {
	fileName := yaml.GetCommRulesSetting()
	if fileName == "" {
		fileName = filepath.Join(os.Getenv("COMMDEPATH"), "static", "rules", "commrules.json")
	}
	err = unmarshalSetting(fileName, &scomms)
	if err != nil {
		logger.Warnf("json file settings conversion to %s failed, please check json file %s is correct or not\n", reflect.TypeOf(scomms).Name(), fileName)
	}
	return
}

// MarshalUnsensitiveWebsiteSetting marshals ussites setting to a json file
func MarshalUnsensitiveWebsiteSetting(ussites Ussites) error {
	fileName := yaml.GetWebsiteRulesSetting()
	if fileName == "" {
		fileName = filepath.Join(os.Getenv("COMMDEPATH"), "static", "rules", "websiterules.json")
	}
	return marshalSetting(fileName, ussites)
}

// UnMarshalUnsensitiveWebsiteSetting unmarshal ussites setting json file to ussites variable
func UnMarshalUnsensitiveWebsiteSetting() (ussites Ussites, err error) {
	fileName := yaml.GetWebsiteRulesSetting()
	if fileName == "" {
		fileName = filepath.Join(os.Getenv("COMMDEPATH"), "static", "rules", "websiterules.json")
	}
	err = unmarshalSetting(fileName, &ussites)
	if err != nil {
		logger.Warnf("json file settings conversion to %s failed, please check json file %s is correct or not\n", reflect.TypeOf(ussites).Name(), fileName)
	}
	return
}

func marshalSetting(fileName string, settings interface{}) error {
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
