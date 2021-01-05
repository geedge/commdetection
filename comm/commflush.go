/*
 * @Author: EnderByEndera
 * @Date: 2020-12-07 09:22:53
 * @LastEditTime: 2021-01-05 13:27:31
 * @LastEditors: Please set LastEditors
 * @Description: this file flushes invalid commands using various types of filters
 * @FilePath: /commdetection/commflush.go
 */

package comm

import (
	"bytes"
	"errors"
	"os/exec"
	"regexp"
	"sort"
	"time"
)

// NOTFOUNDREG is the regular expression to test whether command executed properly
const NOTFOUNDREG = "((.*)[Cc]ommand(.*)not found(.*))|((.*)No such file or directory(.*))"

//MANENTRYREG is the regular expression to test whether "man `Command`" executed properly
const MANENTRYREG = "No manual entry for (.*)"

// Filter defines filter function to clean invalid commadns or ops given by file or net
type Filter func([]Command) []Command

//IntSlice includes a list of n array in order to use sort.Sort() method to sort n
type IntSlice []int

func (s IntSlice) Len() int           { return len(s) }
func (s IntSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s IntSlice) Less(i, j int) bool { return s[i] > s[j] }

type errorNum struct {
	err error
	num int
}

// FlushCommands erase all the invalid commands through various filters
func FlushCommands(commands []Command, filters []Filter) []Command {
	for _, filter := range filters {
		commands = filter(commands)
	}
	return commands
}

// removeOneCommand removes only one command from the list given the index
func removeOneCommand(commands []Command, n int) []Command {
	return append(commands[:n], commands[n+1:]...)
}

// removeCommands removes a list of commands using index
func removeCommands(commands []Command, n IntSlice) []Command {
	if len(n) == 0 {
		return commands
	}
	sort.Sort(n)
	for i := range n {
		commands = removeOneCommand(commands, n[i])
	}
	return commands
}

// SimpleCommandFilter tries to run command without any flag to see whether this command can be run properly
// Deprecated: use WhichCommandFilter instead
func SimpleCommandFilter(commands []Command) []Command {
	return cmdsFilter(commands, func(command string) error {
		return checkOneCmd(command, "", NOTFOUNDREG)
	})
}

// HelpCommandFilter tries to use "`Command` --help", "`Command` -h" and "`Command` help" to judge `Command` is valid or not
// Deprecated: use WhichCommandFilter instead
func HelpCommandFilter(commands []Command) []Command {
	return cmdsFilter(commands, func(command string) error {
		if checkOneCmd(command, "--help", NOTFOUNDREG) != nil {
			return errors.New("command not found or not in $PATH or directory")
		}
		return nil
	})
}

// ManCommandFilter tries to use "man `Command`" cmd to judge `Command` is valid or not
// Deprecated: use WhichCommandFilter instead
func ManCommandFilter(commands []Command) []Command {
	return cmdsFilter(commands, func(command string) error {
		return checkOneCmd("man", command, MANENTRYREG)
	})
}

// WhichCommandFilter tries to use "which `Command`" cmd to judge `Command` is valid or not
func WhichCommandFilter(commands []Command) []Command {
	return cmdsFilter(commands, func(command string) error {
		return checkOneCmd("which", command, "")
	})
}

func checkOneCmd(command, arg, reg string) error {
	compile, _ := regexp.Compile(reg)
	var (
		stderr bytes.Buffer
		stdout bytes.Buffer
	)
	cmd := exec.Command(command, arg)
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout
	if err := cmd.Start(); err != nil {
		return err
	}
	// Wait for the process to finish or kill it after a timeout (whichever happens first):
	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()
	select {
	case <-time.After(400 * time.Millisecond):
		if err := cmd.Process.Kill(); err != nil {
			return err
		}
	case err := <-done:
		if err != nil {
			if compile.MatchString(stderr.String()) || compile.MatchString(stdout.String()) {
				return err
			}
		}
	}
	return nil
}

func cmdsFilter(commands []Command, execute func(string) error) []Command {
	var invalidNum []int
	ec := make(chan errorNum, 1)
	for commNum, command := range commands {
		go func(command string, commNum int) {
			ec <- errorNum{execute(command), commNum}
		}(command.CommName, commNum)
	}
	for range commands {
		if e, ok := <-ec; ok && e.err != nil {
			invalidNum = append(invalidNum, e.num)
		}
	}
	return removeCommands(commands, invalidNum)
}
