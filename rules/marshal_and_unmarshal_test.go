/*
 * @Author: EnderByEndera
 * @Date: 2020-12-16 14:31:00
 * @LastEditTime: 2020-12-31 14:50:23
 * @LastEditors: Please set LastEditors
 * @Description: Test marshalling.go
 * @FilePath: /commdetection/rules/marshalling_test.go
 */

package rules

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestMarshalSensitiveCommSetting(t *testing.T) {
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
	err := MarshalSensitiveCommSetting(scomms)
	if err != nil {
		t.Error(err)
	}
	jsonBuf, _ := json.Marshal(scomms)
	if jsonBuf == nil {
		t.Errorf("results are not as predicted")
	}
}

func TestUnmarshalSensitiveCommSetting(t *testing.T) {
	scomms, err := UnmarshalSensitiveCommSetting()
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
	spaths := SPaths{
		{
			Path:        "/root/go/src/commdetection",
			Coefficient: 0.7,
		},
	}
	err := MarshalSensitivePathSetting(spaths)
	if err != nil {
		t.Error(err)
	}
	jsonBuf, _ := json.Marshal(spaths)
	if jsonBuf == nil {
		t.Errorf("results are not as predicted")
	}
}

func TestUnmarshalSensitivePathSetting(t *testing.T) {
	spaths, err := UnmarshalSensitivePathSetting()
	if err != nil {
		t.Error(err)
	}
	predict := SPaths{
		{
			Path:        "/root/go/src/commdetection",
			Coefficient: 0.7,
		},
	}
	if !reflect.DeepEqual(spaths, predict) {
		t.Errorf("results are not as predicted")
	}
}
