package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/cxr5009/provisionmanager/tui"

	// Import your plugin packages so they register or are available for use.
	_ "github.com/cxr5009/provisionmanager/plugins/processes"
	_ "github.com/cxr5009/provisionmanager/plugins/programs"
	_ "github.com/cxr5009/provisionmanager/plugins/requirements"
)

func main() {
	// Here you might initialize and load your plugins.
	// For example, load all available RequirementPlugin, ProgramPlugin, etc.
	// For simplicity, we omit dynamic loading. You can use a registery patter:
	// reqRegistery := []RequirementPlugin{ &reqs.ADBPlugin{}, &reqs.LSUSBPlugin{} }
	// progRegistery := []ProgramPlugin{ &progs.SampleProgram{} }

	p := tea.NewProgram(tui.NewModel())
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
