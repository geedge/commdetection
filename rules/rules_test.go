/*
 * @Author: EnderByEndera
 * @Date: 2020-12-04 15:03:09
 * @LastEditTime: 2020-12-15 15:57:32
 * @LastEditors: Please set LastEditors
 * @Description: Test commrules.go
 * @FilePath: /commdetection/rules/commrules_test.go
 */
package rules

import (
	"commdetection/preprocessing"
	"fmt"
	"testing"
)

func TestRule(t *testing.T) {
	t.Run("Test AddRule", func(t *testing.T) {
		r := Rules{
			{
				Name:     "RuleA",
				RuleFunc: "EvaluateCommandRule",
			},
			{
				Name:     "RuleB",
				RuleFunc: "EvaluatePathRule",
			},
		}
		r = AddRule(r, Rule{
			Name:     "RuleC",
			RuleFunc: "EvaluateCommandRule",
		})
		predict := Rules{
			{
				Name:     "RuleA",
				RuleFunc: "EvaluateCommandRule",
			},
			{
				Name:     "RuleB",
				RuleFunc: "EvaluatePathRule",
			},
			{
				Name:     "RuleC",
				RuleFunc: "EvaluateCommandRule",
			},
		}
		if len(r) != len(predict) {
			fmt.Println(r)
			t.Errorf("results are not as predicted")
			return
		}
		for i := range r {
			if r[i].Name != predict[i].Name {
				fmt.Println(r)
				t.Errorf("results are not as predicted")
				return
			}
		}
		cs := CommScore{
			Command: preprocessing.Command{CommName: "wget"},
			Score:   100,
		}
		for _, rule := range r {
			cs = ruleFuncMap[rule.RuleFunc](cs)
		}
		if cs.Score != 100*0.8*0.8 {
			t.Errorf("results are not as predicted")
		}
	})

	t.Run("Test AddRule with the same name", func(t *testing.T) {
		r := Rules{
			{
				Name:     "RuleA",
				RuleFunc: "EvaluateCommandRule",
			},
			{
				Name:     "RuleB",
				RuleFunc: "EvaluatePathRule",
			},
		}
		r = AddRule(r, Rule{
			Name:     "RuleA",
			RuleFunc: "EvaluateCommandRule",
		})
		if len(r) == 3 {
			t.Errorf("results are not as predicted")
		}
	})

	t.Run("Test DeleteRuleByName", func(t *testing.T) {
		r := Rules{
			{
				Name:     "RuleA",
				RuleFunc: "EvaluateCommandRule",
			},
			{
				Name:     "RuleB",
				RuleFunc: "EvaluatePathRule",
			},
		}
		r = DeleteRuleByName(r, "RuleB")
		if len(r) != 1 {
			fmt.Println(r)
			t.Errorf("results are not as predicted")
		}
		r = DeleteRuleByName(r, "RuleB")
		if len(r) != 1 {
			fmt.Println(r)
			t.Errorf("results are not as predicted")
		}
	})

	t.Run("Test MarshallRules", func(t *testing.T) {
		rs := Rules{
			{
				Name:     "RuleA",
				RuleFunc: "EvaluateCommandRule",
			},
		}
		err := MarshalRules(rs)
		if err != nil {
			t.Error(err)
		}
	})
}
