/*
 * @Author: EnderByEndera
 * @Date: 2020-12-23 13:17:17
 * @LastEditTime: 2020-12-29 14:05:49
 * @LastEditors: Please set LastEditors
 * @Description: test router.go
 * @FilePath: /commdetection/route/router_test.go
 */

package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect"
)

func TestStart(t *testing.T) {
	srv := httptest.NewServer(router)
	eng := httpexpect.New(t, srv.URL)
	eng.GET("/version").
		Expect().
		Status(http.StatusOK)
}
