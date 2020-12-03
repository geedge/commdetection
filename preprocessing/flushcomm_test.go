package preprocessing

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFlushCommands(t *testing.T) {
	t.Run("Flush Command Using Help Test", func(t *testing.T) {
		commands := []Command{
			{
				command: "wget",
				flags:   []string{},
			},
			{
				command: "apt",
				flags:   []string{"install", "wget"},
			},
			{
				command: "./commDetection",
				flags:   []string{},
			},
			{
				command: "dlv",
				flags:   []string{"debug", "main.go"},
			},
			{
				command: "source",
				flags:   []string{".profile"},
			},
			{
				command: "go",
				flags:   []string{"build", "main.go"},
			},
		}
		predict := []Command{
			{
				command: "wget",
				flags:   []string{},
			},
			{
				command: "apt",
				flags:   []string{"install", "wget"},
			},
			{
				command: "source",
				flags:   []string{".profile"},
			},
			{
				command: "go",
				flags:   []string{"build", "main.go"},
			},
		}
		commands = FlushCommands(commands, []func([]Command) []Command{RemoveCommandUsingHelp})
		if !reflect.DeepEqual(commands, predict) {
			fmt.Println(commands)
			t.Errorf("results are not as predicted")
		}
	})
}
