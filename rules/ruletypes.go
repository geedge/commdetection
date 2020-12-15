/*
 * @Author: EnderByEndera
 * @Date: 2020-12-08 10:59:19
 * @LastEditTime: 2020-12-15 17:30:21
 * @LastEditors: Please set LastEditors
 * @Description: Unmarshal and marshal various types of settings and rules
 * @FilePath: /commdetection/rules/ruletypes.go
 */

package rules

import (
	"commdetection/yaml"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
)

// SPath includes sensitive path dir and its sensitive coefficient
type SPath struct {
	Path        string
	Coefficient float64
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
func EvaluatePathRule(cs CommScore) CommScore {
	fileName := yaml.GetYamlSetting("pathrules")
	if fileName == "" {
		fileName = filepath.Join(os.Getenv("COMMDEPATH"), "rules", "rulesjson", "pathrules.json")
	}
	spaths, err := unmarshalSensitivePathSetting(fileName)
	if err != nil {
		log.Printf("cannot get sensitive paths from file %s, please check the file path", fileName)
		return cs
	}
	for _, spath := range spaths {
		err := filepath.Walk(spath.Path, func(path string, info os.FileInfo, err error) error {
			if info == nil {
				return err
			}
			for _, flag := range cs.Command.Flags {
				similar := 0
				indexPath, indexFlag := len(path)-1, len(flag)-1
				for indexPath >= 0 && indexFlag >= 0 {
					if path[indexPath] == flag[indexFlag] {
						similar++
						indexPath--
						indexFlag--
					} else {
						break
					}
				}
				cs.Score *= 1.0 - float64(similar)/float64(len(spath.Path))*(1.0-spath.Coefficient)
			}
			return nil
		})
		if err != nil {
			log.Print(err)
		}
	}
	return cs
}

// EvaluateCommandRule defines a rule from json file to judge command score by
// whether sensitive command appears in the Command
func EvaluateCommandRule(cs CommScore) CommScore {
	fileName := yaml.GetYamlSetting("commrules")
	if fileName == "" {
		fileName = filepath.Join(os.Getenv("COMMDEPATH"), "rules", "rulesjson", "commrules.json")
	}
	scomms, err := unmarshalSensitiveCommSetting(fileName)
	if err != nil {
		log.Printf("cannot get sensitive paths from file %s, please check the file path", fileName)
		return cs
	}
	for _, scomm := range scomms {
		if scomm.Comm == cs.Command.CommName {
			cs.Score *= scomm.Coefficient
		}
	}
	return cs
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
		log.Printf("json file settings conversion to %s failed, please check json file %s is correct or not\n", reflect.TypeOf(spaths).Name(), fileName)
	}
	return
}

func unmarshalSensitiveCommSetting(fileName string) (scomms SComms, err error) {
	err = unmarshalSetting(fileName, &scomms)
	if err != nil {
		log.Printf("json file settings conversion to %s failed, please check json file %s is correct or not\n", reflect.TypeOf(scomms).Name(), fileName)
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
	dir, err := os.Getwd()
	fmt.Println(dir)
	buf, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	err = json.Unmarshal(buf, settings)
	if err != nil {
		log.Println("Unmarshal Settings from json file failed, please check your json file")
	}
	return err
}
