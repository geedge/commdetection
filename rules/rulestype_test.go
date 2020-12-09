/*
 * @Author: EnderByEndera
 * @Date: 2020-12-08 11:28:49
 * @LastEditTime: 2020-12-09 11:35:31
 * @LastEditors: Please set LastEditors
 * @Description: Test UnmarshalSetting and MarshalSetting
 * @FilePath: /commdetection/rules/rulestype_test.go
 */

package rules

import (
	prep "commdetection/preprocessing"
	"encoding/json"
	"fmt"
	"os"
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
	fp, err := os.Open(fileName)
	if err != nil {
		t.Error(err)
	}

	jsonBuf, _ := json.Marshal(scomms)
	readBuf := make([]byte, len(jsonBuf))
	_, err = fp.Read(readBuf)
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

func TestEvaluateCommandRule(t *testing.T) {
	t.Run("Test Normal Command Rule Evaluation", func(t *testing.T) {
		cs := CommScore{
			Command: prep.Command{
				CommName: "wget",
				Flags:    []string{"https://127.0.0.1:8080"},
			},
			Score: 100,
		}
		cs.EvaluateCommandRule()
		if cs.Score != 80 {
			t.Errorf("score is not as predicted")
		}
	})
}

func TestEvaluatePathRule(t *testing.T) {

}
