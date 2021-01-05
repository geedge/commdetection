/*
 * @Author: EnderByEndera
 * @Date: 2020-12-02 17:08:59
 * @LastEditTime: 2021-01-05 17:19:16
 * @LastEditors: Please set LastEditors
 * @Description: Get commands from file or network
 * @FilePath: /commdetection/preprocessing/commget.go
 */

package comm

import (
	"bytes"
	"commdetection/logger"
	"io/ioutil"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// Command contains command and its flags or symbols
type Command struct {
	CommName  string    `json:"commName"`
	Args      []string  `json:"args"`
	Flags     []string  `json:"flags"`
	TimeStamp time.Time `json:"timestamp"`
	User      string    `json:"user"`
	Mac       string    `json:"mac"`
}

// GetCommandsFromHist gets the commands by using `history` bash command
func GetCommandsFromHist() ([]Command, error) {
	var (
		stderr bytes.Buffer
		stdout bytes.Buffer
	)
	hist := exec.Command("/bin/bash", "-c", `$COMMDEPATH/history.sh`)
	hist.Stderr = &stderr
	hist.Stdout = &stdout
	err := hist.Run()
	commands := commandsFromString(stdout.String())
	return commands, err
}

func commandsFromString(s string) []Command {
	var commands []Command
	commLines := strings.Split(s, "\n")
	for _, comm := range commLines {
		newComm := Command{}
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
				timestamp += " " + commArg
				var err error
				newComm.TimeStamp, err = time.Parse("2006-01-02 15:04:05", timestamp)
				if err != nil {
					newComm.TimeStamp = time.Time{}
				}
			case num == 3:
				newComm.Mac = commArg
			case num == 4:
				newComm.User = commArg
			case num == 5:
				newComm.CommName = commArg
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

// GetCommands returns a list of commands preprocessed which first get commands from file then net
//
// If file and url is set to null, GetCommands will call getCommandsFromHistory func automatically
//TODO: Will temporarily use GetCommandsFromHist to replace GetCommands, will put commands in the mongo db later
func GetCommands() []Command {
	commands, err := GetCommandsFromHist()
	if err != nil {
		logger.Warnln(err)
		return []Command{}
	}
	return commands
}

func getCommandsFromFile(f string) ([]Command, error) {
	var commands []Command
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

func splitCommandsInLine(tComm []string) []Command {
	commands := []Command{}
	newComm := Command{}
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
		commands = append([]Command{newComm}, commands...)
	}
	return commands
}

// TODO: ready to develop until API is given
func getCommandsFromNet(url string) ([]Command, error) {
	return nil, nil
}
