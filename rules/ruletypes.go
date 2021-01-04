/*
 * @Author: EnderByEndera
 * @Date: 2020-12-08 10:59:19
 * @LastEditTime: 2021-01-04 18:56:54
 * @LastEditors: Please set LastEditors
 * @Description: Unmarshal and marshal various types of settings and rules
 * @FilePath: /commdetection/rules/ruletypes.go
 */

package rules

import (
	"commdetection/logger"
	"commdetection/model"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

//TODO: Temporarily use the funcMap to map the stringVal and the funcRule, will use a func to map later
var (
	RuleFuncMap = map[string]model.Evaluation{
		"EvaluateCommandRule": EvaluateCommandRule,
		"EvaluatePathRule":    EvaluatePathRule,
	}
)

// EvaluatePathRule defines a rule from json file to judge command score by
// whether sensitive path appears in the Command
func EvaluatePathRule(cs model.CommScore) model.CommScore {
	spaths, err := model.UnmarshalSensitivePathSetting()
	if err != nil {
		logger.Warnf("cannot get sensitive paths from file, please check the file path")
		return cs
	}
	for _, spath := range spaths {
		ratio := 0.0
		filepath.Walk(spath.Path, func(path string, info os.FileInfo, err error) error {
			for _, arg := range cs.Command.Args {
				if filepath.Base(arg) == info.Name() {
					ratio += float64(len(info.Name())) / float64(len(arg))
				}
			}
			return nil
		})
		// when ratio is larger than threshold 0.1
		if ratio > 0.1 {
			// execute evaluation formula
			cs.Score *= 1.0 - ratio*(1.0-spath.Coefficient)
		}
		if cs.Score < 0 {
			cs.Score = 0.0
			return cs
		}
	}
	return cs
}

// EvaluateCommandRule defines a rule from json file to judge command score by
// whether sensitive command appears in the Command
func EvaluateCommandRule(cs model.CommScore) model.CommScore {
	scomms, err := model.UnmarshalSensitiveCommSetting()
	if err != nil {
		logger.Warnln("cannot get sensitive paths from file, please check the file path")
		return cs
	}
	for _, scomm := range scomms {
		if scomm.Comm == cs.Command.CommName {
			cs.Score *= scomm.Coefficient
		}
	}
	return cs
}

// EvaluateWebsiteRule defines a rule from json file to judge commmand score
//
// 1. Check the arg of one command is a website or not
//
// 2. Check whether the website is in the whitelist, a.k.a ussites(unsensitive websites)
//
// 3. if in, the rule won't lower the score, if not, lower the score by coefficient defined in the ussites
func EvaluateWebsiteRule(cs model.CommScore) model.CommScore {
	ussites, err := model.UnMarshalUnsensitiveWebsiteSetting()
	if err != nil {
		logger.Warnln("cannot get sensitive paths from file, please check the file path")
		return cs
	}
	argsites := []*url.URL{}
	// Check the arg of one command is a website or not
	for _, arg := range cs.Command.Args {
		u, err := url.Parse(arg)
		if err == nil {
			argsites = append(argsites, u)
		}
	}

	// Start checking whether the argsites are in the websites or not
	for _, argsite := range argsites {
		// Get the host of one argsite and compare with website
		included := false
		// Check whether the argsite is in the websites
		for _, website := range ussites.Websites {
			if strings.Contains(website, argsite.Host) {
				included = true
			}
		}
		if !included {
			cs.Score *= ussites.Coefficient
		}
	}
	return cs
}
