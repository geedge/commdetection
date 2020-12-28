/*
 * @Author: EnderByEndera
 * @Date: 2020-12-23 13:17:17
 * @LastEditTime: 2020-12-28 16:20:30
 * @LastEditors: Please set LastEditors
 * @Description: test router.go
 * @FilePath: /commdetection/route/router_test.go
 */

package router

import (
	"commdetection/cmd"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func init() {
	go Start()
	time.Sleep(1 * time.Second)
}

func TestStart(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1" + addr + "/version")
	if err != nil {
		t.Error(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	var jsonBody map[string]string
	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		t.Error(err)
	}
	if string(jsonBody["ver"]) != cmd.Ver {
		fmt.Println(string(body))
		t.Errorf("result is not as predicted")
	}
}
