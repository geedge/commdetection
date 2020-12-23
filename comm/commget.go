/*
 * @Author: EnderByEndera
 * @Date: 2020-12-02 17:08:59
 * @LastEditTime: 2020-12-23 11:58:08
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
	Args     []string `json:"flags"`
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
	}
	for commNum, comm := range tComm {
		if commNum == 0 {
			newComm.CommName = comm
		} else if comm != "&&" {
			newComm.Args = append(newComm.Args, comm)
		} else {
			commands = append(commands, splitCommandsInLine(tComm[commNum+1:])...)
			break
		}
	}
	if !reflect.DeepEqual(newComm, Command{}) {
		commands = append([]Command{newComm}, commands...)
	}
	return commands
}

// TODO: ready to develop until API is given
func getCommandsFromNet(url string) ([]Command, error) {
	return nil, nil
}
