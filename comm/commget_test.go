/*
 * @Author: your name
 * @Date: 2020-12-02 17:09:14
 * @LastEditTime: 2020-12-19 10:42:59
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /commdetection/preprocessing/commget_test.go
 */
package comm

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

	t.Run("Get commands from no-existence file test", func(t *testing.T) {
		commands := GetCommands("dsfa", "")
		runCommands(commands, t)
	})

	t.Run("Get commands from no-existence net", func(t *testing.T) {
		commands := GetCommands("", "sdfsagd")
		runCommands(commands, t)
	})

	t.Run("Get commands from nothing", func(t *testing.T) {
		commands := GetCommands("", "")
		runCommands(commands, t)
	})

	t.Run("Get Commands from Both Test", func(t *testing.T) {
		commands := GetCommands("/root/.bash_history", "http://127.0.0.1:8080")
		runCommands(commands, t)
	})
}

func runCommands(commands []Command, t *testing.T) {
	for commNum, command := range commands {
		if command.CommName == "" {
			t.Errorf("command %d is empty", commNum)
		}
		for flagNum, flag := range command.Flags {
			if flag == "" {
				t.Errorf("command %d: %s flag %d is empty", commNum+1, command.CommName, flagNum)
			}
		}
	}
}

func BenchmarkGetCommands(b *testing.B) {
	GetCommands("/root/.bash_history", "")
}
