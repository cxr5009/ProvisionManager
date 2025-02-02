package programs

import (
	"github.com/cxr5009/provisionmanager/plugins/processes"
)

// SampleProgram implements ProgramPlugin.
type SampleProgram struct{}

func (sp *SampleProgram) Name() string {
	return "Sample Program"
}

// Processes returns a slice of processes to run.
func (sp *SampleProgram) Processes() []processes.ProcessPlugin {
	return []processes.ProcessPlugin{
		&processes.ShellCommandProcess{
			Command: "echo",
			Args:    []string{"Hello from Sample Program!"},
		},
		// Add more processes as needed
	}
}
