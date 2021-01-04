/*
 * @Author: EnderByEndera
 * @Date: 2020-12-04 15:03:00
 * @LastEditTime: 2021-01-04 16:44:06
 * @LastEditors: Please set LastEditors
 * @Description: rules provide all the rules to check the commands' availability and set score of every command
 * @FilePath: /commdetection/rules/commcheck.go
 */

package rules

import (
	"commdetection/logger"
	"commdetection/model"
)

// AddRule adds one rule to the rules
func AddRule(rs model.Rules, rule model.Rule) model.Rules {
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
func DeleteRuleByName(r model.Rules, ruleName string) model.Rules {
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
func CreateRule(name string, rule string) (r model.Rule) {
	r.Name = name
	r.RuleFunc = rule
	return
}

// EvaluateCommScore evaluates scores of the commands in the css
func EvaluateCommScore(css []model.CommScore, rs model.Rules) []model.CommScore {
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
			css[index] = RuleFuncMap[r.RuleFunc](css[index])
		}
	}
	return css
}
