/*
 * @Author: EnderByEndera
 * @Date: 2020-12-04 15:03:00
 * @LastEditTime: 2020-12-29 14:29:31
 * @LastEditors: Please set LastEditors
 * @Description: rules provide all the rules to check the commands' availability and set score of every command
 * @FilePath: /commdetection/rules/commcheck.go
 */

package rules

import (
	"commdetection/logger"
)

// Rule defines a rule's func and its name
type Rule struct {
	Name     string `json:"name"`
	RuleFunc string `json:"rulefunc"`
}

// Rules is the slice of Rule
type Rules []Rule

type evaluation func(CommScore) CommScore

// AddRule adds one rule to the rules
func AddRule(rs Rules, rule Rule) Rules {
	for _, r := range rs {
		if r.Name == rule.Name {
			logger.Warnf("%s already existed", rule.Name)
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
		logger.Warnf("rule %s not existed in the rules", ruleName)
	}
	return r
}

// CreateRule creates one rule with name and rule func and return a Rule
func CreateRule(name string, rule string) (r Rule) {
	r.Name = name
	r.RuleFunc = rule
	return
}

// EvaluateCommScore evaluates scores of the commands in the css
func EvaluateCommScore(css []CommScore, rs Rules) []CommScore {
	defer func() {
		err := recover()
		if err != nil {
			logger.Fatalln("Unexpected error happened, RuleFunc conversion failed, please check the ruleFuncMap")
		}
	}()
	if len(rs) == 0 {
		logger.Warnln("There are no rules in the evaluation")
		return css
	}
	for index := 0; index < len(css); index++ {
		for _, r := range rs {
			css[index] = ruleFuncMap[r.RuleFunc](css[index])
		}
	}
	return css
}
