/*
 * @Author: EnderByEndera
 * @Date: 2020-12-08 10:59:19
 * @LastEditTime: 2020-12-09 11:33:58
 * @LastEditors: Please set LastEditors
 * @Description: Unmarshal and marshal various types of settings and rules
 * @FilePath: /commdetection/rules/ruletypes.go
 */

package rules

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
)

// SPath includes sensitive path dir and its sensitive coefficient
type SPath struct {
	Path        string
	Coefficient int
}

// SPaths is the SPath's multiple type
type SPaths []SPath

// SComm includes sensitive command's name and its sensitive coefficient
type SComm struct {
	Comm        string  `json:"command"`
	Coefficient float64 `json:"coefficient"`
}

// SComms is the SComm's multile types
type SComms []SComm

// EvaluatePathRule defines a rule from json file to judge command score by
// whether sensitive path appears in the Command
func (cs *CommScore) EvaluatePathRule() {
	fileName := getFileSettingsDir("pathrule")
	_, err := unmarshalSensitivePathSetting(fileName)
	if err != nil {
		log.Fatalf("cannot get sensitive paths from file %s, please check the file path", fileName)
		return
	}
	// TODO: Cannot find a way to judge whether flag is a file or not
}

// EvaluateCommandRule defines a rule from json file to judge command score by
// whether sensitive command appears in the Command
func (cs *CommScore) EvaluateCommandRule() {
	fileName := getFileSettingsDir("commrule")
	scomms, err := unmarshalSensitiveCommSetting(fileName)
	if err != nil {
		log.Fatalf("cannot get sensitive paths from file %s, please check the file path", fileName)
		return
	}
	for _, scomm := range scomms {
		if scomm.Comm == cs.Command.CommName {
			cs.Score *= scomm.Coefficient
		}
	}
}

func marshalSensitivePathSetting(spaths SPaths, fileName string) error {
	return marshalSetting(spaths, fileName)
}

func marshalSensitiveCommSetting(scomms []SComm, fileName string) error {
	return marshalSetting(scomms, fileName)
}

func unmarshalSensitivePathSetting(fileName string) (spaths SPaths, err error) {

	err = unmarshalSetting(fileName, &spaths)
	if err != nil {
		fmt.Printf("json file settings conversion to %s failed, please check json file %s is correct or not\n", reflect.TypeOf(spaths).Name(), fileName)
	}
	return
}

func unmarshalSensitiveCommSetting(fileName string) (scomms SComms, err error) {
	err = unmarshalSetting(fileName, &scomms)
	if err != nil {
		fmt.Printf("json file settings conversion to %s failed, please check json file %s is correct or not\n", reflect.TypeOf(scomms).Name(), fileName)
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
	if err != nil {
		fmt.Println("Unmarshal Settings from json file failed, please check your json file")
	}
	return err
}

// TODO: Only a stub, will promote later
func getFileSettingsDir(settingName string) string {
	if settingName == "commrule" {
		return "rulesjson/commrules.json"
	}
	if settingName == "pathrule" {
		return "rulesjson/pathrules.json"
	}
	return ""
}
