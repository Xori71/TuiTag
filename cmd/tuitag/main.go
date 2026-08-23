package main

import (
	tea "charm.land/bubbletea/v2"
)

func main() {
	model := Model.Init()
	for {
		app := tea.NewProgram(View(model))
	}	
}
