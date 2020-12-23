/*
 * @Author: your name
 * @Date: 2020-12-02 17:09:14
 * @LastEditTime: 2020-12-23 11:46:52
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /commdetection/preprocessing/commget_test.go
 */
package comm

import (
	"fmt"
	"reflect"
	"strings"
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

func TestSplitCommandsInLine(t *testing.T) {
	t.Run("split large command test", func(t *testing.T) {
		str := "ls && cp -s abc && goc"
		commands := splitCommandsInLine(strings.Fields(str))
		predict := []Command{
			{
				CommName: "ls",
				Args:     []string{},
			},
			{
				CommName: "cp",
				Args:     []string{"-s", "abc"},
			},
			{
				CommName: "goc",
				Args:     []string{},
			},
		}
		if !reflect.DeepEqual(commands, predict) {
			fmt.Println(commands)
			fmt.Println(predict)
			t.Errorf("results are not as predicted")
		}
	})
}

func runCommands(commands []Command, t *testing.T) {
	for commNum, command := range commands {
		if command.CommName == "" {
			t.Errorf("command %d is empty", commNum)
		}
		for flagNum, flag := range command.Args {
			if flag == "" {
				t.Errorf("command %d: %s flag %d is empty", commNum+1, command.CommName, flagNum)
			}
		}
	}
}

func BenchmarkGetCommands(b *testing.B) {
	GetCommands("/root/.bash_history", "")
}
