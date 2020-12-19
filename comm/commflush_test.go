package comm

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFlushCommands(t *testing.T) {
	t.Run("FlushCommands Test Using SimpleCommandFilter", func(t *testing.T) {
		commands := []Command{
			{
				CommName: "wget",
				Flags:    []string{},
			},
			{
				CommName: "apt",
				Flags:    []string{"install", "wget"},
			},
			{
				CommName: "./commDetection",
				Flags:    []string{},
			},
			{
				CommName: "vim",
				Flags:    []string{"debug", "main.go"},
			},
			{
				CommName: "source",
				Flags:    []string{".profile"},
			},
			{
				CommName: "go",
				Flags:    []string{"build", "main.go"},
			},
			{
				CommName: "aasdf",
				Flags:    []string{},
			},
		}
		predict := []Command{
			{
				CommName: "wget",
				Flags:    []string{},
			},
			{
				CommName: "apt",
				Flags:    []string{"install", "wget"},
			},
			{
				CommName: "vim",
				Flags:    []string{"debug", "main.go"},
			},
			{
				CommName: "go",
				Flags:    []string{"build", "main.go"},
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
				CommName: "wget",
				Flags:    []string{},
			},
			{
				CommName: "apt",
				Flags:    []string{"install", "wget"},
			},
			{
				CommName: "dlv",
				Flags:    []string{"debug", "main.go"},
			},
			{
				CommName: "source",
				Flags:    []string{".profile"},
			},
			{
				CommName: "go",
				Flags:    []string{"build", "main.go"},
			},
			{
				CommName: "aasdf",
				Flags:    []string{},
			},
		}
		predict := []Command{
			{
				CommName: "wget",
				Flags:    []string{},
			},
			{
				CommName: "apt",
				Flags:    []string{"install", "wget"},
			},
			{
				CommName: "go",
				Flags:    []string{"build", "main.go"},
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
				CommName: "wget",
				Flags:    []string{},
			},
			{
				CommName: "apt",
				Flags:    []string{"install", "wget"},
			},
			{
				CommName: "vim",
				Flags:    []string{"debug", "main.go"},
			},
			{
				CommName: "./commDetection",
				Flags:    []string{},
			},
			{
				CommName: "source",
				Flags:    []string{".profile"},
			},
			{
				CommName: "go",
				Flags:    []string{"build", "main.go"},
			},
			{
				CommName: "aasdf",
				Flags:    []string{},
			},
		}
		predict := []Command{
			{
				CommName: "wget",
				Flags:    []string{},
			},
			{
				CommName: "apt",
				Flags:    []string{"install", "wget"},
			},
			{
				CommName: "vim",
				Flags:    []string{"debug", "main.go"},
			},
		}
		commands = FlushCommands(commands, []Filter{ManCommandFilter})
		if !reflect.DeepEqual(commands, predict) {
			fmt.Println(commands)
			t.Errorf("results are not as predicted")
		}
	})

	t.Run("FlushCommands with no commands", func(t *testing.T) {
		commands := []Command{}
		commands = FlushCommands(commands, []Filter{SimpleCommandFilter})
		if !reflect.DeepEqual(commands, []Command{}) {
			fmt.Println(commands)
			t.Errorf("results are not as predicted")
		}
	})
}

func BenchmarkFlushCommands(b *testing.B) {
	b.Run("FlushCommands BenchMark Test", func(b *testing.B) {
		filters := []Filter{WhichCommandFilter}
		commands := GetCommands("/root/.bash_history", "")
		b.ResetTimer()
		commands = FlushCommands(commands, filters)
		b.StopTimer()
	})
}
