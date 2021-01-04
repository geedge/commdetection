/*
 * @Author: EnderByEndera
 * @Date: 2020-12-02 17:08:59
 * @LastEditTime: 2021-01-04 15:59:10
 * @LastEditors: Please set LastEditors
 * @Description: Get commands from file or network
 * @FilePath: /commdetection/preprocessing/commget.go
 */

package comm

import (
	"commdetection/logger"
	"io/ioutil"
	"reflect"
	"strings"
)

// Command contains command and its flags or symbols
type Command struct {
	CommName string   `json:"commName"`
	Args     []string `json:"args"`
	Flags    []string `json:"flags,omitempty"`
}

// GetCommands returns a list of commands preprocessed which first get commands from file then net
func GetCommands(file string, url string) []Command {
	if file != "" {
		commands, err := getCommandsFromFile(file)
		if err == nil {
			return commands
		}
		logger.Warnln(err)
	}
	if url != "" {
		commands, err := getCommandsFromNet(url)
		if err == nil {
			return commands
		}
		logger.Warnln(err)
	}
	logger.Warnln("cannot get commands from any file or net")
	return []Command{}
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
	newComm := Command{
		CommName: "",
		Args:     []string{},
		Flags:    []string{},
	}
	for commNum, comm := range tComm {
		if commNum == 0 {
			newComm.CommName = comm
		} else if comm == "&&" { // comm is the separator of the whole commmand
			commands = append(commands, splitCommandsInLine(tComm[commNum+1:])...)
			break
		} else if comm[0] != '-' { // comm is a flag
			newComm.Args = append(newComm.Args, comm)
		} else { // comm is just a normal argument
			newComm.Flags = append(newComm.Flags, comm)
		}
	}
	if !reflect.DeepEqual(newComm, Command{
		CommName: "",
		Args:     []string{},
		Flags:    []string{},
	}) {
		commands = append([]Command{newComm}, commands...)
	}
	return commands
}

// TODO: ready to develop until API is given
func getCommandsFromNet(url string) ([]Command, error) {
	return nil, nil
}
