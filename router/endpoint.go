/*
 * @Author: EnderByEndera
 * @Date: 2020-12-28 15:53:37
 * @LastEditTime: 2020-12-28 17:31:34
 * @LastEditors: Please set LastEditors
 * @Description: endpoint.go contains all the endpoints
 * @FilePath: /commdetection/router/endpoint.go
 */

package router

import (
	"commdetection/cmd"
	"commdetection/logger"
	"commdetection/rules"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func verEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ver": cmd.Ver,
	})
}

// uploadCommEndpoint uploads the file storing commands to the system
func uploadCommEndpoint(c *gin.Context) {
	commfile, err := c.FormFile("commfile")
	if err != nil {
		handleErr(c, err)
		return
	}

	filename := filepath.Base(commfile.Filename)
	if err := c.SaveUploadedFile(commfile, filename); err != nil {
		handleErr(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

// startEvaluateEndpoint let the system start to evaluate commands and their args and gives the score to the user
func startEvaluateEndpoint(c *gin.Context) {
	css := cmd.StartEvaluateCommands(cmd.GetFileName())
	jsonBuf, _ := json.Marshal(css)
	logger.Debugln("Storing result to \"static/output.json\" file")
	if err := ioutil.WriteFile(filepath.Join(os.Getenv("COMMDEPATH"), "static", "output.json"), jsonBuf, os.ModeAppend); err != nil {
		handleErr(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"scores":  css,
	})
}

// getCommrulesEndpoint gets all the command rules stored in the json file
// if error happened, handleErr will be executed
func getCommrulesEndpoint(c *gin.Context) {
	scomms, err := rules.UnmarshalSensitiveCommSetting()
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
	commrules := rules.SComms{}
	if err := c.ShouldBind(&commrules); err != nil {
		handleErr(c, err)
		return
	}
	if err := rules.MarshalSensitiveCommSetting(commrules); err != nil {
		handleErr(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func getPathrulesEndpoint(c *gin.Context) {
	spaths, err := rules.UnmarshalSensitivePathSetting()
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
	pathrules := rules.SPaths{}
	if err := c.ShouldBind(&pathrules); err != nil {
		handleErr(c, err)
		return
	}
	if err := rules.MarshalSensitivePathSetting(pathrules); err != nil {
		handleErr(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func getEvaluationRulesEndpoint(c *gin.Context) {
	rs, err := rules.UnmarshalSensitivePathSetting()
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

func setEvaluationRulesEndpoint(c *gin.Context) {
	rs := rules.Rules{}
	if err := c.ShouldBind(&rs); err != nil {
		handleErr(c, err)
		return
	}
	if err := rules.MarshalEvaluationRules(rs); err != nil {
		handleErr(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

// handleErr handles all the error, set warning info in the command line, returns the error to the frontend
// and exit the endpoint
func handleErr(c *gin.Context, err error) {
	logger.Warnln(err)
	c.JSON(http.StatusOK, gin.H{
		"message": err.Error(),
	})
}
