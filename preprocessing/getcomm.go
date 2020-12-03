package preprocessing

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// Command contains command and its flags or symbols
type Command struct {
	command string
	flags   []string
}

// GetCommands returns a list of commands preprocessed which first get commands from file then net
func GetCommands(file string, url string) []Command {
	if file != "" {
		commands, err := getCommandsFromFile(file, 100000)
		if err == nil {
			return commands
		}
		log.Fatal(err)
	}
	if url != "" {
		commands, err := getCommandsFromNet(url)
		if err == nil {
			return commands
		}
		log.Fatal(err)
	}
	log.Fatal(fmt.Errorf("cannot get commands from both file or net"))
	return nil
}

func getCommandsFromFile(f string, size int) ([]Command, error) {
	var commands []Command
	fp, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	defer fp.Close()
	buf := make([]byte, size)
	_, err = fp.Read(buf)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(buf), "\n")
	for _, line := range lines {
		tComm := strings.Split(line, " ")
		newComm := Command{}
		for commNum, comm := range tComm {
			if commNum == 0 {
				newComm.command = comm
			} else if comm != "" {
				newComm.flags = append(newComm.flags, comm)
			}
		}
		commands = append(commands, newComm)
	}
	return commands, nil
}

// TODO: ready to develop until API is given
func getCommandsFromNet(url string) ([]Command, error) {
	return nil, nil
}
