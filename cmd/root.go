/*
 * @Author: EnderByEndera
 * @Date: 2020-12-19 11:59:02
 * @LastEditTime: 2020-12-23 14:49:55
 * @LastEditors: Please set LastEditors
 * @Description: root of the commdetection cmd
 * @FilePath: /commdetection/cmd/root.go
 */

package cmd

import (
	"commdetection/comm"
	"commdetection/logger"
	"commdetection/router"
	"commdetection/rules"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

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

	filters = []comm.Filter{}
	rs      = rules.Rules{}
	buf     []byte
	jsonBuf []byte
)

// Execute executes the command
func Execute() error {
	return rootCmd.Execute()
}

func init() {

	rootCmd.PersistentFlags().StringVar(&filterStr, "filter", "which", "choose one filter to filt the data from the file")
	rootCmd.PersistentFlags().StringSliceVarP(&evaluations, "evaluations", "e", []string{"command"}, "choose one or more evaluations to evaluate commands")
	rootCmd.PersistentFlags().StringVar(&file, "file", "/root/.bash_history", "choose one file storing data of the commands")
	rootCmd.PersistentFlags().Uint32Var(&logLevel, "loglevel", uint32(logrus.DebugLevel), "choose log level")
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
	logger.Debugf("Start getting commmands from \"%s\"", file)
	commands := comm.GetCommands(file, "")
	// 清理无效命令，利用filter函数保留有效命令以便提供分析
	logger.Debugf("Start flushing commands using \"%s\" filter", filterStr)
	commands = comm.FlushCommands(commands, filters)
	// 初始化命令得分
	logger.Debugf("Initializing commands' scores")
	css := rules.InitCommScores(commands)
	// 评估命令，利用rs中保留的规则进行评估
	logger.Debugf("Evaluating commands' scores using %s evaluations", evaluations)
	css = rules.EvaluateCommScore(css, rs)
	// 将命令得分保存到json文件中
	jsonBuf, _ = json.Marshal(css)
	logger.Debugln("Storing result to \"static/output.json\" file")
	ioutil.WriteFile(filepath.Join(os.Getenv("COMMDEPATH"), "static", "output.json"), jsonBuf, os.ModeAppend)
	// Start backend engine
	router.Start()
}
