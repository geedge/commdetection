/*
 * @Author: EnderByEndera
 * @Date: 2020-12-19 11:59:02
 * @LastEditTime: 2020-12-22 17:21:05
 * @LastEditors: Please set LastEditors
 * @Description: root of the commdetection cmd
 * @FilePath: /commdetection/cmd/root.go
 */

package cmd

import (
	"commdetection/comm"
	"commdetection/logger"
	"commdetection/rules"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "commdetection",
	Short: "commdetection is the command-line tool for detecting harmful commands",
	Long:  "A fast and precise command-line tool for detecting harmful commands",
	Run: func(cmd *cobra.Command, args []string) {
		root()
	},
}

var (
	filterStr   string
	evaluations []string
	file        string
	logLevel    uint32
	output      string

	filters = []comm.Filter{}
	rs      = rules.Rules{}
	buf     []byte
	jsonBuf []byte
	r       *gin.Engine
	srv     *http.Server
)

// Execute executes the command
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// initilize gin-gonic backend engine
	r = gin.New()
	r.Use(logger.Middleware())
	r.Use(gin.Recovery())

	srv = &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	rootCmd.PersistentFlags().StringVar(&filterStr, "filter", "which", "choose one filter to filt the data from the file")
	rootCmd.PersistentFlags().StringSliceVarP(&evaluations, "evaluations", "e", []string{"command"}, "choose one or more evaluations to evaluate commands")
	rootCmd.PersistentFlags().StringVar(&file, "file", "/root/.bash_history", "choose one file storing data of the commands")
	rootCmd.PersistentFlags().Uint32Var(&logLevel, "loglevel", uint32(logrus.DebugLevel), "choose log level")
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "file", "choose the output method: json(json) or file.json(file), choose json will open a server to provide json output on the net")
	rootCmd.AddCommand(verCmd) // add version as a subcommand
}

// GetLogLevel returns loglevel sent by rootCmd
func GetLogLevel() uint32 {
	return logLevel
}

func root() {
	for _, ev := range evaluations {
		switch ev {
		case "command":
			rs = rules.AddRule(rs, rules.Rule{
				Name:     "Command",
				RuleFunc: "EvaluateCommandRule",
			})
		case "path":
			rs = rules.AddRule(rs, rules.Rule{
				Name:     "Path",
				RuleFunc: "EvaluatePathRule",
			})
		default:
			logger.Warnln("invalid rule name: " + ev)
		}
	}

	switch filterStr {
	case "which":
		filters = append(filters, comm.WhichCommandFilter)
	case "simple":
		filters = append(filters, comm.SimpleCommandFilter)
	case "help":
		filters = append(filters, comm.HelpCommandFilter)
	case "man":
		filters = append(filters, comm.ManCommandFilter)
	}

	// 从文件中获取路径，默认获取路径为/root/.bash_history
	logger.Debugln("Start getting commmands from " + file)
	commands := comm.GetCommands(file, "")
	// 清理无效命令，利用filter函数保留有效命令以便提供分析
	logger.Debugln("Start flushing commands using", filterStr)
	commands = comm.FlushCommands(commands, filters)
	// 初始化命令得分
	logger.Debugln("Initializing commands' scores")
	css := rules.InitCommScores(commands)
	// 评估命令，利用rs中保留的规则进行评估
	logger.Debugln("Evaluating commands' scores using ", evaluations)
	css = rules.EvaluateCommScore(css, rs)

	jsonBuf, _ = json.Marshal(css)
	if output == "file" {
		// 将命令得分保存到json文件中
		logger.Debugln("Storing result to output.json file")
		ioutil.WriteFile(filepath.Join(os.Getenv("COMMDEPATH"), "output.json"), jsonBuf, os.ModeAppend)
	} else if output == "json" {
		// router "/scores"
		r.GET("/scores", func(c *gin.Context) {
			c.JSON(200, css)
		})
	}
	go func() {
		// engine start on http://127.0.0.1:8000
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalln(err)
		}
	}()

	// Make a quit channel to receive SIGINT and SIGTERM signal
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Debugln("Server shutting down...")

	if err := srv.Shutdown(context.Background()); err != nil {
		logger.Fatalln("Server Shutdown with err: ", err)
	}
	logger.Debugln("Server shutted down")
}
