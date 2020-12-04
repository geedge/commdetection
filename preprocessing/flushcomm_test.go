package preprocessing

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFlushCommands(t *testing.T) {
	t.Run("FlushCommands Test Using SimpleCommandFilter", func(t *testing.T) {
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
				command: "vim",
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
			{
				command: "aasdf",
				flags:   []string{},
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
				command: "vim",
				flags:   []string{"debug", "main.go"},
			},
			{
				command: "go",
				flags:   []string{"build", "main.go"},
			},
		}
		commands = FlushCommands(commands, []Filter{SimpleCommandFilter})
		if !reflect.DeepEqual(commands, predict) {
			fmt.Println(commands)
			t.Errorf("results are not as predicted")
		}
	})

	t.Run("FlushCommands Test Using HelpCommandFilter", func(t *testing.T) {
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
			{
				command: "aasdf",
				flags:   []string{},
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
				command: "go",
				flags:   []string{"build", "main.go"},
			},
		}
		commands = FlushCommands(commands, []Filter{HelpCommandFilter})
		if !reflect.DeepEqual(commands, predict) {
			fmt.Println(commands)
			t.Errorf("results are not as predicted")
		}
	})

	t.Run("FlushCommands Test Using ManCommandFilter", func(t *testing.T) {
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
				command: "vim",
				flags:   []string{"debug", "main.go"},
			},
			{
				command: "./commDetection",
				flags:   []string{},
			},
			{
				command: "source",
				flags:   []string{".profile"},
			},
			{
				command: "go",
				flags:   []string{"build", "main.go"},
			},
			{
				command: "aasdf",
				flags:   []string{},
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
				command: "vim",
				flags:   []string{"debug", "main.go"},
			},
		}
		commands = FlushCommands(commands, []Filter{ManCommandFilter})
		if !reflect.DeepEqual(commands, predict) {
			fmt.Println(commands)
			t.Errorf("results are not as predicted")
		}
	})
}

func BenchmarkFlushCommands(b *testing.B) {
	b.Run("FlushCommands BenchMark Test", func(b *testing.B) {
		filters := []Filter{SimpleCommandFilter, HelpCommandFilter, ManCommandFilter}
		commands := GetCommands("/root/.bash_history", "")
		b.ResetTimer()
		commands = FlushCommands(commands, filters)
		b.StopTimer()
	})
}
