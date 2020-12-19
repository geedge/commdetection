/*
 * @Author: EnderByEndera
 * @Date: 2020-12-19 11:59:02
 * @LastEditTime: 2020-12-19 17:05:53
 * @LastEditors: Please set LastEditors
 * @Description: root of the commdetection cmd
 * @FilePath: /commdetection/cmd/root.go
 */

package cmd

import (
	"commdetection/comm"
	"commdetection/logger"
	"commdetection/rules"
	"commdetection/yaml"
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
	filterStr   string = "which"
	evaluations []string
	file        string = "/root/.bash_history"
	logLevel    uint32 = 5

	filters = []comm.Filter{}
	rs      = rules.Rules{}
	buf     []byte
)

// Execute executes the command
func Execute() {
	// 初始化
	initialize()
	rootCmd.Execute()
}

func initialize() {
	rootCmd.PersistentFlags().StringVar(&filterStr, "filter", "which", "choose one filter to filt the data from the file")
	rootCmd.PersistentFlags().StringSliceVarP(&evaluations, "evaluations", "e", []string{}, "choose one or more evaluations to evaluate commands")
	rootCmd.PersistentFlags().StringVar(&file, "file", "/root/.bash_history", "choose one file storing data of the commands")
	rootCmd.PersistentFlags().Uint32Var(&logLevel, "loglevel", uint32(logrus.DebugLevel), "choose log level")
	rootCmd.AddCommand(verCmd)

	logger.Init(logrus.Level(logLevel), &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:07:05",
	})

	if err := yaml.InitYamlSetting(); err != nil {
		logger.Fatalln(err)
	}
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
	// 将命令得分保存到json文件中
	logger.Debugln("Storing result to output.json file")
	jsonBuf, _ := json.Marshal(css)
	ioutil.WriteFile(filepath.Join(os.Getenv("COMMDEPATH"), "output.json"), jsonBuf, os.ModeAppend)
}
