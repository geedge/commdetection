/*
 * @Author: EnderByEndera
 * @Date: 2020-12-08 11:28:49
 * @LastEditTime: 2021-01-04 19:06:53
 * @LastEditors: Please set LastEditors
 * @Description: Test UnmarshalSetting and MarshalSetting
 * @FilePath: /commdetection/rules/rulestype_test.go
 */

package rules

import (
	"commdetection/comm"
	"commdetection/model"
	"log"
	"testing"
)

func TestEvaluateCommandRule(t *testing.T) {
	t.Run("Test Normal Command Rule Evaluation", func(t *testing.T) {
		cs := model.CommScore{
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
		cs := model.CommScore{
			Command: comm.Command{
				CommName: "wget",
				Args:     []string{"rules.json"},
			},
			Score: 100,
		}
		cs = EvaluatePathRule(cs)
		if cs.Score >= 0 && cs.Score < 100 {
			log.Printf("result score is %f", cs.Score)
		} else {
			t.Errorf("score is not as predicted")
		}
	})
}

func TestEvaluateWebsiteRule(t *testing.T) {
	t.Run("Test Evaluating Website Rule", func(t *testing.T) {
		cs := model.CommScore{
			Command: comm.Command{
				CommName: "wget",
				Args:     []string{"https://golang.org/"},
				Flags:    []string{},
			},
			Score: 100.0,
		}
		cs = EvaluateWebsiteRule(cs)
		if cs.Score == 100 {
			return
		}
		t.Errorf("Wrong score: %.2f", cs.Score)
	})
}

func BenchmarkEvaluateCommandRule(b *testing.B) {
	comms := comm.GetCommands("/root/.bash_history", "")
	comms = comm.FlushCommands(comms, []comm.Filter{comm.WhichCommandFilter})
	css := InitCommScores(comms)
	b.ResetTimer()
	for _, cs := range css {
		cs = EvaluateCommandRule(cs)
	}
}

func BenchmarkEvaluatePathRule(b *testing.B) {
	comms := comm.GetCommands("/root/.bash_history", "")
	comms = comm.FlushCommands(comms, []comm.Filter{comm.WhichCommandFilter})
	css := InitCommScores(comms)
	b.ResetTimer()
	for _, cs := range css {
		cs = EvaluatePathRule(cs)
	}
}

func BenchmarkEvaluateWebsiteRule(b *testing.B) {
	comms := comm.GetCommands("/root/.bash_history", "")
	comms = comm.FlushCommands(comms, []comm.Filter{comm.WhichCommandFilter})
	css := InitCommScores(comms)
	b.ResetTimer()
	for _, cs := range css {
		cs = EvaluateWebsiteRule(cs)
	}
}
