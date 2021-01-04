/*
 * @Author: EnderByEndera
 * @Date: 2020-12-08 10:59:19
 * @LastEditTime: 2021-01-04 11:21:49
 * @LastEditors: Please set LastEditors
 * @Description: Unmarshal and marshal various types of settings and rules
 * @FilePath: /commdetection/rules/ruletypes.go
 */

package rules

import (
	"commdetection/logger"
	"io/ioutil"
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

// SComms is the SComm's multiple types
type SComms []SComm

// EvaluatePathRule defines a rule from json file to judge command score by
// whether sensitive path appears in the Command
func EvaluatePathRule(cs CommScore) CommScore {
	spaths, err := UnmarshalSensitivePathSetting()
	if err != nil {
		logger.Warnf("cannot get sensitive paths from file, please check the file path")
		return cs
	}
	for _, spath := range spaths {
		fileInfos, err := readAllFiles(spath.Path, []string{})
		if err != nil {
			logger.Warnf("Error occured during EvaluatePathRule, error is %s", err)
			break
		}
		ratio := 0.0
		in := make(chan float64)
		out := make(chan float64)
		// Creates the outside go func
		go func(spath SPath) {
			r1 := 0.0
			for _, fileInfo := range fileInfos {
				// Creates the inside go func
				go func(fileInfo string) {
					r2 := 0.0
					for _, arg := range cs.Command.Args {
						baseArg, baseSPath := filepath.Base(arg), filepath.Base(fileInfo)
						if baseArg == baseSPath {
							r2 += float64(len(arg)) / float64(len(fileInfo))
						}
					}
					in <- r2
				}(fileInfo)
			}
			for range fileInfos {
				r1 += <-in
			}
			out <- r1
		}(spath)

		for range spaths {
			ratio += <-out
		}
		// when ratio is larger than threshold 0.1
		if ratio > 0.1 {
			// execute evaluation formula
			cs.Score *= 1.0 - ratio*(1.0-spath.Coefficient)
		}
		if cs.Score < 0 {
			cs.Score = 0.0
		}
	}
	return cs
}

func readAllFiles(pathname string, s []string) ([]string, error) {
	fromSlash := filepath.FromSlash(pathname)
	rd, err := ioutil.ReadDir(fromSlash)
	if err != nil {
		logger.Warnln("read dir fail:", err)
		return s, err
	}
	for _, fi := range rd {
		if fi.IsDir() {
			fullDir := filepath.Join(fromSlash, fi.Name())
			s, err = readAllFiles(fullDir, s)
			if err != nil {
				logger.Warnln("read dir fail:", err)
				return s, err
			}
		} else {
			fullName := filepath.Join(fromSlash, fi.Name())
			s = append(s, fullName)
		}
	}
	return s, nil
}

// EvaluateCommandRule defines a rule from json file to judge command score by
// whether sensitive command appears in the Command
func EvaluateCommandRule(cs CommScore) CommScore {
	scomms, err := UnmarshalSensitiveCommSetting()
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
