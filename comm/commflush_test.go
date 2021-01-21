/*
 * @Author: your name
 * @Date: 2020-12-03 12:51:28
 * @LastEditTime: 2021-01-16 08:41:14
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /commdetection/comm/commflush_test.go
 */
package comm

import (
	"commdetection/model"
	"fmt"
	"reflect"
	"testing"
)

func TestFlushCommands(t *testing.T) {
	t.Run("FlushCommands Test Using WhichCommandFilter", func(t *testing.T) {
		commands := []model.Command{
			{
				CommName: "wget",
				Args:     []string{},
			},
			{
				CommName: "echo",
				Args:     []string{"install", "wget"},
			},
			{
				CommName: "./commDetection",
				Args:     []string{},
			},
			{
				CommName: "vim",
				Args:     []string{"debug", "main.go"},
			},
			{
				CommName: "source",
				Args:     []string{".profile"},
			},
			{
				CommName: "go",
				Args:     []string{"build", "main.go"},
			},
			{
				CommName: "aasdf",
				Args:     []string{},
			},
		}
		predict := []model.Command{
			{
				CommName: "wget",
				Args:     []string{},
			},
			{
				CommName: "echo",
				Args:     []string{"install", "wget"},
			},
			{
				CommName: "vim",
				Args:     []string{"debug", "main.go"},
			},
			{
				CommName: "go",
				Args:     []string{"build", "main.go"},
			},
		}
		commands = FlushCommands(commands, []Filter{WhichCommandFilter})
		if !reflect.DeepEqual(commands, predict) {
			fmt.Println(commands)
			t.Errorf("results are not as predicted")
		}
	})

	t.Run("FlushCommands with no commands", func(t *testing.T) {
		commands := []model.Command{}
		commands = FlushCommands(commands, []Filter{WhichCommandFilter})
		if !reflect.DeepEqual(commands, []model.Command{}) {
			fmt.Println(commands)
			t.Errorf("results are not as predicted")
		}
	})

	t.Run("FlushCommands with mongodb", func(t *testing.T) {
		comms := GetCommands()
		comms = FlushCommands(comms, []Filter{WhichCommandFilter})
		commands := model.Commands(comms)
		err := commands.InsertAllTo("test", "commands")
		err = commands.InsertAllTo("test", "commands")
		if err != nil {
			t.Error(err)
		}

	})
}

func BenchmarkFlushCommands(b *testing.B) {
	b.Run("FlushCommands BenchMark Test", func(b *testing.B) {
		filters := []Filter{WhichCommandFilter}
		commands := GetCommands()
		b.ResetTimer()
		commands = FlushCommands(commands, filters)
		b.StopTimer()
	})
}

func BenchmarkInsertFromMongo(b *testing.B) {
	commands := GetCommands()
	commands = FlushCommands(commands, []Filter{WhichCommandFilter})
	commands.InsertAllTo("test", "commands")
	commands.InsertAllTo("test", "commands")
}
