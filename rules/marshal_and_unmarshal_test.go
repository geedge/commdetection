/*
 * @Author: EnderByEndera
 * @Date: 2020-12-16 14:31:00
 * @LastEditTime: 2020-12-16 14:32:44
 * @LastEditors: Please set LastEditors
 * @Description: Test marshalling.go
 * @FilePath: /commdetection/rules/marshalling_test.go
 */

package rules

import (
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
