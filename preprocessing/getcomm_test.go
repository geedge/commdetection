package preprocessing

import (
	"testing"
)

func TestGetCommands(t *testing.T) {
	t.Run("Get Commands From File Test", func(t *testing.T) {
		commands := GetCommands("/root/.bash_history", "")
		runCommands(commands, t)
	})

	t.Run("Get Commands From Net Test", func(t *testing.T) {
		commands := GetCommands("", "http://127.0.0.1:8080")
		runCommands(commands, t)
	})

	t.Run("Get commands from incorrect file test", func(t *testing.T) {
		commands := GetCommands("/root/aaa", "")
		runCommands(commands, t)
	})

	t.Run("Get Commands from Both Test", func(t *testing.T) {
		commands := GetCommands("/root/.bash_history", "http://127.0.0.1:8080")
		runCommands(commands, t)
	})
}

func runCommands(commands []Command, t *testing.T) {
	for commNum, command := range commands {
		if command.command == "" {
			t.Errorf("command %d is empty", commNum)
		}
		for flagNum, flag := range command.flags {
			if flag == "" {
				t.Errorf("command %d: %s flag %d is empty", commNum+1, command.command, flagNum)
			}
		}
	}
}
