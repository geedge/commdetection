/*
 * @Author: EnderByEndera
 * @Date: 2020-12-02 17:08:59
 * @LastEditTime: 2021-01-16 08:39:11
 * @LastEditors: Please set LastEditors
 * @Description: Get commands from file or network
 * @FilePath: /commdetection/preprocessing/commget.go
 */

package comm

import (
	"bytes"
	"commdetection/logger"
	"commdetection/model"
	"io/ioutil"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"time"
)

// GetCommands returns a list of commands preprocessed which first get commands from file then net
func GetCommands() model.Commands {
	commands, err := getCommandsFromHist()
	if err != nil {
		logger.Warnln(err)
		return []model.Command{}
	}
	sort.Sort(commands)
	return commands
}

// getCommandsFromHist gets the commands by using `history` bash command
func getCommandsFromHist() (model.Commands, error) {
	var (
		stderr bytes.Buffer
		stdout bytes.Buffer
	)
	hist := exec.Command("/bin/bash", "-c", `$COMMDEPATH/history.sh`)
	hist.Stderr = &stderr
	hist.Stdout = &stdout
	err := hist.Run()
	if err != nil {
		return nil, err
	}
	commands := commandsFromString(stdout.String())
	return commands, nil
}

func commandsFromString(s string) model.Commands {
	var commands model.Commands
	commLines := strings.Split(s, "\n")
	for _, comm := range commLines {
		newComm := model.Command{}
		var timestamp string
		for num, commArg := range strings.Fields(comm) {
			if _, err := strconv.Atoi(commArg); err != nil && num == 0 {
				break
			}
			switch {
			case num == 0:
				continue
			case num == 1:
				timestamp = commArg
			case num == 2:
				timestamp += " " + commArg // create the timestamp string to parse
				var err error
				newComm.TimeStamp, err = time.Parse("2006-01-02 15:04:05", timestamp)
				if err != nil {
					newComm.TimeStamp = time.Time{}
				}
			case num == 3:
				newComm.Mac = commArg // comm[3] arg is Mac Address
			case num == 4:
				newComm.User = commArg // comm[4] arg is Username
			case num == 5:
				newComm.CommName = commArg // comm[5] arg is command name
			default:
				if commArg[0] != '-' { // comm is a flag
					newComm.Args = append(newComm.Args, commArg)
				} else { // comm is just a normal argument
					newComm.Flags = append(newComm.Flags, commArg)
				}
			}
		}
		if newComm.CommName != "" {
			commands = append(commands, newComm)
		}
	}
	return commands
}

// TODO: will promote it later
func getCommandsFromDB() (model.Commands, error) {
	commands := new(model.Commands)
	commands.GetCommandsFrom("test", "commands")
	return *commands, nil
}

func getCommandsFromFile(f string) ([]model.Command, error) {
	var commands []model.Command
	buf, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(buf), "\n")
	for _, line := range lines {
		commands = append(commands, splitCommandsInLine(strings.Fields(line))...)
	}
	return commands, nil
}

func splitCommandsInLine(tComm []string) []model.Command {
	commands := []model.Command{}
	newComm := model.Command{}
	for commNum, comm := range tComm {
		if commNum == 0 {
			newComm.CommName = comm
		} else if comm == "&&" { // comm is the separator of the whole command
			commands = append(commands, splitCommandsInLine(tComm[commNum+1:])...)
			break
		} else if comm[0] != '-' { // comm is a flag
			newComm.Args = append(newComm.Args, comm)
		} else { // comm is just a normal argument
			newComm.Flags = append(newComm.Flags, comm)
		}
	}
	if newComm.CommName != "" {
		commands = append([]model.Command{newComm}, commands...)
	}
	return commands
}
