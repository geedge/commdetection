package preprocessing

import (
	"fmt"
	"log"
	"os/exec"
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

// RemoveCommandUsingHelp tries to use "`Command` --help", "`Command` -h" and "`Command` help" to judge `Command` is valid or not
func RemoveCommandUsingHelp(commands []Command) []Command {
	return removeCommandsUsingCmds(commands, func(command string) error {
		if exec.Command(command, "--help").Run() != nil &&
			exec.Command(command, "-h").Run() != nil &&
			exec.Command(command, "help").Run() != nil {
			return fmt.Errorf("cmd %s help guide not exist", command)
		}
		return nil
	})
}

// RemoveCommandUsingMan tries to use "man `Command`" cmd to judge `Command` is valid or not
func RemoveCommandUsingMan(commands []Command) []Command {
	return removeCommandsUsingCmds(commands, func(command string) error {
		cmd := exec.Command("man", command)
		return cmd.Run()
	})
}

func removeCommandsUsingCmds(commands []Command, cmds func(string) error) []Command {
	var invalidNum []int
	for commNum, command := range commands {
		if err := cmds(command.command); err != nil {
			log.Print(err)
			invalidNum = append(invalidNum, commNum)
		}
	}
	return removeCommands(commands, invalidNum)
}
