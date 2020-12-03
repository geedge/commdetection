package main

import (
	prep "commdetection/preprocessing"
	"fmt"
)

func main() {
	commands := prep.GetCommands("/root/.bash_history", "")
	methods := []func([]prep.Command) []prep.Command{prep.RemoveCommandUsingHelp}
	commands = prep.FlushCommands(commands, methods)
	fmt.Print(commands)
}
