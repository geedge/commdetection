/*
 * @Author: your name
 * @Date: 2020-12-02 17:09:14
 * @LastEditTime: 2021-01-05 17:25:02
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
	t.Run("Get Commands From history Test", func(t *testing.T) {
		commands, err := GetCommandsFromHist()
		if err != nil {
			t.Error(err)
		}
		checkCommands(commands, t)
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
				Flags:    []string{},
			},
			{
				CommName: "cp",
				Args:     []string{"abc"},
				Flags:    []string{"-s"},
			},
			{
				CommName: "goc",
				Args:     []string{},
				Flags:    []string{},
			},
		}
		if !reflect.DeepEqual(commands, predict) {
			fmt.Println(commands)
			fmt.Println(predict)
			t.Log("result is not as predicted")
		}
	})
}

func checkCommands(commands []Command, t *testing.T) {
	for commNum, command := range commands {
		if command.CommName == "" {
			t.Errorf("command %d is empty", commNum)
		}
		for flagNum, flag := range command.Args {
			if flag == "" {
				t.Errorf("command %d: %s flag %d is empty", commNum+1, command.CommName, flagNum)
			}
		}
		if command.TimeStamp.Unix() <= 0.0 {
			t.Errorf("command %d: %s timestamp is empty", commNum+1, command.CommName)
		}
	}
}

func BenchmarkGetCommands(b *testing.B) {
	GetCommands()
}
