/*
 * @Author: EnderByEndera
 * @Date: 2020-12-02 17:08:59
 * @LastEditTime: 2020-12-09 10:27:45
 * @LastEditors: Please set LastEditors
 * @Description: Get commands from file or network
 * @FilePath: /commdetection/preprocessing/commget.go
 */

package preprocessing

import (
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"strings"
)

// Command contains command and its flags or symbols
type Command struct {
	CommName string   `json:"commName"`
	Flags    []string `json:"flags"`
}

// GetCommands returns a list of commands preprocessed which first get commands from file then net
func GetCommands(file string, url string) []Command {
	if file != "" {
		commands, err := getCommandsFromFile(file)
		if err == nil {
			return commands
		}
		log.Print(err)
	}
	if url != "" {
		commands, err := getCommandsFromNet(url)
		if err == nil {
			return commands
		}
		log.Print(err)
	}
	log.Print(fmt.Errorf("cannot get commands from any file or net"))
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
		tComm := strings.Fields(line)
		newComm := Command{}
		for commNum, comm := range tComm {
			if commNum == 0 {
				newComm.CommName = comm
			} else if comm != "" {
				newComm.Flags = append(newComm.Flags, comm)
			}
		}
		if !reflect.DeepEqual(newComm, Command{}) {
			commands = append(commands, newComm)
		}
	}
	return commands, nil
}

// TODO: ready to develop until API is given
func getCommandsFromNet(url string) ([]Command, error) {
	return nil, nil
}
