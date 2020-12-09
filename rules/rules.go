/*
 * @Author: EnderByEndera
 * @Date: 2020-12-04 15:03:00
 * @LastEditTime: 2020-12-09 11:34:12
 * @LastEditors: Please set LastEditors
 * @Description: rules provide all the rules to check the commands' availability and set score of every command
 * @FilePath: /commdetection/rules/commcheck.go
 */

package rules

import (
	prep "commdetection/preprocessing"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

// CommScore includes command name and its score
type CommScore struct {
	Command prep.Command `json:"command"`
	Score   float64      `json:"score"`
}

// Rule defines a rule's func and its name
type Rule struct {
	name string
	rule func(CommScore) CommScore
}

// Rules is the slice of Rule
type Rules []Rule

// AddRule adds one rule to the rules
func AddRule(r Rules, rule Rule) Rules {
	r = append(r, rule)
	return r
}

// DeleteRuleByName deletes one rule by the name of the rule
func DeleteRuleByName(r Rules, ruleName string) Rules {
	for index, rule := range r {
		if rule.name == ruleName {
			r = append(r[:index], r[index+1:]...)
		}
	}
	return r
}

// MarshalRules marshal rules to a json file
// TODO: Cannot use json to store func, need to find another way to store it
func MarshalRules(r Rules, fileName string) (err error) {
	fp, err := os.Open(fileName)
	if err != nil {
		return
	}
	buf, err := json.Marshal(r)
	if err != nil {
		return
	}
	_, err = fp.Write(buf)
	if err != nil {
		return
	}
	return
}

// UnmarshalRules unmarshal rules from a json file
func UnmarshalRules(fileName string) (r Rules, err error) {
	err = unmarshalSetting(fileName, r)
	if err != nil {
		fmt.Printf("json file settings conversion to %s failed, please check json file %s is correct or not", reflect.TypeOf(r).Name(), fileName)
	}
	return
}

// CreateRule creates one rule with name and rule func and return a Rule
func CreateRule(name string, rule func(CommScore) CommScore) (r Rule) {
	r.name = name
	r.rule = rule
	return
}
