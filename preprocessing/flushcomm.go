package preprocessing

import (
	"bytes"
	"errors"
	"os/exec"
	"regexp"
)

// FlushCommands erase all the invalid commands through various methods
func FlushCommands(commands []Command, methods []func([]Command) []Command) []Command {
	for _, method := range methods {
		commands = method(commands)
	}
	return commands
}

func removeOneCommand(commands []Command, n int) []Command {
	return append(commands[:n], commands[n+1:]...)
}

// removeCommands removes a list of commands using index
// make sure n is an ordered array
func removeCommands(commands []Command, n []int) []Command {
	if len(n) == 0 {
		return commands
	}
	if len(n) == 1 {
		return removeOneCommand(commands, n[0])
	}
	if n[0] > n[len(n)-1] {
		for i := range n {
			commands = removeOneCommand(commands, n[i])
		}
	} else {
		for i := range n {
			commands = removeOneCommand(commands, n[len(n)-1-i])
		}
	}
	return commands
}

// SimpleCommandFilter tries to run command without any flag to see whether this command can be run properly
func SimpleCommandFilter(commands []Command) []Command {
	return cmdsFilter(commands, func(command string) error {
		if commandValiderWithFlags(command, "") != nil {
			return errors.New("command not found or not in $PATH or directory")
		}
		return nil
	})
}

// HelpCommandFilter tries to use "`Command` --help", "`Command` -h" and "`Command` help" to judge `Command` is valid or not
func HelpCommandFilter(commands []Command) []Command {
	return cmdsFilter(commands, func(command string) error {
		if commandValiderWithFlags(command, "--help") != nil &&
			commandValiderWithFlags(command, "help") != nil &&
			commandValiderWithFlags(command, "-h") != nil {
			return errors.New("command not found or not in $PATH or directory")
		}
		return nil
	})
}

func commandValiderWithFlags(command string, flag string) error {
	compile, _ := regexp.Compile("((.*)[Cc]ommand(.*)not found(.*))|((.*)No such file or directory(.*))")
	var (
		stderr bytes.Buffer
		stdout bytes.Buffer
	)

	cmd := exec.Command("bash", "-c", command, flag)
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		errString := stderr.String()
		outString := stdout.String()
		if compile.MatchString(errString) || compile.MatchString(outString) {
			return err
		}
	}
	return nil
}

// ManCommandFilter tries to use "man `Command`" cmd to judge `Command` is valid or not
func ManCommandFilter(commands []Command) []Command {
	return cmdsFilter(commands, func(command string) error {
		cmd := exec.Command("man", command)
		return cmd.Run()
	})
}

func cmdsFilter(commands []Command, cmds func(string) error) []Command {
	var invalidNum []int
	for commNum, command := range commands {
		if err := cmds(command.command); err != nil {
			invalidNum = append(invalidNum, commNum)
		}
	}
	return removeCommands(commands, invalidNum)
}
