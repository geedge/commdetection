/*
 * @Author: EnderByEndera
 * @Date: 2021-01-04 16:30:53
 * @LastEditTime: 2021-01-04 18:38:56
 * @LastEditors: Please set LastEditors
 * @Description: This is the model file used for rules pack
 * @FilePath: /commdetection/model/rule_model.go
 */

package model

import (
	"commdetection/comm"
)

// Rule defines a rule's func and its name
type Rule struct {
	Name     string `json:"name"`
	RuleFunc string `json:"rulefunc"`
}

// Rules is the slice of Rule
type Rules []Rule

// Evaluation is the template of the Evaluation function
type Evaluation func(CommScore) CommScore

// CommScore includes command name and its score
type CommScore struct {
	Command comm.Command `json:"command"`
	Score   float64      `json:"score"`
}

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

// Ussites is the Unsensitive websites, a.k.a website whitelists
// Coefficient defines the intent of lowering the score if the website arg is not in the ussites.Websites
type Ussites struct {
	Websites    []string `json:"websites"`
	Coefficient float64  `json:"coefficient"`
}
