package tui

import (
	"time"

	"github.com/cxr5009/provisionmanager/core"

	tea "github.com/charmbracelet/bubbletea"
)

// Define the three major states.
type State int

const (
	StateRequirements State = iota
	StateProgramSelection
	StateExecution
)

// Model holds the app state.
type Model struct {
	state           State
	message         string
	reqPlugins      []core.RequirementPlugin
	programPlugins  []core.ProgramPlugin
	selectedProgram core.ProgramPlugin
}

// NewModel returns an initial model.
func NewModel() Model {
	return Model{
		state:   StateRequirements,
		message: "Welcome to the Provisioner CLI!",
	}
}

// Init is called when the program starts.
func (m Model) Init() tea.Cmd {
	return tea.Batch(tea.EnterAltScreen, tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg{}
	}))
}

// tickMsg is a custom message type.
type tickMsg time.Time

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q": // Exit the program
			return m, tea.Quit
		case "1":
			// In a real app you'd check the requirements.
			m.message = "Checking requirements..."
			// Transition to program selection (for simplicity).
			m.state = StateProgramSelection
		case "2":
			m.message = "Program selected! Ready to execute."
			m.state = StateExecution
		case "3":
			m.message = "Executing program..."
			// Here you'd trigger the actual execution steps.
		}
	case tickMsg:
		// This could be used for periodic updates.
	}
	return m, nil
}

// View renders the UI based on the current state.
func (m Model) View() string {
	header := "Provisioner CLI\n\n"
	body := ""

	switch m.state {
	case StateRequirements:
		body += "Step 1: Requirements\n"
		body += "Press [1] to check requirements.\n"
	case StateProgramSelection:
		body += "Step 2: Program Selection\n"
		body += "Press [2] to select a program.\n"
	case StateExecution:
		body += "Step 3: Execution\n"
		body += "Press [3] to execute the program.\n"
	}
	footer := "\nPress q to quit."

	return header + m.message + "\n\n" + body + footer
}
