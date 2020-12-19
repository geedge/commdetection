/*
 * @Author: EnderByEndera
 * @Date: 2020-12-08 10:59:19
 * @LastEditTime: 2020-12-19 10:48:21
 * @LastEditors: Please set LastEditors
 * @Description: Unmarshal and marshal various types of settings and rules
 * @FilePath: /commdetection/rules/ruletypes.go
 */

package rules

import (
	"commdetection/logger"
	"commdetection/yaml"
	"os"
	"path/filepath"
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
		logger.Warnf("cannot get sensitive paths from file %s, please check the file path", fileName)
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
				ratio := float64(similar) / float64(len(spath.Path))
				if ratio > 0.1 {
					cs.Score *= 1.0 - ratio*(1.0-spath.Coefficient)
				}
			}
			return nil
		})
		if err != nil {
			logger.Warnf("Error occured during EvaluatePathRule, error is %s", err)
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
		logger.Warnf("cannot get sensitive paths from file %s, please check the file path", fileName)
		return cs
	}
	for _, scomm := range scomms {
		if scomm.Comm == cs.Command.CommName {
			cs.Score *= scomm.Coefficient
		}
	}
	return cs
}
