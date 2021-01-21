/*
 * @Author: EnderByEndera
 * @Date: 2020-12-07 09:22:53
 * @LastEditTime: 2021-01-14 15:31:24
 * @LastEditors: Please set LastEditors
 * @Description: this file flushes invalid commands using various types of filters
 * @FilePath: /commdetection/commflush.go
 */

package comm

import (
	"bytes"
	"commdetection/model"
	"errors"
	"os/exec"
	"regexp"
	"time"
)

// NOTFOUNDREG is the regular expression to test whether command executed properly
const NOTFOUNDREG = "((.*)[Cc]ommand(.*)not found(.*))|((.*)No such file or directory(.*))"

//MANENTRYREG is the regular expression to test whether "man `Command`" executed properly
const MANENTRYREG = "No manual entry for (.*)"

// Filter defines filter function to clean invalid commadns or ops given by file or net
type Filter func([]model.Command) []model.Command

type errorNum struct {
	err error
	num int
}

// FlushCommands erase all the invalid commands through various filters
func FlushCommands(commands []model.Command, filters []Filter) []model.Command {
	for _, filter := range filters {
		commands = filter(commands)
	}
	return commands
}

// SimpleCommandFilter tries to run command without any flag to see whether this command can be run properly
// Deprecated: use WhichCommandFilter instead
func SimpleCommandFilter(commands []model.Command) []model.Command {
	return cmdsFilter(commands, func(command string) error {
		return checkOneCmd(command, "", NOTFOUNDREG)
	})
}

// HelpCommandFilter tries to use "`Command` --help", "`Command` -h" and "`Command` help" to judge `Command` is valid or not
// Deprecated: use WhichCommandFilter instead
func HelpCommandFilter(commands []model.Command) []model.Command {
	return cmdsFilter(commands, func(command string) error {
		if checkOneCmd(command, "--help", NOTFOUNDREG) != nil {
			return errors.New("command not found or not in $PATH or directory")
		}
		return nil
	})
}

// ManCommandFilter tries to use "man `Command`" cmd to judge `Command` is valid or not
// Deprecated: use WhichCommandFilter instead
func ManCommandFilter(commands []model.Command) []model.Command {
	return cmdsFilter(commands, func(command string) error {
		return checkOneCmd("man", command, MANENTRYREG)
	})
}

// WhichCommandFilter tries to use "which `Command`" cmd to judge `Command` is valid or not
func WhichCommandFilter(commands []model.Command) []model.Command {
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

func cmdsFilter(commands []model.Command, execute func(string) error) []model.Command {
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
	return model.RemoveCommands(commands, invalidNum)
}
