/*
 * @Author: EnderByEndera
 * @Date: 2021-01-16 08:12:49
 * @LastEditTime: 2021-01-16 08:16:53
 * @LastEditors: Please set LastEditors
 * @Description: rule_endpoint has the function for all the command rules' endpoint, including getter and setter
 * @FilePath: /commdetection/router/rule_endpoint.go
 */

package router

import (
	"commdetection/logger"
	"commdetection/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// getCommrulesEndpoint gets all the command rules stored in the json file
// if error happened, handleErr will be executed
func getCommrulesEndpoint(c *gin.Context) {
	scomms, err := model.UnmarshalSensitiveCommSetting()
	if err != nil {
		handleErr(c, err)
		return
	}
	logger.Debugln("Start to send scomms to the front end")
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"scomms":  scomms,
	})
}

// setCommrulesEndpoint get the command rule setting in the post form and marshal it to the json file
// Post form should be like:
//
// {
//     "commrules": [
//	      {
//            "comm": "command name",
//            "coefficient": "coefficient num",
//        }, ...
//     ]
// }
func setCommrulesEndpoint(c *gin.Context) {
	commrules := model.SComms{}
	if err := c.ShouldBind(&commrules); err != nil {
		handleErr(c, err)
		return
	}
	if err := model.MarshalSensitiveCommSetting(commrules); err != nil {
		handleErr(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func getPathrulesEndpoint(c *gin.Context) {
	spaths, err := model.UnmarshalSensitivePathSetting()
	if err != nil {
		handleErr(c, err)
		return
	}
	logger.Debugln("Start to send spaths to the front end")
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"spaths":  spaths,
	})
}

// setPathrulesEndpoint get the paths rule setting in the post form and marshal it to the json file
// Post form should be like:
//
// {
//     "pathrules": [
//	      {
//            "path": "command name",
//            "coefficient": "coefficient num",
//        }, ...
//     ]
// }
func setPathrulesEndpoint(c *gin.Context) {
	pathrules := model.SPaths{}
	if err := c.ShouldBind(&pathrules); err != nil {
		handleErr(c, err)
		return
	}
	if err := model.MarshalSensitivePathSetting(pathrules); err != nil {
		handleErr(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func getEvaluationRulesEndpoint(c *gin.Context) {
	rs, err := model.UnmarshalSensitivePathSetting()
	if err != nil {
		handleErr(c, err)
		return
	}
	logger.Debugln("Start to send evaluation rules to the front end")
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"rs":      rs,
	})
}

// setPathrulesEndpoint get the paths rule setting in the post form and marshal it to the json file
// Post form should be like:
//
// {
//     "rules": [
//	      {
//            "name": "rule name",
//            "rulefunc": "func name",
//        }, ...
//     ]
// }
func setEvaluationRulesEndpoint(c *gin.Context) {
	rs := model.Rules{}
	if err := c.ShouldBind(&rs); err != nil {
		handleErr(c, err)
		return
	}
	if err := model.MarshalEvaluationRules(rs); err != nil {
		handleErr(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
