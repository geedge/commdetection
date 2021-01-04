/*
 * @Author: your name
 * @Date: 2020-12-28 16:05:24
 * @LastEditTime: 2021-01-04 19:02:57
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /commdetection/router/endpoint_test.go
 */

package router

import (
	"commdetection/cmd"
	"commdetection/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect"
)

var (
	e          *httpexpect.Expect
	url        string
	testscomms = model.SComms{
		{
			Comm:        "sudo",
			Coefficient: 0.1,
		},
		{
			Comm:        "wget",
			Coefficient: 0.8,
		},
		{
			Comm:        "apt",
			Coefficient: 0.2,
		},
	}
	testspaths = model.SPaths{
		{
			Path:        "/root/go/src/commdetection",
			Coefficient: 0.5,
		},
	}
	testrs = model.Rules{
		{
			Name:     "pathrule",
			RuleFunc: "EvaluatePathRule",
		},
	}
)

func init() {
	testsrv := httptest.NewServer(router)
	url = testsrv.URL
}

//TODO: will finish the test later
func TestEndpoint(t *testing.T) {
	e = httpexpect.New(t, url)

	t.Run("Test verEndpoint", func(t *testing.T) {
		e.GET("/version").
			Expect().
			Status(http.StatusOK).
			JSON().Object().
			ValueEqual("ver", cmd.Ver)
	})

	t.Run("Test uploadCommEndpoint", func(t *testing.T) {
		e.POST("/evaluation/upload").
			WithMultipart().
			WithFile("commfile", "/root/.bash_history").
			Expect().
			Status(http.StatusOK).
			JSON().Object().
			ValueEqual("message", "ok")
	})

	t.Run("Test startEvaluateEndpoint", func(t *testing.T) {
		e.GET("/evaluation/start").
			Expect().
			Status(http.StatusOK).
			JSON().Object().
			ValueEqual("message", "ok")
	})

	t.Run("Test getCommrulesEndpoint", func(t *testing.T) {
		e.GET("/commrules/getrules").
			Expect().
			Status(http.StatusOK).
			JSON().Object().
			ContainsKey("message").
			ContainsKey("scomms").
			ValueEqual("message", "ok").
			Value("scomms").NotNull()
	})

	t.Run("Test setCommrulesEndpoint", func(t *testing.T) {
		e.POST("/commrules/setrules").
			WithJSON(testscomms).
			Expect().
			Status(http.StatusOK).
			JSON().Object().ValueEqual("message", "ok")
	})

	t.Run("Test getPathrulesEndpoint", func(t *testing.T) {
		e.GET("/pathrules/getrules").
			Expect().
			Status(http.StatusOK).
			JSON().Object().
			ContainsKey("spaths").
			ValueEqual("message", "ok").
			Value("spaths").NotNull()
	})

	t.Run("Test setPathrulesEndpoint", func(t *testing.T) {
		e.POST("/pathrules/setrules").
			WithJSON(testspaths).
			Expect().
			Status(http.StatusOK).
			JSON().Object().
			ValueEqual("message", "ok")
	})

	t.Run("Test getEvaluationRulesEndpoint", func(t *testing.T) {
		e.GET("/rules/getrules").
			Expect().
			Status(http.StatusOK).
			JSON().Object().
			ContainsKey("rs").
			ValueEqual("message", "ok").
			Value("rs").NotNull()
	})

	t.Run("Test setEvaluationRulesEndpoint", func(t *testing.T) {
		e.POST("/rules/setrules").
			WithJSON(testrs).
			Expect().
			Status(http.StatusOK).
			JSON().Object().
			ValueEqual("message", "ok")
	})
}

func BenchmarkEndpoint(b *testing.B) {
	e = httpexpect.New(b, url)
	b.Run("Benchmark verEndpoint", func(b *testing.B) {
		e.GET("/version").Expect()
	})

	b.Run("Benchmark uploadCommEndpoint", func(b *testing.B) {
		e.POST("/evaluation/upload").
			WithMultipart().
			WithFile("commfile", "/root/.bash_history").
			Expect()
	})

	b.Run("Benchmark startEvaluateEndpoint", func(b *testing.B) {
		e.GET("/evaluation/start").Expect()
	})

	b.Run("Benchmark getCommrulesEndpoint", func(b *testing.B) {
		e.GET("/commrules/getrules").Expect()
	})

	b.Run("Benchmark setCommrulesEndpoint", func(b *testing.B) {
		e.POST("/commrules/setrules").
			WithJSON(testscomms).Expect()
	})

	b.Run("Benchmark getPathrulesEndpoint", func(b *testing.B) {
		e.GET("/pathrules/getrules").Expect()
	})

	b.Run("Benchmark setPathrulesEndpoint", func(b *testing.B) {
		e.POST("/pathrules/setrules").
			WithJSON(testspaths).Expect()
	})

	b.Run("Benchmark getEvaluationRulesEndpoint", func(b *testing.B) {
		e.GET("/rules/getrules").Expect()
	})

	b.Run("Benchmark setEvaluationRulesEndpoint", func(b *testing.B) {
		e.POST("rules/setrules").
			WithJSON(testrs).Expect()
	})
}
