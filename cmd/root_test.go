/*
 * @Author: your name
 * @Date: 2020-12-22 14:13:40
 * @LastEditTime: 2020-12-23 14:50:02
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /commdetection/cmd/root_test.go
 */

package cmd

import (
	"commdetection/rules"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func init() {
	go root()
	time.Sleep(2 * time.Second)
}

func BenchmarkRoot(b *testing.B) {
	b.ResetTimer()
	if res, err := http.Get("http://127.0.0.1:8080/scores"); err == nil {
		css := []rules.CommScore{}
		jsonBuf, _ := ioutil.ReadAll(res.Body)
		if err := json.Unmarshal(jsonBuf, &css); err != nil {
			b.Error(err)
		}
	} else {
		b.Error(err)
	}
}
