package main

import (
	prep "commdetection/preprocessing"
)

func main() {
	filters := []prep.Filter{prep.HelpCommandFilter}
	commands := prep.GetCommands("/root/.bash_history", "")
	commands = prep.FlushCommands(commands, filters)
}
