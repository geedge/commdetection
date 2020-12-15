/*
 * @Author: EnderByEndera
 * @Date: 2020-12-08 11:28:49
 * @LastEditTime: 2020-12-15 17:36:38
 * @LastEditors: Please set LastEditors
 * @Description: Test UnmarshalSetting and MarshalSetting
 * @FilePath: /commdetection/rules/rulestype_test.go
 */

package rules

import (
	prep "commdetection/preprocessing"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestMarshalSensitiveCommSetting(t *testing.T) {
	fileName := "rulesjson/commrules.json"
	scomms := SComms{
		{
			Comm:        "wget",
			Coefficient: 0.8,
		},
		{
			Comm:        "apt",
			Coefficient: 1.0,
		},
	}
	err := marshalSensitiveCommSetting(scomms, fileName)
	if err != nil {
		t.Error(err)
	}
	jsonBuf, _ := json.Marshal(scomms)
	readBuf, err := ioutil.ReadFile(fileName)
	if !reflect.DeepEqual(readBuf, jsonBuf) {
		fmt.Println(readBuf)
		t.Errorf("results are not as predicted")
	}
}

func TestUnmarshalSensitiveCommSetting(t *testing.T) {
	fileName := "rulesjson/commrules.json"
	scomms, err := unmarshalSensitiveCommSetting(fileName)
	if err != nil {
		t.Error(err)
	}
	predict := SComms{
		{
			Comm:        "wget",
			Coefficient: 0.8,
		},
		{
			Comm:        "apt",
			Coefficient: 1.0,
		},
	}
	if !reflect.DeepEqual(scomms, predict) {
		t.Errorf("results are not as predicted")
	}
}

func TestMarshalSensitivePathSetting(t *testing.T) {
	fileName := "rulesjson/pathrules.json"
	spaths := SPaths{
		{
			Path:        "/root/go/src/commdetection/rules",
			Coefficient: 0.7,
		},
	}
	err := marshalSensitivePathSetting(spaths, fileName)
	if err != nil {
		t.Error(err)
	}
	jsonBuf, _ := json.Marshal(spaths)
	readBuf, err := ioutil.ReadFile(fileName)
	if !reflect.DeepEqual(readBuf, jsonBuf) {
		fmt.Println(readBuf)
		t.Errorf("results are not as predicted")
	}
}

func TestUnmarshalSensitivePathSetting(t *testing.T) {
	fileName := "rulesjson/pathrules.json"
	spaths, err := unmarshalSensitivePathSetting(fileName)
	if err != nil {
		t.Error(err)
	}
	predict := SPaths{
		{
			Path:        "/root/go/src/commdetection/rules",
			Coefficient: 0.7,
		},
	}
	if !reflect.DeepEqual(spaths, predict) {
		t.Errorf("results are not as predicted")
	}
}

func TestEvaluateCommandRule(t *testing.T) {
	t.Run("Test Normal Command Rule Evaluation", func(t *testing.T) {
		cs := CommScore{
			Command: prep.Command{
				CommName: "wget",
				Flags:    []string{"https://127.0.0.1:8080"},
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
			Command: prep.Command{
				CommName: "wget",
				Flags:    []string{},
			},
			Score: 100,
		}
		cs = EvaluatePathRule(cs)
		if cs.Score >= 0 {
			fmt.Printf("result score is %f", cs.Score)
		} else {
			t.Errorf("score is not as predicted")
		}
	})
}

func BenchmarkEvaluateCommandRule(b *testing.B) {
	cs := CommScore{
		Command: prep.Command{
			CommName: "wget",
			Flags:    []string{"https://127.0.0.1:8080"},
		},
		Score: 100,
	}
	cs = EvaluateCommandRule(cs)
	cs = EvaluateCommandRule(cs)
}

func BenchmarkEvaluatePathRule(b *testing.B) {
	//TODO: will finish after the EvaluatePathRule is finished
}
