/*
 * @Author: EnderByEndera
 * @Date: 2020-12-04 15:03:00
 * @LastEditTime: 2020-12-15 10:05:33
 * @LastEditors: Please set LastEditors
 * @Description: rules provide all the rules to check the commands' availability and set score of every command
 * @FilePath: /commdetection/rules/commcheck.go
 */

package rules

import (
	prep "commdetection/preprocessing"
	"commdetection/yaml"
	"fmt"
	"log"
	"reflect"
)

// CommScore includes command name and its score
type CommScore struct {
	Command prep.Command `json:"command"`
	Score   float64      `json:"score"`
}

// Rule defines a rule's func and its name
type Rule struct {
	Name     string `json:"name"`
	RuleFunc string `json:"rulefunc"`
}

// Rules is the slice of Rule
type Rules []Rule

//TODO: Temporarily use the funcMap to map the stringVal and the funcRule, will use a func to map later
var (
	ruleFuncMap = map[string]func(CommScore) CommScore{
		"EvaluateCommandRule": EvaluateCommandRule,
		"EvaluatePathRule":    EvaluatePathRule,
	}
)

// AddRule adds one rule to the rules
func AddRule(rs Rules, rule Rule) Rules {
	for _, r := range rs {
		if r.Name == rule.Name {
			log.Printf("%s already existed", rule.Name)
			return rs
		}
	}
	rs = append(rs, rule)
	return rs
}

// DeleteRuleByName deletes one rule by the name of the rule
func DeleteRuleByName(r Rules, ruleName string) Rules {
	existed := false
	for index, rule := range r {
		if rule.Name == ruleName {
			r = append(r[:index], r[index+1:]...)
			existed = true
		}
	}
	if !existed {
		log.Printf("rule %s not existed in the rules", ruleName)
	}
	return r
}

// MarshalRules marshal rules to a json file
func MarshalRules(rs Rules) (err error) {
	fileName := "rulesjson/rules.json"
	err = marshalSetting(rs, fileName)
	return
}

// UnmarshalRules unmarshal rules from a json file
func UnmarshalRules() (r Rules, err error) {
	fileName := yaml.GetYamlSetting("rules")
	err = unmarshalSetting(fileName, r)
	if err != nil {
		fmt.Printf("json file settings conversion to %s failed, please check json file %s is correct or not", reflect.TypeOf(r).Name(), fileName)
	}
	return
}

// CreateRule creates one rule with name and rule func and return a Rule
func CreateRule(name string, rule string) (r Rule) {
	r.Name = name
	r.RuleFunc = rule
	return
}
