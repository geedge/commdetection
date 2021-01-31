/*
 * @Author: EnderByEndera
 * @Date: 2020-12-19 11:59:02
 * @LastEditTime: 2021-02-01 06:53:44
 * @LastEditors: Please set LastEditors
 * @Description: root of the commdetection cmd
 * @FilePath: /commdetection/cmd/root.go
 */

package cmd

import (
	"commdetection/comm"
	"commdetection/logger"
	"commdetection/model"
	"commdetection/router"
	"commdetection/rules"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

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
	rs      = model.Rules{}
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
	logger.LogLevel = logrus.Level(logLevel)
}

func root() {
	for _, ev := range evaluations {
		switch ev {
		case "command":
			rs = rules.AddRule(rs, model.Rule{
				Name:     "Command",
				RuleFunc: "EvaluateCommandRule",
			})
		case "path":
			rs = rules.AddRule(rs, model.Rule{
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

	go func() {
		for {
			var comms model.Commands
			comms = comm.GetCommands()
			comms = comm.FlushCommands(comms, filters)
			comms.InsertAllTo("test", "commands")
			comms.GetCommandsFrom("test", "commands")
			css := rules.InitCommScores(comms)
			css = rules.EvaluateCommScore(css, rs)
			jsonBuf, _ := json.Marshal(css)
			ioutil.WriteFile(filepath.Join(os.Getenv("COMMDEPATH"), "static", "base", "output.json"), jsonBuf, os.ModeAppend)
			logger.Debugln("New output.json file is built")
			time.Sleep(3 * time.Minute)
		}
	}()

	router.StartRouter()
}
