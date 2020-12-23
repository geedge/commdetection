/*
 * @Author: EnderByEndera
 * @Date: 2020-12-08 11:28:49
 * @LastEditTime: 2020-12-19 10:45:05
 * @LastEditors: Please set LastEditors
 * @Description: Test UnmarshalSetting and MarshalSetting
 * @FilePath: /commdetection/rules/rulestype_test.go
 */

package rules

import (
	"commdetection/comm"
	"log"
	"testing"
)

func TestEvaluateCommandRule(t *testing.T) {
	t.Run("Test Normal Command Rule Evaluation", func(t *testing.T) {
		cs := CommScore{
			Command: comm.Command{
				CommName: "wget",
				Args:     []string{"https://127.0.0.1:8080"},
			},
			Score: 100,
		}
		cs = EvaluateCommandRule(cs)
		if cs.Score != 80 {
			t.Errorf("score is not as predicted")
		}
	})
}

func TestEvaluatePathRule(t *testing.T) {
	t.Run("Test Normal Path Rule Evaluation", func(t *testing.T) {
		cs := CommScore{
			Command: comm.Command{
				CommName: "wget",
				Args:     []string{"/root/go/src/commdetection/rules/rulesjson/rules.json"},
			},
			Score: 100,
		}
		cs = EvaluatePathRule(cs)
		if cs.Score >= 0 {
			log.Printf("result score is %f", cs.Score)
		} else {
			t.Errorf("score is not as predicted")
		}
	})
}

func BenchmarkEvaluateCommandRule(b *testing.B) {
	cs := CommScore{
		Command: comm.Command{
			CommName: "wget",
			Args:     []string{"https://127.0.0.1:8080"},
		},
		Score: 100,
	}
	cs = EvaluateCommandRule(cs)
}

func BenchmarkEvaluatePathRule(b *testing.B) {
	comms := comm.GetCommands("/root/.bash_history", "")
	comms = comm.FlushCommands(comms, []comm.Filter{comm.ManCommandFilter})
	css := InitCommScores(comms)
	b.ResetTimer()
	for _, cs := range css {
		cs = EvaluatePathRule(cs)
	}
}
